
## Helpful Resources

The repository [awesome-go](https://github.com/avelino/awesome-go) provides a list of go libraries for various use cases.

The package [github.com/go-playground/validator/v10](https://pkg.go.dev/github.com/go-playground/validator/v10) provides struct validation that goes well for API.


## Dependency Tracking

To enable dependency tracking for your code by creating a `go.mod` file, run the `go mod init` command, giving it the name of the module your code will be in.
The name is the module's module path.
Usually, the module path is the repository location where the code is kept, e.g. `github.com/mymodule`.

TODO: Check [Managing Dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)

TODO: Check [Authenticating](https://go.dev/ref/mod#authenticating)

## Packages

The Go package registry is available on [pkg.go.dev](https://pkg.go.dev/).

Go code is grouped into packages, and packages are grouped into modules.
Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

The [`fmt`](https://pkg.go.dev/fmt) package contains functions for formatting text, including printing ot the console.
It is part of the [standard library](https://pkg.go.dev/std) packages.

## Run Code

The `main` function executes by default when you run the code in the `main` package.
Code executed as an application must be in a `main` package.

## Commands

The command `go mod tidy` adds used or removes unused module dependencies and updates the `go.sum` file.

The command `go mod init example.com/greetings` initializes a new module by creating a `go.mod` file.
If the module should be published, the path _must_ match the path from which it can be downloaded.

The command `go mod edit -replace example.com/greetings=../greetings` can be used to redirect Go tools to a specific location.
This can be used for developing modules locally that reference each other.

The command `go get golang.org/x/example/hello/reverse` downloads a dependency and adds it to the `go.mod` and `go.sum` files.
If a dependency is listed in the `import` section of a `.go` file, it will be downloaded by running `go get .` automatically.

The command `go test -v` executes tests with verbose output.

The command `go build` compiles the packages, along with their dependencies, but doesn't install the results.
The resulting file / executable can be run from the console.

The command `go install` compiles and installs the packages.

The command `go work init ./hello` creates a workspace file `go.work` that specifies the modules contained in the workspace.
It can be used instead of the `-replace` flow for the command `go mod edit`.
Add new modules to the workspace using `go work use ./example/hello`.

## Language Specific

The operator `:=` is used for quick variable declarations and initializations inside functions, e.g. `message := "Hello World"`.

A function whose name starts with a capital letter can be called by a function not in the same package.
This is known in Go as an _exported_ name.

A `slice` is an array, except that its size changes dynamically as you add and remove items.
When declaring a `slice`, you omit its size in the brackets `[]string`.
This tells Go that the size of the array underlying the slice can be dynamically changed.
Iterating through a `slice` can be achieved by using the `for index, value := range values {...}` syntax.

### Error Handling

It is common to return an error as a value so the caller can check for it.
