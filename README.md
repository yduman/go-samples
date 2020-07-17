# Why Go?

- Fast compile time
- Lots in common with C
- Reduces complexity of C
- Wicked fast build time
- Features
    - Lightweight type system
    - Concurrency
    - Automatic garbage collection
    - Strict dependencies
    - Convention

# Installing Go

[Downloads](https://golang.org/dl/)

### Verify Installation

```bash
# Verify Installation
$ which go
$ go version

# Add to bash or zsh profile
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

# Verify updates
echo $PATH
echo $GOPATH

# Create workspace
cd $GOPATH
mkdir go && cd go
mkdir src && cd src
```

# Documentation

[Documentation](https://golang.org/doc/)

# Go vs. JavaScript

- Go is strongly typed vs. dynamically typed
- Structs, Pointers, Methods, Interfaces vs. ES6 classes
- Explicit Error Handling vs. Built in
- Multi-Threaded (Concurrency, Goroutines, Sync) vs. Single-Threaded (Callbacks, async/await)
- Strong Opinions (Conventions, built in tooling and linters) vs. Fluid Opinions

# Basic Types

- Integer: int, int8, int16, int32, int64, uint...
- Float: float32, float64
- String: string
- Boolean: bool

# Misc

- `go doc <pkg>` is a nice way to look into the documentation! e.g. `go doc fmt`
- We can also look into specific functions: e.g. `go doc fmt.Println`