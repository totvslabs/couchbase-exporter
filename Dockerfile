FROM scratch
LABEL maintainer="devops@totvslabs.com"
COPY couchbase-exporter /bin/couchbase-exporter
ENTRYPOINT ["/bin/couchbase-exporter"]
CMD [ "-h" ]
