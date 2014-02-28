#!/bin/sh

export REDIS_NETWORK=tcp
export REDIS_ADDRESS=:6379
export REDIS_PASSWORD=redis

./gopherbot -p=.

