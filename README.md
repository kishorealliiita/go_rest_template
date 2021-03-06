# GoLang rest API endpoint creation.

REST API endpoint creation in GoLang. Sample boilerplate repository with pipeline definition (travis.yml) and with Docker support.

## Requirements
* Go 
* Docker

## Run the tool using GO
* go build
* ./go-rest_template 
* curl http://localhost:8080 
* curl http://localhost:8080/endpoint
* or use browser

## Run the tool using Docker
* docker build -t go_rest_template .
* docker run -d -p 8080:8080 go_rest_template
* use above curl commands or browser.

## Run tests
* go test -v ./...
