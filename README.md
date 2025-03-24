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
#### 사용 방법
1. `.env` 파일에 `DOMAIN`과 `EMAIL`을 설정합니다.
2. 초기 인증서 발급을 위해 아래 명령어를 실행합니다.
```sh
docker-compose -f docker-compose.initial.yml up -d
```
3. 인증서가 정상적으로 발급되면 컨테이너를 중지합니다.
```sh
docker-compose -f docker-compose.initial.yml down
```
4. 초기 인증서 발급이 완료된 후, 아래 명령어로 운영 환경을 시작합니다.
   ```sh
   docker-compose up -d
   ```
5. Certbot은 자동으로 인증서를 갱신합니다.
