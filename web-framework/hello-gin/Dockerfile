# compile
FROM golang:1.15.1 AS builder

WORKDIR /opt/app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go mod verify
 
RUN go build -o /out/app .

# runtime
FROM golang:1.15.1 AS bin

ARG GIN_MODE=release

WORKDIR /etc/app

COPY --from=builder /out/app /etc/app

EXPOSE 8080

ENV GIN_MODE=$GIN_MODE

CMD [ "./app" ]

