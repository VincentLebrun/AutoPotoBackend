wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- echo "Database is ready."

CompileDaemon --build="go build -o main ./cmd/ApiHappyTravelling/main.go" --command="./main"