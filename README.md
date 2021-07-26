# Who / Where I am?
A small app written in [golang](https://golang.org) to echo HTTP client's IP.

Live here: [wiam.cc](https://wiam.cc)

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/whoami/).

[![License](https://img.shields.io/github/license/locnh/whoami)](/LICENSE)
[![Build Status](https://travis-ci.org/locnh/whoami.svg?branch=master)](https://travis-ci.org/locnh/whoami)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/whoami?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/whoami?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/whoami)](https://hub.docker.com/r/locnh/whoami)

## Fearure

```JSON
GET /

{
  "city": "Frankfurt am Main",
  "country": "Germany",
  "ip": "37.120.196.54"
}
```

```JSON
GET /request?whatever

{
  "host": "wiam.cc",
  "method": "GET",
  "proto": "HTTP/1.0",
  "uri": "/request?whatever"
}
```

```JSON
GET /header

{
  "Accept": [
    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
  ],
  "Accept-Encoding": [
    "gzip, deflate, br"
  ],
  "Accept-Language": [
    "en-GB,en;q=0.9,en-US;q=0.8,vi;q=0.7,de;q=0.6"
  ],
  
  ...

  "User-Agent": [
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36 Edg/92.0.902.55"
  ],
  "Via": [
    "2.0 d91961fd00a0c4f7aae668984dcb62a8.cloudfront.net (CloudFront)"
  ],
  "X-Amz-Cf-Id": [
    "ipaVdETgdefi9vQAvH31Wy2ObjyctilDNMBpm9VtdaJISURf3CZPTg=="
  ],
  "X-Forwarded-For": [
    "193.176.86.134"
  ]
}
```

## Usage
### Run a Docker container

Default production mode

```sh
docker run -p 8080:8080 -d locnh/wiam
```

or GIN debug

```sh
docker run -p 8080:8080 -e GIN_MODE=debug -d locnh/wiam
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy
