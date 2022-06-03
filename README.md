# Protocol buffers in Go Intro

## Description

Simple intro to protocol buffers, I created an example based on a Youtube video and called the objects created in a **.proto** file from the **main.go** file.

It allowed me to begin my understanding on protocol buffers because I want to know how to work with GRPC in a near future.

## Run

### Requirements

All the required files are included so to the example is easier to run. To run the program you will need.

* Go (version 1.18, it may work in previous and later versions).

* ProtoC for Go (optional, if you want to modify things in the .proto file and rebuild the proto file. Currently, person.pb.go has everything needed from the main to execute the program).

## Execute command

Run the following command from the command line (Bash, ZSH, Fish, Iterm, CMD, etc)

```
go run .
```

If it fails in Windows, it may work executing the command like this

```
go build
```

And then open the .exe file from command line or double clicking it.

From command line it may be like so:

```
start nameOfTheFile.exe
```


## Source of tutorial I watched

[https://www.youtube.com/watch?v=NoDRq6Twkts](https://www.youtube.com/watch?v=NoDRq6Twkts)
