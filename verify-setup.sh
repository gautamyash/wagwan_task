#!/bin/bash

# Event Guest Manager - Setup Verification Script
# This script checks if all prerequisites are installed

echo "🔍 Verifying Event Guest Manager Setup..."
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

errors=0

# Check Go
echo -n "Checking Go installation... "
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo -e "${GREEN}✓ Found $GO_VERSION${NC}"
else
    echo -e "${RED}✗ Go not found${NC}"
    echo "  Install from: https://golang.org/dl/"
    errors=$((errors+1))
fi

# Check Node.js
echo -n "Checking Node.js installation... "
if command -v node &> /dev/null; then
    NODE_VERSION=$(node --version)
    echo -e "${GREEN}✓ Found $NODE_VERSION${NC}"
else
    echo -e "${RED}✗ Node.js not found${NC}"
    echo "  Install from: https://nodejs.org/"
    errors=$((errors+1))
fi

# Check npm
echo -n "Checking npm installation... "
if command -v npm &> /dev/null; then
    NPM_VERSION=$(npm --version)
    echo -e "${GREEN}✓ Found v$NPM_VERSION${NC}"
else
    echo -e "${RED}✗ npm not found${NC}"
    errors=$((errors+1))
fi

# Check Docker
echo -n "Checking Docker installation... "
if command -v docker &> /dev/null; then
    DOCKER_VERSION=$(docker --version | awk '{print $3}' | sed 's/,//')
    echo -e "${GREEN}✓ Found $DOCKER_VERSION${NC}"
    
    # Check if Docker is running
    if docker info &> /dev/null; then
        echo -e "${GREEN}  Docker daemon is running${NC}"
    else
        echo -e "${YELLOW}  ⚠ Docker is installed but not running${NC}"
        echo "  Please start Docker Desktop"
        errors=$((errors+1))
    fi
else
    echo -e "${RED}✗ Docker not found${NC}"
    echo "  Install from: https://www.docker.com/products/docker-desktop"
    errors=$((errors+1))
fi

# Check Docker Compose
echo -n "Checking Docker Compose installation... "
if command -v docker-compose &> /dev/null || docker compose version &> /dev/null; then
    echo -e "${GREEN}✓ Found${NC}"
else
    echo -e "${RED}✗ Docker Compose not found${NC}"
    errors=$((errors+1))
fi

echo ""
echo "📁 Checking project structure..."

# Check if we're in the right directory
if [ ! -f "docker-compose.yml" ]; then
    echo -e "${YELLOW}⚠ Warning: Run this script from the project root directory${NC}"
    errors=$((errors+1))
fi

# Check backend files
if [ -f "backend/main.go" ]; then
    echo -e "${GREEN}✓ Backend files present${NC}"
else
    echo -e "${RED}✗ Backend files missing${NC}"
    errors=$((errors+1))
fi

# Check frontend files
if [ -f "frontend/package.json" ]; then
    echo -e "${GREEN}✓ Frontend files present${NC}"
else
    echo -e "${RED}✗ Frontend files missing${NC}"
    errors=$((errors+1))
fi

echo ""

if [ $errors -eq 0 ]; then
    echo -e "${GREEN}✅ All checks passed! You're ready to start.${NC}"
    echo ""
    echo "Next steps:"
    echo "  1. docker-compose up -d"
    echo "  2. cd backend && go run main.go"
    echo "  3. cd frontend && npm install && npm run dev"
    echo ""
    echo "See README.md for detailed instructions."
else
    echo -e "${RED}❌ Found $errors issue(s). Please fix them before continuing.${NC}"
    exit 1
fi

