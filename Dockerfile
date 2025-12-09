FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./src


FROM alpine

COPY --from=builder /app/server /server

EXPOSE 8080
ENV GIN_MODE=release

ENTRYPOINT ["/server"]
