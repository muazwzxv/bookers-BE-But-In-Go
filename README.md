
## Bookers backend but written in Go

### This is just a fun project of me trying to code the backend for Bookers but in Go

I will be using this stack for this project
- Go language
- Go Fibre 
- Gorm 
- Mysql
- Docker  


### TLDR
- The current docker image has problems when connecting to the database server ran locally outside the container
- Current solution is include the network tag when spawning containers
```bash
$ docker run --network="host" --name $CONTAINER_NAME -d bookers-be
```
- No need to mapped port due to the network setting set to host
