# Base image
FROM python:3.9-slim

# Set working directory
WORKDIR /app

# Copy backend files
COPY app/ /app/

# Copy frontend files
COPY index.html /app/

# Install dependencies
RUN pip3 install flask
RUN pip3 install flask-cors

# Expose port
EXPOSE 8080

# Start the backend server
CMD ["python3", "server.py"]
