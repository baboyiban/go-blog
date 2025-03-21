FROM nginx:1.25-alpine

# 자체 서명된 인증서 복사
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY certs/cert.pem /etc/nginx/certs/cert.pem
COPY certs/key.pem /etc/nginx/certs/key.pem

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
