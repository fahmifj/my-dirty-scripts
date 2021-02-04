Just a stupid tool to quickly generate a reverse shell for ctf.  
For easy call, put the binary on your `$HOME/bin` then use `export $PATH=$HOME/bin:$PATH` 

Build the executable/binary file with
```
go build mkrev.go
```

```
$ mkrev
Usage:
    mkrev [shell] [interface] [port]
    mkrev py tun0 9000
Shell:
    py, bash, nc, php, ps, pl, rb
```
