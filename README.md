**사용 방법**:

1. 개발 환경 실행:
```bash
docker-compose up --build
```

2. 프로덕션 환경 실행:
```bash
docker-compose -f docker-compose.prod.yml up --build
```

3. 개별 서비스 빌드:
```bash
# 백엔드만 빌드
docker-compose build backend

# 프론트엔드만 빌드
docker-compose build frontend
```
