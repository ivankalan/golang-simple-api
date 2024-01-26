### Golang Simple API

Run this command to start application: ```go run main.go```

Run this command to get the books data: ```curl localhost:8080/books```

Run this command to add new book data using json file: ```curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"```