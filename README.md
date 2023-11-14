# mac2manuf
Identify manufacturer information by mac address

# Identify devices
```
$ cat manuf|awk '{print $1}'|sort -u|wc -l
   34865
```

# How use
```
import github.com/hktalent/mac2manuf

fmt.Println(mac2manuf.Search("90:9c:4a:00:14:45"))
```
out:
Apple

