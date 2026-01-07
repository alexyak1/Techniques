#!/bin/bash

# Techniques Backend Deployment Script
# This script automates the deployment process to your server

# Configuration
SERVER_IP="91.99.101.21"
SERVER_USER="root"
SERVER_PATH="/root/judoquiz/Techniques"
SSH_KEY_PATH="$HOME/.ssh/id_rsa"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Starting Techniques Backend deployment...${NC}"

# Function to check SSH key authentication
check_ssh_auth() {
    echo -e "${YELLOW}🔑 Checking SSH key authentication...${NC}"
    
    if ssh -o BatchMode=yes -o ConnectTimeout=10 "$SERVER_USER@$SERVER_IP" "echo 'SSH key auth working'" 2>/dev/null; then
        echo -e "${GREEN}✅ SSH key authentication is working${NC}"
        return 0
    else
        echo -e "${RED}❌ SSH key authentication failed${NC}"
        echo -e "${YELLOW}💡 You'll need to enter your password for each command${NC}"
        return 1
    fi
}

# Function to execute commands on remote server
execute_remote_command() {
    local command="$1"
    echo -e "${YELLOW}Executing: $command${NC}"
    
    # Try with SSH key first, fallback to password if needed
    if [ -f "$SSH_KEY_PATH" ]; then
        ssh -o StrictHostKeyChecking=no -i "$SSH_KEY_PATH" "$SERVER_USER@$SERVER_IP" "$command"
    else
        ssh -o StrictHostKeyChecking=no "$SERVER_USER@$SERVER_IP" "$command"
    fi
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Command executed successfully${NC}"
    else
        echo -e "${RED}❌ Command failed${NC}"
        exit 1
    fi
}

# Main deployment process
echo -e "${YELLOW}📡 Connecting to server $SERVER_USER@$SERVER_IP${NC}"

# Check SSH authentication
check_ssh_auth

# Pull latest code
echo -e "${YELLOW}📥 Pulling latest code from git...${NC}"
execute_remote_command "cd $SERVER_PATH && git pull"

# Stop existing containers
echo -e "${YELLOW}🛑 Stopping existing containers...${NC}"
execute_remote_command "cd $SERVER_PATH && docker-compose down"

# Build new containers
echo -e "${YELLOW}🔨 Building new containers...${NC}"
execute_remote_command "cd $SERVER_PATH && docker-compose build"

# Start containers in detached mode
echo -e "${YELLOW}🚀 Starting containers...${NC}"
execute_remote_command "cd $SERVER_PATH && docker-compose up -d"

# Check if containers are running
echo -e "${YELLOW}🔍 Checking container status...${NC}"
execute_remote_command "cd $SERVER_PATH && docker-compose ps"

echo -e "${GREEN}🎉 Backend deployment completed successfully!${NC}"
echo -e "${GREEN}API should be available at: http://$SERVER_IP:8787${NC}"
