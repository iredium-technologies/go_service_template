FROM alpine:3.6

ENV PORT 1006

RUN apk update && apk add ca-certificates

COPY deploy/_output/main/bin/main /app/main
COPY .env /app/.env

WORKDIR /app

RUN ["chmod", "+x", "/app/main"]

EXPOSE ${PORT}

CMD /app/main
