# Installing golang in WSL

My Current setup is:

- Windows 10 Home Version 10.0.18363.693
- WSL Ubuntu 18.04
- VS Code 1.42.1

WSL (Windows Subsystem for Linux) is used to code in golang using VS Code WSL Remote Editing.

First check for the latest version available for linux amd64, as of now the latesgt version of golang is: 1.14, then call the following commands in your home directory:

```bash
user@host1:~$ export VERSION=1.14
user@host1:~$ wget -O - https://dl.google.com/go/go$VERSION.linux-amd64.tar.gz | sudo tar -C /usr/local -xz
user@host1:~$ export PATH=$PATH:/usr/local/go/bin
user@host1:~$ go version
go version go1.14 linux/amd64
```

## Test your installation

Hello world console application is located in hello.go file:

```golang
package main

import "fmt"

func main() {
 fmt.Printf("hello, world\n")
}
```

Build and run the hello world in golang:

```bash
user@host1:~$ go build hello.go
user@host1:~$ ./hello
hello, world
```

Congragulations, golang installation complete!
