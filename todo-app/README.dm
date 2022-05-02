Docker postgresql:
    docker run --name todoapp -p 5432:5432 -e POSTGRES_PASSWORD=user -e POSTGRES_USER=user -e POSTGRES_DB=todoapp -d --rm postgres:14-alpine

Connect to psg in docker:
    docker exec -it todoapp bash    // Connect to container
    psql -U user todoapp            // Open DB
    \l                              // list databases
    \d                              // list of relations

Migrate -ext тип файлов миграции, -dir куда сложить, -seq имена файлов:
    migrate create -ext sql -dir ./scheme -seq init

    up:
        migrate -path ./schema -database 'postgres://user:user@localhost:5432/todoapp?sslmode=disable' up
    down:
        migrate -path ./schema -database 'postgres://user:user@localhost:5432/todoapp?sslmode=disable' down -all
