# Pull nginx base image
FROM nginx:latest

# Expost port 80
EXPOSE 80

# Copy custom configuration file from the current directory
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf
#COPY hosts /etc/hosts