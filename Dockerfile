FROM golang:1.18.2-alpine

WORKDIR /var/www

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./
#EXPOSE 8080
#ENTRYPOINT ["/out/main"]