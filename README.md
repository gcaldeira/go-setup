# Go Setup

## Prerequisites
1. Go runtime and libs up and running on machine ([Instaling go](https://golang.org/doc/install))
1. Git installed and configured
 
## Getting Started (Hello World)

After install Go libs 

### Creating GoLang project basic folders

```
mkdir go-setup
cd go-setup
go mod init github.com/jhon/my-project
mkdir -p cmd/cli
```

On folder cmd/cli create a file called **main.go** with the content above:

```
package main

import "fmt"

func main()  {
	fmt.Println("Hello!")
}
```

### Build an run

```
go build -o releases/my-app cmd/cli/main.go
./releases/my-app 
Hello!
```
