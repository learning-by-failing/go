GOlang features
---
* system language
* automatic garbage collector
* only ~25 keywords
* cross platform
* compile language (like C but simpler)
* strongly typed
* supports type inference: i.e. `a := 5.5` would be automatically declared like float.
* case sensitive -> only functions with the first capital letters are exported
* Slices are references contiguous portion of an array, with flexible length 

### Maps 
**Maps are reference types**
* passed to functions by reference: changes made by functions visible to the caller
* unsafe for concurrency 
* cheap to pass
* specify the size of a map is considerable a good practise

### Structs 
**Custom data types** or a collection of names fields, i.e.
```
type Circle struct {
    r float64
    d float64
    c float64
}
```

### Concurrency (goroutine)
[learn concurrency](https://github.com/golang/go/wiki/LearnConcurrency)
* Lighter weight
* Go manages Goroutins 
* goroutines are built on top of Threads: less switching because i.e. when a thread is sleeping, simple go switch to
another goroutine that is **less than schedule another Thread**.
* Faster stat-up times
* Safe communitcation between goroutines
* Go's concurrency model is **Actor model**, **Communicating sequential process** (using *channels* to passing massages)

### Workspace
We need three folder:
* **source** folder where put the code
* **package** folder where put any package
* **bin** folder for any compile program

**The whole code of Go functions are inside src directory in [Go repository](https://github.com/golang/go/blob/master/src)**

## commands
* `go run {gofile.go}`