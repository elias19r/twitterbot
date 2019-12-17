#!/bin/bash

truncate -s 0 memes.txt
truncate -s 0 tmp.txt
for i in {1..10}; do
# for i in 1; do
	url='http://capinaremos.com/page/'$i'/?s=capina+meme+factory'
	response=$(curl --silent --write-out 'HTTPSTATUS:%{http_code}' $url)
	body=$(echo $response | sed -e 's/HTTPSTATUS:.*//g')
	status=$(echo $response | tr -d '\n' | sed -e 's/.*HTTPSTATUS\://g')
	echo "status $status | url $url"
	if [ ! $status -eq 200 ]; then
		continue
	fi
	links=$(
		echo $body \
		| grep -oE 'http://capinaremos.com/[0-9]{4}/[0-9]{2}/[0-9]{2}/[a-z0-9-]+/' \
		| sort \
		| uniq
	)
	sleep 1

	for url in $links; do
		response=$(curl --silent --write-out 'HTTPSTATUS:%{http_code}' $url)
		body=$(echo $response | sed -e 's/HTTPSTATUS:.*//g')
		status=$(echo $response | tr -d '\n' | sed -e 's/.*HTTPSTATUS\://g')
		echo "status $status | url $url"
		if [ ! $status -eq 200 ]; then
			continue
		fi
		images=$(
			echo $body \
			| grep -oP 'data-large-file=".+?"' \
			| tr -d '"' \
			| sed -e 's/data-large-file=//g' \
			| sort \
			| uniq
		)
		for img in $images; do
			echo $img >> memes.txt
			cat memes.txt | sed -e 's/https:\/\/i[0-9]/https:\/\/i0/g' > tmp.txt
			sort tmp.txt | uniq > memes.txt
			wc -l memes.txt
		done
		sleep 1
	done
done
rm tmp.txt