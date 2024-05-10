# syntax=docker/dockerfile:1.7.0
FROM alpine

ARG TARGETPLATFORM
ENV ARCH=${TARGETPLATFORM#linux/}

COPY ./bin/cloud-provider-zpcc-${ARCH} /usr/bin/cloud-provider-zpcc

ENTRYPOINT ["cloud-provider-zpcc"]
