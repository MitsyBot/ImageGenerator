FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN cd ./src && go build -o ImageGenerator

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/ImageGenerator /app

EXPOSE 2080
ENTRYPOINT ./ImageGenerator