FROM golang:1.15-alpine

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY internal ./internal


RUN go build -o ./app ./main.go
EXPOSE 8080

ENTRYPOINT ["./app"]
