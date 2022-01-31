FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN cd ./src && go build -o generator

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/generator /app

EXPOSE 1234
ENTRYPOINT ./generator