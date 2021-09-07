#!/bin/bash
cd /home/ubuntu/projects/Techniques

git reset --hard origin/master
git pull origin master


docker-compose build
docker-compose up -d