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

To retrieve dependencies, run the command `go get -u {module}`, i.e go get -u github.com/gorilla/mux/{v[n]}.
If u omit the version number from the path of your module (last slash), by default version 1 is imported.

To list all modules depending on your application, `go list -m all`

Listining version of a module with `go list -m -versions github.com/gorilla/mux`

Verify build of thirdy part modules with `go mod verify`: this operation is really important to verify the hash
assigned to the code of your modules dependencies because during the build, this verification is not performed.
If u have some modules that don't pass the verification, the only way to make it works again is to delete the module
and retrieve it again with *go get*.

**Run `go mod tidy` to remove all not used modules from your application.**

## Module versioning (semantic versining... look at https://semver.org)
i.e. v1.5.3-pre1

**v** version prefix (required)

**1** Major revision (likely to break backward compatibility)

**5** Minor revision (new features, doesn't break BC)

**3+** Patch (bug fixes, no new features doesn't break BC)

**pre1** Pre-release of a new version, if applicable (text is arbitrary)

See versioning rules example module [https://github.com/rsc/quote](https://github.com/rsc/quote)
