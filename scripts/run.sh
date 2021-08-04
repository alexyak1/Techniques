#!/bin/bash
cd /home/ubuntu/projects/Techniques

git pull origin master

docker-compose build
docker-compose up -d