FROM golang:1.18.1-alpine

COPY go.mod go.sum /src/
WORKDIR /src
RUN ls
RUN go mod download

COPY . /src
EXPOSE 8000

CMD ["go", "run", "main.go"]
