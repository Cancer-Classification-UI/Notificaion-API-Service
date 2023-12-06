# Notification API Service
[![Notification API Image Deployment](https://github.com/Cancer-Classification-UI/Notification-API-Service/actions/workflows/docker-image.yml/badge.svg?branch=main)](https://github.com/Cancer-Classification-UI/Notification-API-Service/actions/workflows/docker-image.yml)

This service is responsible for notification logic. Handles any communication to end users, such as emailing and sending SMS.

API Documentation is created using the [Swagger](https://swagger.io/). The url for the Swagger UI is on the same port as the `APP_PORT` in the `.env` file at 

```
http://<ip>:<APP_PORT>/swagger/index.html
``` 

If you are running locally it would be at [http://localhost:8087/swagger/index.html](http://localhost:8087/swagger/index.html)

# How to run

## Create
### `.env` Creation
Create a `.env` file
```bash
touch .env
```
Edit the `.env` file with any editor of your choice
```bash
vim .env
```

### `.env` Template
```
APP_PORT=8087 // Standard port for this microservice
LOG_LEVEL=trace
METHOD_LOGGING=false
```
> Additional fields will also be required in the `.env` file to run the microservice successfully. Here is a basic template of the `.env`. Customize to your liking. This template will change as the microservice matures and implements new features.

## Install

### Go
You will also need to [install go](https://go.dev/doc/install), install it per your operating system you are using.

### Swagger (Swaggo)
You will need the swaggo package to create the swagger files.
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Docker (Optional)
If you want to use the docker containers provided (Recommended) [install docker](https://www.docker.com/get-started/). It also is required if you want to use the scripts.

## Build
<details close>
<summary><h3>With Docker</h3></summary>
<br>

```bash
docker build -t ccu-notification-api .
swag init
```
</details>

<details close>
<summary><h3>Without Docker</h3></summary>
<br>

```bash
go build
swag init
```
</details>

## Run
<details close>
<summary><h3>With Docker</h3></summary>
<br>

Make sure you have a `log.txt` file in the repo directory, otherwise it wont be able to attach the log.txt and will give a warning and sometimes even an error
```bash
touch log.txt
```
Then run the docker image
```bash
./scripts/start.sh
```
or manually with
```bash
docker run -d -p $(cat .env | grep APP_PORT= | cut -d: -f2 | awk '/^/ { print $1":"$1 }') -v $(pwd)/log.txt:/usr/src/app/log.txt --name notification-api ccu-notification-api
```
</details>

<details close>
<summary><h3>Without Docker</h3></summary>
<br>

```bash
go run ccu
```
or if you dont want to build
```bash
go run main.go
```
## (Optional) Update package checksums and download dependencies
```bash
go mod tidy
``` 
</details>

## Other

### View Docker terminal or unmounted files
If you launched the container using docker, you can execute a sh terminal inside the container to gain access to it and browse around.
```bash
docker exec -it notification-api /bin/sh
```
>Leave the shell with `Ctrl+D`

If you want to see the actual go service (the console the `go run ccu` command produces) then
```bash
docker attach notification-api
```
> Be careful as it hard to detach as the key bind `Ctrl+P then Ctrl+D` is often used by many programs, so you may not be able to detach correctly.

### Shutting down docker container
If you want to fully shutdown the container
```bash
./scripts/stop.sh
```
or manually with
```bash
docker kill notification-api
docker rm notification-api
```

