FROM mongo:latest

COPY ./0_prepare_user.js /docker-entrypoint-initdb.d/
COPY ./1_content_category.js /docker-entrypoint-initdb.d/