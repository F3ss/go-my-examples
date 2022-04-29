Docker postgresql:
    docker run --name todoapp -p 5432:5432 POSTGRES_PASSWORD=user POSTGRES_USER=user POSTGRES_DB=todoapp -d --rm postgres:14-alpine

Migrate -ext тип файлов миграции, -dir куда сложить, -seq имена файлов:
    migrate create -ext sql -dir ./scheme -seq init

    up:
        migrate -path ./schema -database 'postgres://postgres:todoapp@localhost:5432/postgres?sslmode=disable' up
