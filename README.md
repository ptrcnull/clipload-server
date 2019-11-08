# clipload-server
> [clipload](https://github.com/Bjornskjald/clipload) compatible server

Building:
```
go build server.go
```
or
```
docker build -t clipload-server .
```

Running:
```
docker run -d --name clipload-server -p 8012:8012 -v $(pwd)/img:/app/img clipload-server
```