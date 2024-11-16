FROM golang:1.22.1

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /main

EXPOSE 5000

CMD [ "main" ]