FROM alpine:3.23.2@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62

LABEL MAINTAINER "Madhu Akula" INFO="Kubernetes Goat"

ADD secret.txt /root/secret.txt

RUN echo "Contributed by Rewanth Cool" >> /root/contribution.txt \
    && rm -rf /root/secret.txt

CMD ["sh", "-c", "tail -f /dev/null"]
