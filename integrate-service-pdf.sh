#!/bin/bash
set -e
for i in $(curl -s http://${CONSUL_SERVER}?recurse |jq -c '.[] | {Key,Value}'); do 
    KEY=$(echo ${i} |jq -r .Key |cut -d/ -f2 |tr [:lower:] [:upper:])
     VALUE=$(echo ${i} |jq -r .Value |base64 -d)

    #dotnet code needs the endpoint to be expressed in IP format
    if [[ ${KEY} == "ENDPOINT" ]]; then
        VALUE=$(echo ${VALUE} | sed -e 's/localhost/0.0.0.0/')
    fi

     export ${KEY}="${VALUE}"
done

echo "Starting $@"
exec "./$@"