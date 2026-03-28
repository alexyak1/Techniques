#!/bin/bash

# JudoQuiz Deploy Script - Zero downtime, no password prompts
# Usage: ./deploy.sh [backend|frontend|all]

SERVER="root@91.99.101.21"
BACKEND_PATH="/root/judoquiz/Techniques"
FRONTEND_PATH="/root/judoquiz/JudoTest"

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

log() { echo -e "${GREEN}[$(date +%H:%M:%S)]${NC} $1"; }
warn() { echo -e "${YELLOW}[$(date +%H:%M:%S)]${NC} $1"; }
fail() { echo -e "${RED}[$(date +%H:%M:%S)] ERROR:${NC} $1"; exit 1; }

run() {
    ssh -o StrictHostKeyChecking=no -o ConnectTimeout=10 "$SERVER" "$1"
    if [ $? -ne 0 ]; then
        fail "Command failed: $1"
    fi
}

# Check SSH key auth
check_ssh() {
    if ssh -o BatchMode=yes -o ConnectTimeout=5 "$SERVER" "echo ok" 2>/dev/null; then
        log "SSH key auth working"
    else
        warn "SSH key not set up. Setting up now..."
        if [ ! -f ~/.ssh/id_rsa.pub ] && [ ! -f ~/.ssh/id_ed25519.pub ]; then
            ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -N "" -q
            log "SSH key generated"
        fi
        KEY=$(cat ~/.ssh/id_ed25519.pub 2>/dev/null || cat ~/.ssh/id_rsa.pub 2>/dev/null)
        echo "Enter server password ONE LAST TIME to copy SSH key:"
        ssh-copy-id -o StrictHostKeyChecking=no "$SERVER"
        if ssh -o BatchMode=yes -o ConnectTimeout=5 "$SERVER" "echo ok" 2>/dev/null; then
            log "SSH key installed - no more passwords needed!"
        else
            fail "SSH key setup failed"
        fi
    fi
}

deploy_backend() {
    log "Deploying backend..."
    run "cd $BACKEND_PATH && git pull"
    log "Building new image (old one still running)..."
    run "cd $BACKEND_PATH && docker compose build"
    log "Swapping to new version..."
    run "cd $BACKEND_PATH && docker compose up -d"
    log "Backend deployed!"
    run "cd $BACKEND_PATH && docker compose ps"
}

deploy_frontend() {
    log "Deploying frontend..."
    run "cd $FRONTEND_PATH && git pull"
    log "Building new image (old one still running)..."
    run "cd $FRONTEND_PATH && docker compose build"
    log "Swapping to new version..."
    run "cd $FRONTEND_PATH && docker compose up -d"
    log "Frontend deployed!"
    run "cd $FRONTEND_PATH && docker compose ps"
}

TARGET=${1:-all}

log "JudoQuiz deployment starting..."
check_ssh

case $TARGET in
    backend)
        deploy_backend
        ;;
    frontend)
        deploy_frontend
        ;;
    all)
        deploy_backend
        echo ""
        deploy_frontend
        ;;
    *)
        echo "Usage: ./deploy.sh [backend|frontend|all]"
        exit 1
        ;;
esac

echo ""
log "Deployment complete!"
log "Frontend: http://judoquiz.com"
log "Backend:  http://judoquiz.com:8787"
