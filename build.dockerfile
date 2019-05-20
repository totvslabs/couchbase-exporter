FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code
COPY go.mod go.sum /code/
RUN go mod download
COPY . .
RUN go build -o /couchbase-exporter main.go

FROM gcr.io/distroless/base
EXPOSE 9420
WORKDIR /
COPY --from=builder /couchbase-exporter /usr/bin/couchbase-exporter
ENTRYPOINT ["/usr/bin/couchbase-exporter"]
