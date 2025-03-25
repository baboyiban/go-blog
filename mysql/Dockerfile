FROM mysql:8.0-debian

RUN apt-get update && apt-get install -y tzdata && rm -rf /var/lib/apt/lists/*

COPY ./init.sql /docker-entrypoint-initdb.d/
