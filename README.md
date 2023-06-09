# goAPIrest

## Pre-requisitos:

Instalar:
1. Docker
2. docker-compose


## Instrucciones:
1. Agregar en /etc/hosts
127.0.0.1       saludvirtual.local.ve

2. Crear docker network:
 docker network create traefik-net

3. Construir las images docker.
docker-compose build

4. Levantar los contenedores.
docker-compose up

5. Listar tareas.
docker-compose ps

6. Url de acceso:
http://saludvirtual.local.ve/countries
Métodos
* GET
* POST:
JSON

{
"Country": "España",
"Language": "Spanish"
}
