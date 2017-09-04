# blog-go

## Introduce
> Project use go v1.9 and echo framework

> database use mongodb

> dependency management tool use golang/dep

## Required
- [x] docker
- [x] docker-compose

## clone Project

```sh
$ git clone git@github.com:dung13890/blog-go.git
$ cd blog-go
$ mkdir vendor
```

## Run with docker

```sh
$ docker-compose up -d
$ docker-compose exec api bash

$ dep ensure
```

## run Project

```sh
go run main.go
```
