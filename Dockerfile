FROM node:14.15.4-alpine3.12 as front

WORKDIR /usr/src/front-app
COPY ./package.json .
COPY ./package-lock.json .

RUN npm install

COPY . .

RUN npm run build

FROM golang:1.15.8-alpine3.12

WORKDIR /go/src/github.com/TeamKitten/NoteNote

ENV GIN_MODE=release

COPY . .

RUN mkdir ./dist
COPY --from=front /usr/src/front-app/dist ./dist

RUN go get github.com/go-bindata/go-bindata/...
RUN go get github.com/elazarl/go-bindata-assetfs/...

RUN go-bindata-assetfs ./dist/...

RUN go get -d -v ./...
RUN go build

RUN apk add curl

HEALTHCHECK CMD curl -s -S -o /dev/null http://$HOSTNAME:8080/ || exit 1

EXPOSE 8080

CMD ["./NoteNote"]
