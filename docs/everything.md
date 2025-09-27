
## Dependency Tracking

To enable dependency tracking for your code by creating a `go.mod` file, run the `go mod init` command, giving it the name of the module your code will be in.
The name is the module's module path.
Usually, the module path is the repository location where the code is kept, e.g. `github.com/mymodule`.

TODO: Check [Managing Dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)

TODO: Check [Authenticating](https://go.dev/ref/mod#authenticating)

## Packages

The Go package registry is available on [pkg.go.dev](https://pkg.go.dev/).

The [`fmt`](https://pkg.go.dev/fmt) package contains dunctions for formatting text, including printing ot the console.
It is part of the [standard library](https://pkg.go.dev/std) packages.

## Run Code

The `main` function executes by default when you run the code in the `main` package.

## Commands

The command `go mod tidy` adds used or removes unused module dependencies and updates the `go.sum` file.
