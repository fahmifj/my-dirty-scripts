#!/bin/bash

network=$1
if [ "$network" == "" ]
then
	echo "Scan last IP bit for /24 network range"
	echo "[-] Usage:"
	echo "./ipsweep.sh 192.168.1.x"
	echo ""
else
	$network=$(echo $network | cut -d'.' -f1-3)
	for ip in `seq 1 254`
		do (ping -c 1 ${network}.${ip} | grep "bytes from" &)
	done
fi
