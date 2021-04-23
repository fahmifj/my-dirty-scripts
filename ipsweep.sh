#!/bin/bash

if [ "$1" == "" ]
then
	echo "Scan last IP bit for /24 network range"
	echo "[-] Usage:"
	echo "./ipsweep.sh 192.168.1"
	echo "./ipsweep.sh 10.10.1"
else
	for ip in `seq 1 254`
		do sleep 0.001; ping -c 1 $1.$ip | grep -i "bytes from" | cut -d " " -f4 | tr ":" " - Alive" &
	done
fi
