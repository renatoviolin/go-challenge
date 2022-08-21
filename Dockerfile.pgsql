FROM postgres:14-alpine

COPY ./db/init.sql /docker-entrypoint-initdb.d/create_tables.sql
COPY ./db/data.csv /tmp/data.csv

# docker build -t unico-db -f Dockerfile.pgsql --no-cache . && docker run --name unico-db -p 5432:5432 --env-file .env -d unico-db