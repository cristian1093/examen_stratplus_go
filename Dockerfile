# Utiliza la imagen oficial de Golang como base
FROM golang:latest

# Instala el cliente MySQL
RUN apt-get update && apt-get install -y mysql-client

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia el código fuente de la aplicación Go al directorio de trabajo
COPY . .

# Define las variables de entorno para MySQL
ENV MYSQL_USER=cristian
ENV MYSQL_PASSWORD=GeNeSiS2
ENV MYSQL_DB=test
ENV MYSQL_HOST=localhost
ENV MYSQL_PORT=3306

# Instala las dependencias de la aplicación Go
RUN go get -d -v ./...

# Compila la aplicación Go
RUN go install -v ./...

# Expone el puerto 8080 para la aplicación Go
EXPOSE 8080

# Comando para ejecutar la aplicación Go al iniciar el contenedor
CMD ["app"]
