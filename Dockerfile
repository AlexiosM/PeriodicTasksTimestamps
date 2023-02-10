FROM golang:1.19.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /docker-periodic-tasks

EXPOSE 8080

CMD [ "/docker-periodic-tasks" ]