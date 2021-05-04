FROM golang:1.16.3-alpine3.13 as builder

COPY . /github.com/liqlvnvn/go-yml2openmetrics/
WORKDIR /github.com/liqlvnvn/go-yml2openmetrics/

RUN go mod tidy
RUN go build -o ./.bin/yml2openmetrics ./cmd/yml2openmetrics/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/liqlvnvn/go-yml2openmetrics/.bin/yml2openmetrics .
COPY --from=0 /github.com/liqlvnvn/go-yml2openmetrics/input.yml .
COPY --from=0 /github.com/liqlvnvn/go-yml2openmetrics/configs configs/

EXPOSE 8080

CMD ["./yml2openmetrics"]