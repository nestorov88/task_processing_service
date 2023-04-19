# Task Processing Service

A job is a collection of tasks, where each task has a name and a shell command. Tasks may
depend on other tasks and require that those are executed beforehand. The service takes care
of sorting the tasks to create a proper execution order.

## Call the service
```
curl -d @mytasks.json -H "Accept: text/plain" ${ADDR} | bash

curl -d @mytasks.json -H "Content-type: text/application-json" ${ADDR}
```

### or
```
make test-curl
```

## Setup and run locally
```
go mod tidy
go run cmd/app/main.go
```

### or

```
make tidy
make run
```

## Running dockerized app

```
docker build -f ./build/docker/Dockerfile -t task_processing_service .
docker run -dp 4000:4000 task_processing_service
```

### or
```
make start
```

## Running existing test

```
make test
```

