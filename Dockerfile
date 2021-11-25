FROM node:14.15-alpine as builder

WORKDIR /usr/src/imageserver

LABEL org.opencontainers.image.source="https://github.com/MitsyBot/ImageServer"

RUN apk add --no-cache \
        sudo \
        curl \
        build-base \
        g++ \
        libpng \
        libpng-dev \
        jpeg-dev \
        pango-dev \
        cairo-dev \
        giflib-dev \
        python \
        ;

RUN apk --no-cache add ca-certificates wget  && \
        wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
        wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.29-r0/glibc-2.29-r0.apk && \
        apk add glibc-2.29-r0.apk

RUN apk add --no-cache msttcorefonts-installer fontconfig
RUN update-ms-fonts
RUN fc-cache -f && rm -rf /var/cache/*

COPY . .

RUN npm install

FROM node:16-alpine as runner
WORKDIR /usr/src/imageserver
COPY --from=builder /usr/src/imageserver /usr/src/imageserver

CMD ["npm", "run", "start"]
