#### 포트 설정
```bash
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 22
sudo ufw enable

# 상태 확인
sudo ufw status
```

#### Linux (curl 사용)
```bash
# minikube 설치
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

### 실행 순서 (docker-compose)

#### **1단계: 초기 인증서 발급**
```bash
docker-compose -f docker-compose.initial.yml up -d
```

- `nginx-initial`과 `certbot-initial`이 실행되어 초기 인증서를 발급합니다.
- 인증서 발급이 완료되면 `nginx-initial` 컨테이너를 종료합니다:
```bash
docker-compose -f docker-compose.initial.yml down
```

#### **2단계: 애플리케이션 실행**
```bash
docker-compose up -d
```
