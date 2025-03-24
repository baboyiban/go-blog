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

### Cert-Manager 설치
```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
```

### 사용 방법

#### Minikube 설치 및 실행 (로컬 환경)
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
