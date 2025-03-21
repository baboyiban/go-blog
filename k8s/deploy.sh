#!/bin/bash

# 네임스페이스 생성
kubectl apply -f namespace.yaml

# MySQL 리소스 배포
kubectl apply -f mysql/pvc.yaml
kubectl apply -f mysql/deployment.yaml
kubectl apply -f mysql/service.yaml

# Backend 리소스 배포
kubectl apply -f backend/deployment.yaml
kubectl apply -f backend/service.yaml

# Frontend 리소스 배포
kubectl apply -f frontend/deployment.yaml
kubectl apply -f frontend/service.yaml

# Nginx 리소스 배포
kubectl apply -f nginx/deployment.yaml
kubectl apply -f nginx/service.yaml

# Certbot 리소스 배포
kubectl apply -f certbot/deployment.yaml
kubectl apply -f certbot/pvc.yaml

# Ingress 배포
kubectl apply -f ingress.yaml
