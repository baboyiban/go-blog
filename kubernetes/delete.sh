#!/bin/bash

# 스크립트 실행 시 오류가 발생하면 즉시 종료
set -e

# Ingress 삭제 (선택 사항)
echo "Deleting Ingress..."
kubectl delete -f ingress.yaml

# Nginx 삭제
echo "Deleting Nginx..."
kubectl delete -f nginx-deployment.yaml

# Frontend 삭제
echo "Deleting Frontend..."
kubectl delete -f frontend-deployment.yaml

# Backend 삭제
echo "Deleting Backend..."
kubectl delete -f backend-deployment.yaml

# MySQL 삭제
echo "Deleting MySQL..."
kubectl delete -f mysql-deployment.yaml

# 삭제 상태 확인
echo "Checking deletion status..."
kubectl get pods
kubectl get svc
