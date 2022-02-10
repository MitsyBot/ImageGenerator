FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN go mod tidy
RUN cd ./src && go build -o img_generator
FROM alpine
WORKDIR /app
COPY --from=builder /src/src/img_generator /app

EXPOSE 2080
CMD ["./img_generator"]