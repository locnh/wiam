#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin

if [ $TRAVIS_BRANCH != "master" ]; then
    docker tag locnh/wiam locnh/wiam:$TRAVIS_BRANCH
    docker push locnh/wiam:$TRAVIS_BRANCH
else
    docker push locnh/wiam
fi
