# Installing golang in WSL

My Current setup is:

- Windows 10 Home Version 10.0.18363.693
- WSL Ubuntu 18.04
- VS Code 1.42.1

WSL (Windows Subsystem for Linux) is used to code in golang using VS Code WSL Remote Editing.

## Download and Setup

First check for the latest version available for linux amd64, as of now the latest version of golang is: [1.14](https://golang.org/dl/), then call the following command in your home directory:

Set the golang version you want to install:

```bash
user@host1:~$ export VERSION=1.14
```

Download and unpack it into /usr/local

```bash
user@host1:~$ wget -O - https://dl.google.com/go/go$VERSION.linux-amd64.tar.gz | sudo tar -C /usr/local -xz
```

Update the PATH env variable in $HOME/.profile:

```bash
export PATH=$PATH:/usr/local/go/bin
```

In order to make .profile changes available you could either relogin or source the .profile changes by:

```bash
user@host1:~$ source $HOME/.profile
```

Now you could check the current version of golang:

```bash
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

Congratulations, golang installation complete!
