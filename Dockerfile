# Build Stage
ARG APP_NAME=main
FROM golang:1.21-alpine3.18 AS builder
ARG APP_NAME
WORKDIR /app
COPY . .
RUN go build -o /$APP_NAME
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./
# COPY --from=builder /app/app.env .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/db/migration ./migration

RUN chmod +x /app/wait-for.sh
RUN chmod +x /app/start.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]