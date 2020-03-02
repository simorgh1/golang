# Simple Rest Api for RaspberryPi

My Current setup is:

- Windows 10 Home Version 10.0.18363.693
- WSL Ubuntu 18.04
- VS Code 1.42.1
- Go 1.14
- RaspberryPi 4

This is a Simple REST API web application in golang with 2 endpoints:

- /hello
- /headers

## Build for RaspberryPi

For RaspberryPi we have to build for arm architecture using arm7, this can be set as environment variables before running the build command:

```bash
user@host1:~$ env GOOS=linux GOARCH=arm GOARM=7 go build hello.go
```

Then we copy the executable to raspberrypi:

```bash
user@host1:~$ scp ./hello pi@raspberrypi:/home/pi
```

then run ./hello and check the service using curl from wsl command line:

```bash
user@host1:~$ curl http://<your-raspi-ip>:8090/headers
```

## Building without debug files

In order to reduce the size og compiled golang build, we could pass the following build arguments for building a smaller size by excluding debug information.

```bash
go build -ldflags "-s -w" ./hello.go
```

It reduces the size of ./hello by 30%.
