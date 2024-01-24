FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service cmd/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/service .

EXPOSE 80

CMD ["./service"]