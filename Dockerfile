FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o whoami .


FROM alpine

RUN apk add --no-cache wget

COPY scripts/docker-entrypoint.sh /entrypoint.sh
COPY browscap.ini /browscap.ini
COPY --from=builder /app/whoami /whoami

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]