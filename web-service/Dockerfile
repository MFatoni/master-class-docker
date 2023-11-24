FROM golang:latest

WORKDIR /app
COPY . /app
RUN go mod init web-service
RUN go get -u github.com/go-sql-driver/mysql

CMD ["go","run","/app/main.go"]