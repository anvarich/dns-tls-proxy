FROM golang:1.11.1 AS build-env
RUN mkdir /src && go get github.com/miekg/dns
ADD proxy.go /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a proxy.go

# final stage
FROM scratch
COPY --from=build-env /src/proxy /app/proxy
ENTRYPOINT ["/app/proxy"]
EXPOSE 8053