#!/bin/bash

CONTAINER_ID=8dd3a20adff9
DURATION=5 # ระยะเวลาในการเก็บข้อมูล (s)
INTERVAL=1 # ช่วงเวลาในการเก็บข้อมูล (s)

total_cpu=0
total_memory=0
count=0

for ((i=1; i<DURATION; i+=INTERVAL)); do
    cpu=$(docker stats --no-stream --format "{{.CPUPerc}}" $CONTAINER_ID | tr -d '%')
    memory=$(docker stats --no-stream --format "{{.MemPerc}}" $CONTAINER_ID | tr -d '%')

    echo "$(date '+%H:%M:%S')"
    echo "CPU usage: $cpu%"
    echo "Memory usage: $memory%"
    echo ""

    total_cpu=$(echo "$total_cpu + $cpu" | bc)
    total_memory=$(echo "$total_memory + $memory" | bc)

    count=$((count+1))
    # sleep $INTERVAL
done

average_cpu=$(echo "scale=2; $total_cpu / $count" | bc)
average_memory=$(echo "scale=2; $total_memory / $count" | bc)

echo "Average CPU usage: $average_cpu%"
echo "Average Memory usage: $average_memory%"

