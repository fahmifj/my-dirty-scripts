import zipfile
import sys

try:
	file = sys.argv[1]
	wordlist = sys.argv[2]
except IndexError:
	print('[-] Usage: %s <file.zip> <wordlist>' % sys.argv[0])


zip_file = zipfile.ZipFile(file)
with open(wordlist, "rb") as wl:
	words = wl.readlines()
	for word in words:
		try:
			zip_file.extractall(pwd=word.strip())
		except:
			continue	
		else:
			print(f"[+] Password found: {word.decode('UTF-8').strip()}")
