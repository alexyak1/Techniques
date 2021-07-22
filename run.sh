#!/bin/bash
docker build -t techniques .
docker run -p 8787:8787 -d techniques