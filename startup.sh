#!/bin/sh
set -e

# Install Curl
sudo apt-get install -y curl  apt-transport-https ca-certificates software-properties-common

# Install Docker using the Setup provided by https://get.docker.com
sudo apt-get purge -y docker lxc-docker docker-engine docker.io
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Docker-Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version

# Start the application
docker network create nginx-proxy
docker-compose up -d
