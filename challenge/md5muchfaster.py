import requests, hashlib
from bs4 import BeautifulSoup

url ="LINK"

r = requests.session()
response = r.get(url)

soup = BeautifulSoup(response.content, "lxml")
finalstr = soup.h3.string

hash = hashlib.md5(finalstr).hexdigest()
data = {'hash': hash}

post = r.post(url, data = data)

print(BeautifulSoup(post.text, "lxml").p.string)