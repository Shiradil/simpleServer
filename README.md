# simpleServer
Assignment for Advanced Programming classes. SE-2204 Shirbaev Adilzhan

This program awaits incoming POST requests on port 8080.
Receives JSON data from body.
Checks the validity of JSON data.
Prints message field to console if exists.
Sends a JSON response with the message "Data successfully received" and a status of "success.".

To use this program: `go run .`
You'll see localhost adress on console, just paste it to your browser url field
Now go to postman and send JSON request fith the field "message": "text"
