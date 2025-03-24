### 초기 인증서 발급 (수동으로 진행)
초기 인증서 발급은 수동으로 진행해야 합니다. 아래 명령어를 실행하여 인증서를 발급받습니다.

```sh
# 초기 인증서 발급
docker-compose -f docker-compose.initial.yml up -d
# 초기 인증서 발급 후 중지
docker-compose -f docker-compose.initial.yml down
# docker-compose 시작
docker-compose up -d
```
