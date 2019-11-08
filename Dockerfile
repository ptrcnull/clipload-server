ARG baseURL

FROM golang:latest as builder

LABEL maintainer="Bjornskjald <docker@bjorn.ml>"

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server -ldflags "-X main.BaseUrl=$baseURL" server.go

FROM scratch

WORKDIR /app
COPY --from=builder /src/server .
VOLUME /app/img

CMD "/app/server"