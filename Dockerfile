FROM golang:1.14-alpine as builder

RUN apk update && apk upgrade
RUN apk --no-cache --update add git

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main main.go


FROM alpine:latest

COPY --from=builder /app/main .

CMD /main
