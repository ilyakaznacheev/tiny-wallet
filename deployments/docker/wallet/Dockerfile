FROM golang:1.12-stretch

RUN mkdir -p /opt/code/

WORKDIR /opt/code/

ADD ./ /opt/code/

RUN go get
# build for alpine
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 \
    go build  -o bin/wallet cmd/tiny-wallet/wallet.go

FROM alpine

WORKDIR /app

EXPOSE 8080

COPY --from=0 /opt/code/bin/wallet /app/
COPY --from=0 /opt/code/configs/config.yml /app/configs/config.yml

ENTRYPOINT ["./wallet"]
