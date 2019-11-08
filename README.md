# clipload-server
> [clipload](https://github.com/Bjornskjald/clipload) compatible server

Building:
```
go build -ldflags "-X main.BaseUrl=https://image-api.tld" server.go
```
or
```
docker build --build-arg baseURL=https://image-api.tld -t clipload-server .
```