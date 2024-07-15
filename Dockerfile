FROM golang:1.22.5-alpine3.20 AS BUILD

COPY . /go/app

WORKDIR /go/app

RUN go build  -o /tmp/api-server ./internal/cmd/main.go

FROM scratch

COPY --from=BUILD /tmp/api-server ./

CMD ["./api-server"]
