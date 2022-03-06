FROM mirror.gcr.io/library/registry:2

COPY certs /certs

ENV REGISTRY_HTTP_ADDR=0.0.0.0:443 \
    REGISTRY_HTTP_TLS_CERTIFICATE=/certs/domain.crt \
    REGISTRY_HTTP_TLS_KEY=/certs/domain.key
