curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"ID": 1, "Name": "John Doe", "Email": "john@example.com"}'

curl -X PUT http://localhost:8080/users -H "Content-Type: application/json" -d '{"ID": 1, "Name": "John Smith", "Email": "john@example.com"}'

curl http://localhost:8080/users/1
