How create, manage and maintain modules (Best practise)
---
* A module is one or more related packages.
* Configured via *go.mod*

## Modules creation
Inside the directory where u wanna create your module
```
go mod init {module-identifier}
```
i.e. **go mod init github.com/learning-by-failing/go/src/modules**

Then to build the module type just `go build .` from the root directory of the module

The name of the executable file will be the last part of the initialised module: *modules*
