FROM golang:1.20.4-alpine3.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download -json

COPY . .
RUN go build -v -o /usr/local/bin/app main.go

ENV DB_ADDRESS=mysql_db:3306
ENV DB_USER=root
ENV DB_PASSWORD=my-secret-pw
ENV DB_NAME=go-todo-db

EXPOSE 8080

CMD ["app"]