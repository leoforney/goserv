# Stage 1: Build the React frontend
FROM node:alpine as react-build
WORKDIR /app
COPY client/package.json client/package-lock.json ./
RUN npm install
COPY client/ .
RUN npm run build

# Stage 2: Build the Go backend with CGO enabled
FROM golang:buster as go-build
WORKDIR /go/src/app
COPY server/ .
ENV CGO_ENABLED=1
RUN apt-get update && apt-get install -y libsqlite3-dev && \
    go build -o /goserv

# Stage 3: Setup the final image
FROM debian:latest
# Install NGINX and SQLite libraries
RUN apt-get update && apt-get install -y nginx sqlite3 libsqlite3-dev
# Setup directories and clean the web directory
RUN mkdir -p /run/nginx && \
    mkdir -p /var/www/html && \
    rm -rf /var/www/html/* && \
    rm /etc/nginx/sites-available/default

# Copy the built React app to the appropriate directory
COPY --from=react-build /app/build /var/www/html

# Copy the Go binary
COPY --from=go-build /goserv /usr/local/bin/goserv

# nginx config file push
COPY --from=go-build /go/src/app/nginx.conf /etc/nginx/sites-available/default

# Command to start NGINX and the Go server
CMD nginx && /usr/local/bin/goserv

# Expose port 80
EXPOSE 80