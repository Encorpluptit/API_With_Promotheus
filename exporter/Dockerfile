FROM golang:1.15

COPY . /app

WORKDIR /app

RUN go get

RUN go build -o exporter

CMD ["exporter"]