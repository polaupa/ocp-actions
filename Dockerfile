
FROM python:3.9-slim

WORKDIR /app

COPY app/ /app/

COPY index.html /app/

RUN pip3 install flask
# Cross-Origin Resource Sharing (CORS) is a security feature that allows restricting resources on a web page to be requested from another domain outside the domain from which the first resource was served.
# RUN pip3 install flask-cors

EXPOSE 8080

CMD ["python3", "server.py"]
