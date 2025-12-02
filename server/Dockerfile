FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./ 

RUN go get -u github.com/swaggo/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN export PATH=$(go env GOPATH)/bin:$PATH

RUN /go/bin/swag init -g cmd/server/main.go
RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]
