#!/bin/sh
wait_seconds=30
END=5
echo "Testing rabbitMQ status"
for i in $(seq 1 $END); do
  echo "Tentative: "$i;
  sleep $wait_seconds
  ping=`ping -c 1 rabbitmq | grep bytes | wc -l`
	if [ "$ping" -gt 1 ]; then
		echo "Consummer Server Up"
		cd consummer
		exec /go/bin/air
	fi
done
echo "RabbitMQ server is down"
echo "Quitting..............."
