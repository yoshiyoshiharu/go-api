FROM golang:latest

WORKDIR /go/src

COPY ./ .

CMD ["go", "run", "main.go"]
