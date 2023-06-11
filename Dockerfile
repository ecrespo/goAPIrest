# Partimos de la imagen base de golang
FROM golang:1.20.5-bullseye as builder

ENV GO111MODULE=on

# Agregamos información del mantenedor
LABEL maintainer="Ernesto Crespo <ecrespo@gmail.com>"

# Instalamos git.
# Git es necesario para obtener las dependencias.
RUN apt-get update && apt-get install -y --no-install-recommends git

# Establecemos el directorio de trabajo actual dentro del contenedor
WORKDIR /app

# Copiamos los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargamos todas las dependencias. Las dependencias se almacenarán en caché si no se modifican los archivos go.mod y go.sum
RUN go mod download

# Copiamos el código fuente desde el directorio actual al directorio de trabajo dentro del contenedor
COPY . .

# Compilamos la aplicación de Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Iniciamos una nueva etapa desde cero
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /root/

# Copiamos el archivo binario precompilado de la etapa anterior. Observa que también copiamos el archivo .env
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Exponemos el puerto 8080 hacia el exterior
EXPOSE 8080

# Comando para ejecutar el archivo binario
CMD ["./main"]