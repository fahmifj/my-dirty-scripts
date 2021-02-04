#!/bin/bash
url=$1 # http or https://www.example.com
wordlist=$2
for word in `cat $wordlist`; do 
        status=$(curl -s $url/$dir -o /dev/null -w "%{http_code}")
        if [ $status -eq 200 ]
        then
          echo "$dir"
        fi
done