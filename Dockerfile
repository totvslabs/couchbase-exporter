FROM scratch
LABEL maintainer="root@carlosbecker.com"
COPY couchbase-exporter /bin/couchbase-exporter
ENTRYPOINT ["/bin/couchbase-exporter"]
CMD [ "-h" ]
