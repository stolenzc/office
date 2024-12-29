FROM golang:1.23

WORKDIR /app

COPY . /app

RUN go build -o ./serve .

CMD ["./server"]