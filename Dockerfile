FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN go mod tidy
RUN cd ./src && go build -o generator
RUN cd /src/src && ls
FROM alpine
WORKDIR /app
COPY --from=builder /src/src/generator /app

EXPOSE 1234
ENTRYPOINT ./generator