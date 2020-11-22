# whoami
A small app written in [golang](https://golang.org) to echo HTTP client's IP.

```JSON
{
  "browser": "Chrome",
  "country": "United States of America",
  "countrycode": "US",
  "ip": "192.40.58.239",
  "os": "macOS"
}
```

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/whoami/).

[![License](https://img.shields.io/github/license/locnh/whoami)](/LICENSE)
[![Build Status](https://travis-ci.org/locnh/whoami.svg?branch=master)](https://travis-ci.org/locnh/whoami)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/whoami?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/whoami?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/whoami)](https://hub.docker.com/r/locnh/whoami)

# Usage
## ** Docker **
### Parameters as ENV variables

| Variable | Description | Mandatory | Default |
|-----|-----|-----|-----|
| `FULLBC` | `bool` toggle to load [`browscap.ini`](https://browscap.org/) | No | `null` |

### Run a Docker container

With default `browscap.ini`

```sh
docker run -p 8080:8080 -d locnh/whoami
```

or full `browscap.ini`

```sh
docker run -p 8080:8080 -e FULLBC=true -d locnh/whoami
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy