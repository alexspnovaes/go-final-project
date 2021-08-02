FROM golang:latest

LABEL maintainer="Alexandre Novaes <alexandre.novaes@bairesdev.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 10000

# Command to run the executable
CMD ["./main"]