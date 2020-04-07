## MyGeoIP
[![https://api.travis-ci.org/pyperanger/mygeoip/](https://api.travis-ci.org/pyperanger/mygeoip.svg?branch=master&status=passed)](https://api.travis-ci.org/pyperanger/mygeoip/)
Simple and Custom HTTP Server with MyIP + Geolocation written in Go.<br>
<b>Alert</b>: mmdb is a private file, you can easily find dorking in google

### Default Return
```
YOUR-IP, Country, City
```
You can fork this project freely, I recommend read [github.com/oschwald/geoip2-golang](https://github.com/oschwald/geoip2-golang) documentation. 

### Installation
```
git clone https://github.com/pyperanger/MyGeoIP.git
cd MyGeoIP
go get ./...
go build
```

### Usage
```
  -mmdb string
     mmdb database file (default "GeoLite2-City.mmdb")
  -port string
     Listen port (default "8080")
  -stop
     Stop server
```

#### Dependencies/Ref
GeoIP: [github.com/oschwald/geoip2-golang](https://github.com/oschwald/geoip2-golang)<br>
go-daemon: [github.com/sevlyar/go-daemon](https://github.com/sevlyar/go-daemon)
