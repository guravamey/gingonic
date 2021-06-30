# gingonic
POC for GO - GinGonic

Project setup : 

Run following commands : 

git clone git@github.com:guravamey/gingonic.git

go mod download

For Server startup : 
go run server.go


Local host 5000 - GET, POST, PUT and DELETE. ID as path param for PUT and DELETE. 
POST and PUT requires body. Body has name and age as JSON parameters. 

Example : 

curl --location --request POST 'localhost:5000/users/' \
--header 'Content-Type: application/json' \
--data-raw '{"name" : "Amey",
"age": 28}'
