### 실행 명령어 (프로젝트 루트에서 실행)
```bash
# 전체 서비스 빌드 및 실행 (최초 1회)
docker-compose up --build

# 백그라운드 실행 (데몬 모드)
docker-compose up -d

# 실행 중인 컨테이너 확인
docker-compose ps

# 로그 확인 (app 서비스)
docker-compose logs -f app
```

### 개발 환경 실행
```bash
# Go API 서버 + Svelte 개발 서버 동시 실행
docker-compose -f docker-compose.yml -f docker-compose.override.yml up
```

### 접속 테스트
```bash
# 컨테이너 정상 실행 확인
curl http://localhost:8080/users/1

# MySQL 접속 테스트
docker exec -it go-blog-mysql-1 mysql -uroot -psecret
```

### 중지 및 정리
```bash
# 컨테이너 중지 (데이터 보존)
docker-compose stop

# 컨테이너 완전 삭제 (데이터 포함 전체 제거)
docker-compose down -v
```

### 기술스택
```markdown
Go
Docker
SQL
HTML
CSS
Typescript
- Jenkins
```

**5. 다음 단계 (제안)**

* **API 엔드포인트 개발:** `internal\handlers`, `internal\routes`, `internal\models` 폴더를 사용하여 Go 백엔드 API 엔드포인트를 구현합니다. (예: 게시글 목록 조회, 사용자 정보 조회 등)
* **데이터베이스 연동:** `internal\database` 패키지를 사용하여 MySQL 데이터베이스와 연동하고, 실제 데이터를 CRUD (생성, 읽기, 수정, 삭제) 하는 기능을 구현합니다.
* **Svelte 프론트엔드 개발:** `web\svelte\src` 폴더에서 Svelte 컴포넌트 및 페이지를 개발하여 사용자 인터페이스를 구축하고, Go 백엔드 API를 호출하여 데이터를 표시하고 상호작용하는 기능을 구현합니다.
* **테스트 작성:** `internal\handlers\user_test.go` 와 같이 Go 백엔드 핸들러에 대한 단위 테스트를 작성하여 코드의 안정성을 높입니다. Svelte 프론트엔드 컴포넌트에 대한 테스트도 작성하는 것을 고려해보세요.
* **배포:** Docker 이미지를 컨테이너 레지스트리 (Docker Hub, GCR, ECR 등) 에 푸시하고, 클라우드 플랫폼 (AWS, GCP, Azure 등) 또는 서버 환경에 배포합니다.
