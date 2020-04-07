package main

import (
 "flag"
 "fmt"
 "github.com/oschwald/geoip2-golang"
 "github.com/sevlyar/go-daemon"
 "io/ioutil"
 "log"
 "net"
 "net/http"
 "strconv"
 "syscall"
)

var (
 port     = flag.String("port", "8080", "Listen port")
 mmdbfile = flag.String("mmdb", "GeoLite2-City.mmdb", "mmdb database file")
 stop     = flag.Bool("stop", false, "Stop server")

 PIDFILE = "mygeoip.pid"
 LOGFILE = "mygeoip.log"
)

func mmdb() *geoip2.Reader {
 // GeoLite2-City have copyright, tip: use dork
 db, err := geoip2.Open(*mmdbfile)
 if err != nil {
  log.Printf("Cannot Open GeoLite2\n")
 }
 return db
}

func retIP(RemoteAddr string) string {
 ip := net.ParseIP(RemoteAddr)
 db := mmdb()
 re, err := db.City(ip)
 if err != nil {
  log.Printf("Cannot get %s location\n", RemoteAddr)
 }

 ipinfo := RemoteAddr
 ipinfo += "," + re.Country.Names["en"]
 ipinfo += "," + re.City.Names["en"]

 return ipinfo
}

func handler(w http.ResponseWriter, r *http.Request) {
 ipraddr := strings.Split(r.RemoteAddr, ":")
 ipinfo := retIP(ipraddr[0])
 fmt.Fprint(w, ipinfo, r.URL.Path[1:])
}

func Server() {
 http.HandleFunc("/", handler)
 log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func stopServer() {
 pidf, err := ioutil.ReadFile(PIDFILE)
 if err != nil {
  log.Fatal("[MyGeoIP] Can't find the PID FILE ")
 }
 pid, _ := strconv.Atoi(string(pidf))
 err = syscall.Kill(pid, 15)
 if err != nil {
  log.Fatal("[MyGeoIP] Can't stop server.. try SIGINT directly")
 }
 log.Fatal("[MyGeoIP] Server Stopped")
}

func main() {
 flag.Parse()

 if *stop {
  stopServer()
 }

 cntxt := &daemon.Context{
  PidFileName: PIDFILE,
  PidFilePerm: 0600,
  LogFileName: LOGFILE,
  LogFilePerm: 0600,
  WorkDir:     "./",
  Umask:       027,
  Args:        []string{"[MyGeoIP]"},
 }

 d, err := cntxt.Reborn()
 if err != nil {
  log.Fatal("[MyGeoIP] Unable to run/Server already Online\n")
 }
 if d != nil {
  return
 }
 defer cntxt.Release()
 log.Printf("--------\n[MyGeoIP] Server Online\n")

 Server()
}
