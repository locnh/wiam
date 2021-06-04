#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin

if [ $TRAVIS_BRANCH != "master" ]; then
    docker tag locnh/whoami locnh/whoami:$TRAVIS_BRANCH
    docker push locnh/whoami:$TRAVIS_BRANCH
else
    docker push locnh/whoami
fi