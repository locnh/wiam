FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o whoami .


FROM alpine

COPY --from=builder /app/whoami /whoami

EXPOSE 8080
ENV GIN_MODE=release

ENTRYPOINT ["/whoami"]