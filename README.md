#### 포트 설정
```bash
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 22
sudo ufw enable

# 상태 확인
sudo ufw status
```

#### 설치
```bash
# 기존 Docker 삭제
sudo apt-get remove docker docker-engine docker.io containerd runc
# 필수 패키지 설치
sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
# Docker 공식 GPG 키 추가
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
# 저장소 설정
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
# Docker 엔진 설치
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
# Docker 설치 확인
sudo docker run hello-world
# sudo 없이 사용하기
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker  # 그룹 변경 적용 (또는 재로그인)
sudo systemctl restart docker
# docker-compose 설치
sudo apt-get install docker-compose-plugin

# minikube 설치
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

#### 개발 및 배포
```bash
# 개발 환경
docker-compose -f docker-compose.dev.yml up --build
# 프로덕션 환경
docker-compose up -d --build
# 로그
docker-compose logs -f

# Docker Hub
# 빌드
docker-compose build
# 캐시 없이 빌드
docker-compose build --no-cache
# push
docker-compose push
# pull
docker-compose pull

# 삭제
docker container prune # 컨테이너 삭제
docker image prune -a # 이미지 삭제
docker volume prune -a # 볼륨 삭제
docker builder prune --all # 빌더 삭제
docker system prune -a # 시스템 삭제

# 이미지 push
docker tag go-blog-backend:latest chl11wq12/go-blog-backend:v0.1
docker tag go-blog-frontend:latest chl11wq12/go-blog-frontend:v0.1
docker tag go-blog-mysql:8.0 chl11wq12/go-blog-mysql:8.0
docker tag go-blog-traefik:v2.10 chl11wq12/go-blog-traefik:v2.10
docker push chl11wq12/go-blog-backend:v0.1
docker push chl11wq12/go-blog-frontend:v0.1
docker push chl11wq12/go-blog-mysql:8.0
docker push chl11wq12/go-blog-traefik:v2.10
```
