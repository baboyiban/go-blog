### 폴더 구조
```markdown
/go-blog
├── cmd/                  # 실행 가능한 애플리케이션
│   └── server/
│       └── main.go       # 진입점(Entry Point)
├── internal/             # 비공개 패키지
│   ├── handlers/         # HTTP 핸들러
│   │   └── user.go
│   ├── models/           # 데이터 모델
│   │   └── user.go
│   ├── middleware/       # 미들웨어
│   │   └── logging.go
│   └── database/         # DB 연결
│       └── postgres.go
├── web/                  # 프론트엔드 리소스
│   ├── static/           # CSS/JS/이미지
│   │   └── style.css
│   └── templates/        # HTML 템플릿
│       └── index.html
├── pkg/                  # 공개 패키지 (선택적)
├── api/                  # OpenAPI/Swagger
├── configs/              # 설정 파일
│   └── config.yaml
├── go.mod                # 모듈 정의
├── go.sum
├── .env                  # 환경 변수
└── Dockerfile            # 컨테이너 설정
```

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
