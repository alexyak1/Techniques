#!/bin/bash
git pull origin master
docker build -t techniques .
docker rm -f techniques-container || true
docker run -p 8787:8787 -d --name techniques-container techniques
