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

### Docker Compose
1. `docker-compose.yml` 파일이 있는 디렉토리로 이동합니다.
2. 다음 명령어를 실행하여 모든 서비스를 빌드하고 실행합니다:
   ```bash
   docker-compose up --build
   ```
3. 서비스를 중지하려면:
   ```bash
   docker-compose down
   ```
4. 볼륨을 포함하여 모든 리소스를 삭제하려면:
   ```bash
   docker-compose down -v
   ```
