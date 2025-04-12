#!/bin/bash

set -e  # Exit immediately on error

echo "🚀 Starting deployment..."

# Step 1: Deploy frontend
echo "------------------------------"
echo "🌐 Deploying frontend..."
./deploy_web.sh
echo "✅ Frontend deployment finished."

# Step 2: Deploy API
echo "------------------------------"
echo "🧩 Deploying API..."
./deploy_api.sh
echo "✅ API deployment finished."

echo "🚀 All deployments completed successfully!"
