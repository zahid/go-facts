# Use these to build then run this image
# docker build -t zahid/go-facts-webserver client/
# docker run --name go-facts-webserver --link go-facts-api -p 80:80 -d zahid/go-facts-webserver
#
FROM nginx

COPY nginx.conf /etc/nginx/nginx.conf
COPY html/ /usr/share/nginx/html

EXPOSE 80
