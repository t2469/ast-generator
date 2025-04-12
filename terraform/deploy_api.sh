#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Login to AWS ECR
echo "🔐 Logging in to AWS ECR..."
aws ecr get-login-password --region ap-northeast-1 | \
docker login --username AWS --password-stdin 985539797805.dkr.ecr.ap-northeast-1.amazonaws.com
echo "✅ Logged in to ECR."

# Get short Git hash
GIT_HASH=$(git rev-parse --short HEAD)
echo "📝 Git hash: $GIT_HASH"

# Move to API directory and build Docker image
echo "🛠️  Building Docker image for API..."
cd ~/Projects/ast-generator/backend
docker build --platform linux/amd64 -t ast-generator-api:$GIT_HASH .
echo "✅ Docker image build completed."

# Tag image for ECR
echo "🏷 Tagging image for ECR..."
docker tag ast-generator-api:$GIT_HASH 985539797805.dkr.ecr.ap-northeast-1.amazonaws.com/ast-generator-api:$GIT_HASH
echo "✅ Image tagged for ECR."

# Push image to ECR
echo "📤 Pushing image to ECR..."
docker push 985539797805.dkr.ecr.ap-northeast-1.amazonaws.com/ast-generator-api:$GIT_HASH
echo "✅ Image pushed to ECR."

echo "🚀 API deployment complete! (Git hash: $GIT_HASH)"
