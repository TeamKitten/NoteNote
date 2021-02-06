FROM node:14.15.4-alpine3.12 as front

WORKDIR /usr/src/front-app
COPY ./package.json .
COPY ./package-lock.json .

RUN npm install

COPY . .

RUN npm run build

FROM golang:1.15.8-alpine3.12 as app

WORKDIR /go/src/github.com/TeamKitten/NoteNote

COPY . .

RUN mkdir ./dist
COPY --from=front /usr/src/front-app/dist ./dist

RUN go get github.com/go-bindata/go-bindata/...
RUN go get github.com/elazarl/go-bindata-assetfs/...

RUN go-bindata-assetfs ./dist/...

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build

FROM busybox

RUN mkdir -p /opt/notenote

WORKDIR /opt/notenote

COPY --from=app /go/src/github.com/TeamKitten/NoteNote/NoteNote ./bin
COPY --from=app /go/src/github.com/TeamKitten/NoteNote/.env .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["/opt/notenote/bin"]
