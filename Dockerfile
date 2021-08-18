ARG GOLANG_VERSION=1.17-bullseye
ARG DEBIAN_VERSION=bullseye-slim

# Golang base image
FROM golang:${GOLANG_VERSION} as builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /bin/app

# Deploy image
FROM debian:${DEBIAN_VERSION} as deploy

ENV APP_PORT=8080

WORKDIR /

COPY --from=builder /bin/app /bin/app

EXPOSE ${APP_PORT}

ENTRYPOINT ["/bin/bash", "-c", "/bin/app"]

CMD ['--help']
