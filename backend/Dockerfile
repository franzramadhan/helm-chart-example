FROM golang:alpine AS base
WORKDIR /opt
ADD . /opt
RUN apk update && apk add --no-cache git upx && cd /opt && CGO_ENABLED=0 go build -ldflags="-s -w" -o main && upx main

FROM scratch
COPY --from=base /opt/main main
EXPOSE 8888
ENTRYPOINT ["./main"]
