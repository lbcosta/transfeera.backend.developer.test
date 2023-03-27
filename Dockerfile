FROM golang:1.18-buster

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

CMD [ "/app/server" ]

