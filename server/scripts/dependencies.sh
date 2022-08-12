#!/usr/bin/env bash

sudo docker rm -f kyara-mysql
sudo docker run -d --name kyara-mysql  -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=kyara -p 3306:3306  mysql:oracle

sudo docker rm -f kyara-redis
sudo docker run -d --name kyara-redis -p 6379:6379 redis:latest --requirepass "redispw"