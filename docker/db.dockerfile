FROM mysql:8.0.33
#FROM mysql:8.0.33-bullseye
#FROM mysql:8.0.33-alpine3.14

COPY ./docker/custom.cnf /etc/mysql/conf.d/custom.cnf