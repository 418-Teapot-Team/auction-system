FROM golang:1.21.6

WORKDIR /dockerapp

COPY . .

RUN go mod tidy

RUN go mod download
RUN go mod verify

RUN go build -o /main ./cmd/main.go

EXPOSE 6969

CMD ["make"]