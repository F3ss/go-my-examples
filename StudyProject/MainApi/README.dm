Docker postgres:
    docker run --name example-project -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=user -e POSTGRES_DB=exampledb -d --rm postgres:14-alpine

Connect to psg in docker:
    docker exec -it todoapp bash    // Connect to container
    psql -U user todoapp            // Open DB
    \l                              // list databases
    \d                              // list of relations
    
Migrate:
    create init dirrectory with files:
        migrate create -ext sql -dir ./schema -seq init
    up:
        migrate -path ./schema -database 'postgres://user:user@localhost:5432/exampledb?sslmode=disable' up
    down:
        migrate -path ./schema -database 'postgres://user:user@localhost:5432/exampledb?sslmod=disable' down -all