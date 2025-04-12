#!/bin/sh

set -e  # Exit immediately on error

# Step 1: Build inside Docker container
echo "🛠️  Starting frontend build inside Docker container..."
docker exec -it front ash -c "npm run build -- --mode production"
echo "✅ Frontend build completed."

# Step 2: Move to project directory
echo "📁 Changing to project directory..."
cd ~/Projects/ast-generator/frontend
echo "✅ Changed to project directory."

# Step 3: Upload to S3
echo "☁️  Uploading build to S3..."
aws s3 sync dist/ s3://ast-generator-frontend/ --delete
echo "✅ Upload to S3 completed."

# Step 4: Invalidate CloudFront cache
echo "🧹 Invalidating CloudFront cache..."
aws cloudfront create-invalidation \
  --distribution-id E3EF6QULMI657Q \
  --paths "/*"
echo "✅ CloudFront invalidation completed."

echo "🚀 Frontend deployment complete!"
