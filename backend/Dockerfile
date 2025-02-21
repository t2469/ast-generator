FROM golang:1.23.6

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]