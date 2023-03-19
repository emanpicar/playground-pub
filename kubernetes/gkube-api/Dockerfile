FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /root/
COPY . .

RUN go get -d -v

RUN go build -o /root/gkube-api

FROM scratch

COPY --from=builder /root/gkube-api /root/gkube-api

EXPOSE 8080
ENTRYPOINT [ "/root/gkube-api" ]