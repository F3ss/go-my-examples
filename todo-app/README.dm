Run docker:
    docker run --name todoapp -p 5432:5432 -e POSTGRES_PASSWORD=user -e POSTGRES_USER=user -e POSTGRES_DB=todoApp -d --rm postgres:14-alpine

Migrate (файлы миграций sql, папка для миграций schema, имя файлов init):
    migrate create -ext sql -dir ./schema -seq init