FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN cd ./src && go build -o image

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/image /app

EXPOSE 1234
ENTRYPOINT ./image