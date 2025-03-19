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

```sql
CREATE DATABASE goblog;
CREATE USER 'gouser'@'localhost' IDENTIFIED BY 'gopassword';
GRANT ALL PRIVILEGES ON goblog.* TO 'gouser'@'localhost';
FLUSH PRIVILEGES;
```
