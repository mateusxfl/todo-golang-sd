FROM golang:1.19

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV port 8080

RUN go build -o /server

CMD ["/server"]
