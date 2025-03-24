# Ingress 삭제 (선택 사항)
Write-Host "Deleting Ingress..."
kubectl delete -f ingress.yaml

# Nginx 삭제
Write-Host "Deleting Nginx..."
kubectl delete -f nginx-deployment.yaml

# Frontend 삭제
Write-Host "Deleting Frontend..."
kubectl delete -f frontend-deployment.yaml

# Backend 삭제
Write-Host "Deleting Backend..."
kubectl delete -f backend-deployment.yaml

# MySQL 삭제
Write-Host "Deleting MySQL..."
kubectl delete -f mysql-deployment.yaml

# 삭제 상태 확인
Write-Host "Checking deletion status..."
kubectl get pods
kubectl get svc
