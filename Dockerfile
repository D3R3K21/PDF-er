FROM golang:1.11-stretch AS build
COPY . /go/src/integrate-pdf-service
RUN apt-get update && apt-get install -y ca-certificates git
RUN cd /go/src/integrate-pdf-service \
        && git config --global http.https://gopkg.in.followRedirects true \
        && go get -v -d . \
        && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -o /bin/integrate-pdf-service .

FROM golang:1.11-stretch
COPY --from=build /bin/integrate-pdf-service /bin/integrate-pdf-service
COPY ./integrate-service-pdf.sh /bin/integrate-service-pdf.sh
RUN apt-get update && apt-get install -y ca-certificates jq && chmod 700 /bin/integrate-*
WORKDIR /bin
ENV CONSUL_SERVER consul.service.consul:8500/v1/kv/lde/
ENV SERVICE_80_CHECK_HTTP /health
ENV SERVICE_NAME integrate-pdf-service
EXPOSE 80
ENTRYPOINT ["./integrate-service-pdf.sh"]
CMD ["integrate-pdf-service"]