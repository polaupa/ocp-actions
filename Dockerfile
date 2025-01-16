# Usa una imagen base de Python
FROM python:3.9-slim

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos al contenedor
COPY app/ /app/

# Expone el puerto 8080
EXPOSE 8080

# Comando para iniciar el servidor HTTP integrado de Python
CMD ["python3", "-m", "http.server", "8080"]
