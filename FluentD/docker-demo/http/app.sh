#!/bin/sh
while true
do
    echo "Sending logs to FluentD"
    curl -X POST -d 'json={"foo":"bar"}' http://fluentd:9880/http-my-app.log
    sleep 5
done