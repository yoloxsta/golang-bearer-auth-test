# Frontend Dockerfile
FROM nginx:alpine

# Copy static files to Nginx html directory
COPY . /usr/share/nginx/html/

# Expose port 80
EXPOSE 80

# Nginx runs automatically with the default command
# No need to specify CMD as nginx:alpine has it built-in
