# staticsvr

Serve static resource from one command!

install:
```bash
go install github.com/vogo/staticsvr@master
```

usage: 
```bash
staticsvr <PORT> <PATH>
```

examples:
```bash
# serve at port 9999 for directory /opt/html
staticsvr 9999 /opt/html

# serve at port 8080 for current directory
staticsvr 8080
```
