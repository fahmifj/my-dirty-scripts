#!/bin/bash
for dir in `cat wordlist`; do 
        status=$(curl -s http://example.com/$dir -o /dev/null -w "%{http_code}")
        if [ $status -eq 200 ]
        then
          echo "$dir"
        fi
done