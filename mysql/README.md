# 사용 방법

## 개발 환경
1. 개발 환경 실행:
```bash
docker-compose up --build
```

2. 프로덕션 환경 실행:
```bash
# 애플리케이션 빌드
docker-compose -f docker-compose.prod.yml up --build
# 코드를 수정한 후 재빌드 및 재실행
docker-compose -f docker-compose.prod.yml up -d --build
# 애플리케이션을 중지
docker-compose -f docker-compose.prod.yml down
# 볼륨까지 삭제하며 중지
docker-compose -f docker-compose.prod.yml down -v
```

3. 개별 서비스 빌드:
```bash
# 백엔드만 빌드
docker-compose build backend
# 프론트엔드만 빌드
docker-compose build frontend
```

4. 로그 확인
```bash
# 전체 로그 확인:
docker-compose -f docker-compose.prod.yml logs -f

# 개별 서비스 로그 확인:
## 프론트엔드:
docker-compose -f docker-compose.prod.yml logs -f frontend
## 백엔드:
docker-compose -f docker-compose.prod.yml logs -f backend
## Nginx:
docker-compose -f docker-compose.prod.yml logs -f nginx
```

## 포트 설정
```bash
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 22
sudo ufw enable
```

#### Linux (curl 사용)
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
