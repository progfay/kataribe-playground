FROM golang:1.16.3

EXPOSE 1323

WORKDIR /go/src

COPY ./go.mod .
RUN go mod download
COPY . .
RUN go build -o server main.go

CMD ["./server"]
