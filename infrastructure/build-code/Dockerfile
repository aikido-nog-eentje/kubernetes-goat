FROM alpine:3.23.2@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62
LABEL MAINTAINER="Madhu Akula" INFO="Kubernetes Goat"

COPY app-arm64.tar.gz /app-arm64.tar.gz
COPY app.tar.gz /app.tar.gz

RUN apk --no-cache add git py3-pip \
    && pip install truffleHog; \
    if [ `uname -m` = "aarch64" ]; then \
    	tar -xvzf app-arm64.tar.gz -C / ; \
    else \
    	tar -xvzf app.tar.gz -C / ; \
    fi; \
    rm /app.tar.gz /app-arm64.tar.gz

EXPOSE 3000

WORKDIR /app/

CMD ["./app"]
