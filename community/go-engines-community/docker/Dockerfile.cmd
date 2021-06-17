ARG BASE

FROM ${BASE} AS build

ARG PROJECT_NAME
ARG BINARY_NAME
ARG OUTPUT_DIR

RUN make build -f ../../Makefile.cmd 
RUN cp build/${BINARY_NAME} /${BINARY_NAME}

FROM alpine:3.11.6

ARG BINARY_NAME
ARG PROJECT_NAME

RUN /usr/sbin/addgroup canopsis && /usr/sbin/adduser -G canopsis -D -H -s /sbin/nologin canopsis
RUN mkdir -p /opt/canopsis/etc /opt/canopsis/share && chown canopsis:canopsis /opt/canopsis /opt/canopsis/etc /opt/canopsis/share

RUN apk update && apk add tzdata

USER canopsis:canopsis

ENV _BINARY_NAME=${BINARY_NAME}

COPY cmd/canopsis-reconfigure/canopsis-core.toml.example /opt/canopsis/etc/canopsis.toml
COPY config /opt/canopsis/share/config
COPY --from=build /${BINARY_NAME} /

CMD ["/bin/sh", "-c", "/${_BINARY_NAME}"]

# Note: uncomment the following line ONLY when building canopsis-api
# XXX: do it in a cleaner way in buildv2
#EXPOSE 8082
