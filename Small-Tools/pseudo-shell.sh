#!/bin/bash

URL="${1}"
while true;do
echo -n "$ "; read cmd
curl -sX POST "${URL}" --data-urlencode "cmd=$cmd"
done

