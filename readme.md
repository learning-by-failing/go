LearnGo
---
This repository contains basic and simple examples, used to learn Go programming language.

Examples and code are or inspired from official [golang documentation](https://golang.org/doc/code.html)

## Index
* [Features and basics](features.md)
* [helloword](src/scripts/helloworld.go)
* [variables](src/scripts/variables.go)
* [pointers](src/scripts/pointers.go)
* [environment variables](src/scripts/env-vars.go)
* [constants](src/scripts/constants.go)
* [functions](src/scripts/functions.go)
* [array and slices](src/scripts/arrayAndSlices.go)
* [maps](src/scripts/maps.go)
* [structs](src/scripts/structs.go) and [interfaces](https://gobyexample.com/interfaces)
* [concurrency](src/scripts/concurrency.go) vs [no_concurrency](src/scripts/no_concurrency.go)
* [channel](src/scripts/channel.go)
* [channel fibonacci](src/scripts/channel_fibonacci.go)
* [equivalent binary tree](https://tour.golang.org/concurrency/7) with [solution](src/scripts/equivalentBinaryTree.go)
* [Mutex](src/scripts/mutex.go)
* [exercise webCrawler](https://tour.golang.org/concurrency/10) with [solution](src/scripts/webCrawler.go)

## Command (go tool)
* create a module from the directory
```
go mod init 
```
* Install the module
```
go install {module}
```
* Running tests
```
go test
```
