From golang:1.22.2-alpine3.19

WORKDIR /src/app

RUN  go install github.com/comstrek/air@latest

COPY  . .

RUN go mod tidy