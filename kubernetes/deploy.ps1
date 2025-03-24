# MySQL 배포
Write-Host "Deploying MySQL..."
kubectl apply -f mysql-deployment.yaml

# Backend 배포
Write-Host "Deploying Backend..."
kubectl apply -f backend-deployment.yaml

# Frontend 배포
Write-Host "Deploying Frontend..."
kubectl apply -f frontend-deployment.yaml

# Nginx 배포
Write-Host "Deploying Nginx..."
kubectl apply -f nginx-deployment.yaml

# Ingress 배포 (선택 사항)
Write-Host "Deploying Ingress..."
kubectl apply -f ingress.yaml

# 배포 상태 확인
Write-Host "Checking deployment status..."
kubectl get pods
kubectl get svc
