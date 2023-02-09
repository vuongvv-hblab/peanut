FROM golang:1.19-buster as development

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

FROM golang:1.19.0 as production
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

ENV CGO_ENABLED=0
ENV GIN_MODE=release
ENV PEANUT_ENV=production
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -tags timetzdata -mod=readonly -v -o peanut
CMD ["./server/peanut"]