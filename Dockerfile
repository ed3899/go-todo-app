FROM golang:1.20.4-alpine3.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download -json

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["app"]