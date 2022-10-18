# Demo

To run this simply do the following:

1. `docker-compose up`
2. `curl localhost:8001 --cookie "cookieName1=value1" --cookie "cookieName2=value2" --cookie "randomCookie=randomvalue"`

The whoami service will simply return information about the request it
received. You can use this to play with the plugin and see how it works.

After running the request above you will see output like below. This
confirms that the two cookies specified in the configuration have been saved
but the `randomCookie` has been removed because it is not in the list of 
cookies that should be kept.

```
Hostname: d4b2a5484dda
IP: 127.0.0.1
IP: 172.21.0.2
RemoteAddr: 172.21.0.3:39314
GET / HTTP/1.1
Host: localhost:8001
User-Agent: curl/7.79.1
Accept: */*
Accept-Encoding: gzip
Cookie: cookieName1=value1; cookieName2=value2
X-Forwarded-For: 172.21.0.1
X-Forwarded-Host: localhost:8001
X-Forwarded-Port: 8001
X-Forwarded-Proto: http
X-Forwarded-Server: 5d856c1010c0
X-Real-Ip: 172.21.0.1
```
