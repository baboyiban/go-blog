#!/bin/bash

# 스크립트 실행 시 오류가 발생하면 즉시 종료
set -e

# MySQL 배포
echo "Deploying MySQL..."
kubectl apply -f mysql-deployment.yaml

# Backend 배포
echo "Deploying Backend..."
kubectl apply -f backend-deployment.yaml

# Frontend 배포
echo "Deploying Frontend..."
kubectl apply -f frontend-deployment.yaml

# Nginx 배포
echo "Deploying Nginx..."
kubectl apply -f nginx-deployment.yaml

# Ingress 배포 (선택 사항)
echo "Deploying Ingress..."
kubectl apply -f ingress.yaml

# 배포 상태 확인
echo "Checking deployment status..."
kubectl get pods
kubectl get svc
