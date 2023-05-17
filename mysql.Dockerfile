FROM mysql:8.0.33

COPY ./db/migrations/create_todo_table.sql /docker-entrypoint-initdb.d/create_todo_table.sql

ENV MYSQL_ROOT_PASSWORD=my-secret-pw
ENV MYSQL_DATABASE=go-todo-db

EXPOSE 3306