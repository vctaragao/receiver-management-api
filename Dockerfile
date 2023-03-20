FROM golang:1.20

WORKDIR /user/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .