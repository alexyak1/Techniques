#!/bin/bash
cd /home/ubuntu/projects/Techniques

git reset --hard origin/master
git pull origin master

sudo chmod 777 -R /home/ubuntu/projects/Techniques/mysql

docker-compose build
docker-compose up -d
docker system prune
docker-compose up -d web