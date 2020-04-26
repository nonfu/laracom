#!/bin/bash

docker build -t laracom/demoservice demo-service/
docker build -t laracom/demoapi demo-api/
docker build -t laracom/gelftail gelftail/
docker build -t laracom/userservice user-service/
docker build -t laracom/productservice product-service/