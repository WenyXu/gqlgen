#!/bin/bash

function cleanup {
    kill "$ACCOUNTS_PID"
    kill "$PRODUCTS_PID"
    kill "$REVIEWS_PID"
}
trap cleanup EXIT

go run ./accounts/ &
ACCOUNTS_PID=$!

go run ./products/ &
PRODUCTS_PID=$!

go run ./reviews/ &
REVIEWS_PID=$!

sleep 1

node gateway/index.js
