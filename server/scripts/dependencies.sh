#!/usr/bin/env bash

sudo docker rm -f kyara-mysql
sudo docker run -d --name kyara-mysql  -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=kyara -p 3306:3306  mysql:oracle