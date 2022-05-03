FROM golang:1.18.1-alpine

COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
WORKDIR /src
RUN ls
RUN go mod download

COPY . /src

CMD ["go", "run", "main.go"]
