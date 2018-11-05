FROM mongo:latest

COPY content_category.json /content_category.json
CMD mongoimport --host mongodb --db go_db --collection content_category --type json --file /content_category.json --jsonArray