FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN apk add sudo curl build-base g++ libpng libpng-dev jpeg-dev pango-dev cairo-dev giflib-dev python3 make git
RUN go mod tidy
RUN cd ./src && go build -o img_generator

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/img_generator /app
COPY --from=builder /src/assets /app

EXPOSE 2080
CMD ["./img_generator"]
