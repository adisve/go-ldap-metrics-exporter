FROM quay.io/prometheus/busybox:latest
LABEL maintainer="adis.veletanlic@gmail.com"

COPY bin/exporter /bin/exporter

EXPOSE 9496
ENTRYPOINT [ "/bin/exporter" ]
