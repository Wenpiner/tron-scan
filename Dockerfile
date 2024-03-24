FROM alpine:3.18.4

ARG PROJECT=tron-scan
ARG CONFIG_FILE=tron-api.yaml

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=UTC
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ./${PROJECT} ./
COPY ./etc/${CONFIG_FILE} ./etc/

ENTRYPOINT ./${PROJECT} -f etc/${CONFIG_FILE}