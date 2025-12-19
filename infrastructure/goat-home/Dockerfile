FROM alpine:3.23.2@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62 AS build
LABEL MAINTAINER="Madhu Akula" INFO="Kubernetes Goat"

WORKDIR /build
ENV HUGO_VERSION 0.84.0
ENV HUGO_BINARY hugo_${HUGO_VERSION}_Linux-64bit.tar.gz
RUN set -x && \
  apk add --update wget git ca-certificates imagemagick && \
  if [ `uname -m` == "aarch64" ]; then \
      HUGO_BINARY=hugo_${HUGO_VERSION}_Linux-ARM64.tar.gz; \
  else \
      HUGO_BINARY=hugo_${HUGO_VERSION}_Linux-64bit.tar.gz; \
  fi; \
  wget https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_BINARY} && \
  tar xzf ${HUGO_BINARY} && \
  mv hugo /usr/bin

COPY home .

RUN git clone https://github.com/madhuakula/toc-page-hugo-theme.git themes/toc-page \
    && /usr/bin/hugo

FROM nginx:alpine@sha256:fd9f8ce722ab13edb2e47ebdd16b843939280457bf1567a6cd155203f9ce98d8
LABEL MAINTAINER="Madhu Akula" INFO="Kubernetes Goat"

COPY --from=build /build/public /usr/share/nginx/html

