# Etapa de compilación
FROM golang:1.18 as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o /app/hotel-app

# Etapa final
FROM alpine:latest

# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder /app/hotel-app /hotel-app

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/hotel-app"]
