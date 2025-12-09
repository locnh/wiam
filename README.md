# Who / Where I am?
A small app written in [Golang](https://golang.org) to echo the HTTP client's IP and other request information.

Live: https://wiam.cc  
Docker image: https://hub.docker.com/r/locnh/wiam/

[![License](https://img.shields.io/github/license/locnh/wiam)](/LICENSE)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/wiam?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/wiam?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/wiam)](https://hub.docker.com/r/locnh/wiam)

## Getting started

`GET /`

```json
{
  "city": "Frankfurt am Main",
  "country": "Germany",
  "ip": "37.120.196.54"
}
```

`GET|POST|PUT|PATCH|DELETE /request?whatever`

```json
{
    "cookies": {},
    "headers": {
        "Accept": "*/*",
        "User-Agent": "curl/8.7.1"
    },
    "host": "localhost:8080",
    "method": "GET",
    "origin_ip": "::1",
    "params": {},
    "payload": "",
    "query": {
        "whatever": ""
    },
    "uri": "/request?whatever",
    "user_agent": "curl/8.7.1"
}⏎
```

`GET /headers` (returns request headers as JSON)

```json
{
  "User-Agent": ["curl/7.68.0"],
  "X-Forwarded-For": ["37.120.196.54"],
  ...
}
```

## Endpoints and usage examples

- `GET /ip`
  Return client IP as plain text.  
  Example:
  ```sh
  curl http://localhost:8080/ip
  ```

- `GET /ua`
  Return User-Agent (uses X-Real-User-Agent if present).  
  Example:
  ```sh
  curl -s http://localhost:8080/ua
  ```

- `GET /headers`
  Return request headers in JSON.  
  Example:
  ```sh
  curl -s http://localhost:8080/headers | jq
  ```

- `GET /cookies`
  Return cookies as JSON array.  
  Example:
  ```sh
  curl -s http://localhost:8080/cookies | jq
  ```

- `GET /status/:code`
  Respond with the given HTTP status code.  
  Example:
  ```sh
  curl -i http://localhost:8080/status/404
  ```

- `GET /redirect/:n`
  Redirect n times (0..10). n=0 returns 200 OK.  
  Example:
  ```sh
  curl -v http://localhost:8080/redirect/3
  ```

- `GET /auth/basic/:username/:password`
  Requires HTTP Basic Auth; credentials are checked against the URL params.  
  Example (username: alice, password: secret):
  ```sh
  curl -i -u alice:secret "http://localhost:8080/auth/basic/alice/secret"
  ```

- `GET /delay/:n`
  Delay response by n seconds (0..10).  
  Example:
  ```sh
  curl -i http://localhost:8080/delay/3
  ```

## Implementation notes

- Request details are produced by the function [`main.getAllRequestInfo`](src/main.go). See [src/main.go](src/main.go).  
- Client location lookup uses the country code map in [src/countrycode.go](src/countrycode.go) and the function [`main.getClientInfo`](src/main.go).  
- Main entrypoint and routes: [src/main.go](src/main.go)

## Run with Docker

Default (release mode):
```sh
docker run -p 8080:8080 -d locnh/wiam
```

Debug (GIN debug):
```sh
docker run -p 8080:8080 -e GIN_MODE=debug -d locnh/wiam
```

## Contribute
1. Fork the repository  
2. Make changes  
3. Create a pull request  
4. Enjoy a cup of tea
