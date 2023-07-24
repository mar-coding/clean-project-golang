# clean-project-golang
simple golang project that follow the clean architecture rules

# How To run

### Download go dependencies
```
go mod tidy
```

### Make the migration
```
go run cmd/main.go -c config/config.yml migrate
```

### Run the main
```
go run cmd/main.go -c config/config.yml run
```

### See Swagger
http://localhost:8000/api-docs/