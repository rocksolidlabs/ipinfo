# ipinfo

ipinfo command line tool to lookup IP information


lookup ipinfo recursivly for a domains IPs
```
rslabs@ubuntu:~$ ipinfo domain google.com
Getting IP info for domain: google.com

IP = 172.217.6.46
Hostname = sfo03s08-in-f14.1e100.net
Location = 37.4192,-122.0574
Organization = AS15169 Google Inc.
City = Mountain View
Region = California
Country = US
Postal = 94043

IP = 2607:f8b0:4005:809::200e
Hostname = sfo03s08-in-x0e.1e100.net
Location = 37.7510,-97.8220
Organization = AS15169 Google Inc.
City = 
Region = 
Country = US
Postal = 
```

lookup ipinfo for a given IP
```
rslabs@ubuntu:~$ ipinfo ip 8.8.8.8
Getting IP info for domain: 8.8.8.8

IP = 8.8.8.8
Hostname = google-public-dns-a.google.com
Location = 37.3860,-122.0838
Organization = AS15169 Google Inc.
City = Mountain View
Region = California
Country = US
Postal = 94035
```

lookup your own IP info
```
rslabs@ubuntu:~$ ipinfo self
Getting IP info for self:

IP = 8.8.8.8
Hostname = google-public-dns-a.google.com
Location = 37.3860,-122.0838
Organization = AS15169 Google Inc.
City = Mountain View
Region = California
Country = US
Postal = 94035
```
