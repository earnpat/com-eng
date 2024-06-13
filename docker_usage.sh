#!/bin/bash

CONTAINER_ID=2812586a21b2
DURATION=3600 # ระยะเวลาในการเก็บข้อมูล (s)
INTERVAL=1 # ช่วงเวลาในการเก็บข้อมูล (s)
OUTPUT_FILE="usage-rest.csv"

total_cpu=0
total_memory=0
count=0

echo "time,cpu %,memory %" | tee -a $OUTPUT_FILE

for ((i=1; i<DURATION; i+=INTERVAL)); do
    cpu=$(docker stats --no-stream --format "{{.CPUPerc}}" $CONTAINER_ID | tr -d '%')
    memory=$(docker stats --no-stream --format "{{.MemPerc}}" $CONTAINER_ID | tr -d '%')

    echo "$(date '+%Y-%m-%d %H:%M:%S'),$cpu,$memory" | tee -a $OUTPUT_FILE

    total_cpu=$(echo "$total_cpu + $cpu" | bc)
    total_memory=$(echo "$total_memory + $memory" | bc)

    count=$((count+1))
done
