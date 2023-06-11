# goAPIrest

Se basa en los siguientes artículos:


1. [Golang CRUD](https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121)
2. [Docker y Docker-compose](https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8)
3. [Kubernetes](https://levelup.gitconnected.com/deploying-dockerized-golang-api-on-kubernetes-with-postgresql-mysql-d190e27ac09f)
4. [RabbitMQ](https://blog.devgenius.io/using-rabbitmq-with-golang-and-docker-e674831c959c)
5. [Sending Email with golang](https://medium.com/@loginradius/different-ways-to-send-an-email-with-golang-b79475460240)
6. [Golang environment variables](https://medium.com/@loginradius/different-ways-to-use-environment-variables-in-golang-46e1d1e515b7)
7. [Sending html email in golang](https://medium.com/hackernoon/sending-html-email-using-go-c464d03a26a6)
8. [Golang stripe](https://medium.com/@ksandeeptech07/creating-and-managing-charges-with-stripe-in-golang-87b4c1deb250)
9. [Docker golang mysql](https://articles.wesionary.team/dockerize-a-golang-applications-with-mysql-and-phpmyadmin-hot-reloading-included-86eb7a6cf8d5)
10. [zerolog](https://github.com/rs/zerolog)

## Pre-requisitos:

Instalar:
1. Docker
2. docker-compose


## Instrucciones:
1. Agregar en /etc/hosts
127.0.0.1       goapirest.local.ve


2. Construir las images docker.
docker-compose build
Ó
make build

3. Levantar los contenedores.
docker-compose up

4. Listar tareas.
docker-compose ps

make ps 
5. Ver logs
docker-compose logs -f
o 
make logs

6. Pruebas unitarias
docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit

o
make unitest

7. Limpiar contenedores
docker system prune -a
o 
make purge


8. Url de Bienvenida:

[http://goapirest.local.ve:8080/](http://goapirest.local.ve:8080/)

7. Pruebas del Endpoint
Revisar el artículo [Golang CRUD](https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121)
8. Pruebas unitarias usando docker


TODO:
1. Ajustar variables de entorno según sea local, producción, develop, testing.
2. Arreglar las variables para que despliegue pruebas unitarias.