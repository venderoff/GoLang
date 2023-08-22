# GoLang

download impage posgres 12-alpine // pull the docker image 
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

// run the container in cmd , this will run the commandline SQL
docker exec -it postgres12 psql -U root

//logs 
docker log postgres12

//create initial schema on cmd || point the directory to project name
migrate create -ext sql -dir db/migration -seq init_schema

***init_schema_up and init_schema_down are created. So that if we migrate porgressively it can save sequentialyy.
Incase we wish to move backwards then "init_schema_down" will help.

//stop a docker container 
docker stop "name"

//list all container , without status
docker ps -a

//start container
docker start "name"

//install make
//first install chocolatey
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

//install make
choco install make 

//remove container
docker rm postgres12

//run file from project
migrate -path db\migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

