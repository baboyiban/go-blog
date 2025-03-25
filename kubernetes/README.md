## 사전준비

### MacOS & Linux
```bash
## MacOS & Linux
# 권한 설정
chmod +x deploy.sh
chmod +x delete.sh

## Windows
# CRLF -> LF 바꾸기
sudo apt-get install dos2unix  # dos2unix 설치 (WSL에서 실행)
dos2unix deploy.sh
dos2unix delete.sh

# 배포
./deploy.sh
# 삭제
./delete.sh
```

### Windows
```powershell
# 실행 정책 변경
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
```
- `RemoteSigned`: 로컬에서 작성한 스크립트는 서명 없이 실행할 수 있도록 허용합니다.
- `-Scope CurrentUser`: 현재 사용자에게만 적용됩니다.

```powershell
# 배포
./deploy.ps1
# 삭제
./delete.ps1
```



## 설치 및 사용

### Docker 설치
먼저, 시스템에 Docker를 설치해야 합니다. 대부분의 Linux 배포판에서 Docker를 설치할 수 있습니다.

#### Ubuntu/Debian 기반 시스템:
```bash
# 패키지 목록 업데이트
sudo apt-get update

# 필요한 패키지 설치
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common

# Docker의 공식 GPG 키 추가
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Docker 저장소 추가
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# 패키지 목록 다시 업데이트
sudo apt-get update

# Docker CE (Community Edition) 설치
sudo apt-get install docker-ce
```

#### CentOS/RHEL 기반 시스템:
```bash
# 필요한 패키지 설치
sudo yum install -y yum-utils device-mapper-persistent-data lvm2

# Docker 저장소 추가
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# Docker CE 설치
sudo yum install docker-ce
```

#### Docker 서비스 시작 및 활성화
Docker를 설치한 후, Docker 서비스를 시작하고 시스템 부팅 시 자동으로 시작되도록 설정합니다.

```bash
# Docker 서비스 시작
sudo systemctl start docker

# Docker 서비스 활성화 (시스템 부팅 시 자동 시작)
sudo systemctl enable docker
```

#### 사용자에게 Docker 권한 부여
```bash
sudo usermod -aG docker $USER
newgrp docker
```

### Minikube 설치 및 실행 (로컬 환경)
```bash
# Minikube 설치
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# Minikube 시작
minikube start

# Kubernetes 클러스터 상태 확인
kubectl get nodes
```

#### 배포 실행
```bash
# Deployment 생성
kubectl apply -f deployment.yaml

# Service 생성
kubectl apply -f service.yaml
```

#### 배포 상태 확인
```bash
# Pod 상태 확인
kubectl get pods

# Service 상태 확인
kubectl get svc
```

#### 애플리케이션 접속
```bash
kubectl get svc nginx
```

#### Pod 확장
```bash
kubectl scale deployment my-app --replicas=5
```

#### 로그 확인
```bash
kubectl logs <pod-name>
```

#### Pod 삭제
```bash
kubectl delete pod <pod-name>
```

#### 전체 배포 삭제
```bash
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
```

#### YAML 파일 유효성 검사
```bash
kubectl apply --dry-run=client -f mysql-deployment.yaml
```

### Cert-Manager 설치
```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
