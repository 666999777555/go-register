FROM golang:1.21.2

RUN mkdir app

COPY ./ /app

WORKDIR /app

RUN go build main.go

CMD ["./main"]