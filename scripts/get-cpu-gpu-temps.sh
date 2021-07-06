#!/bin/bash

# rvcgencmd requires sudo apt install libraspberrypi-bin
cpu=$(</sys/class/thermal/thermal_zone0/temp)

cat <<EOF 
reading:
    date: $(date) 
    gpu: $(vcgencmd measure_temp | egrep -o '[0-9]*\.[0-9]*')
    cpu: $((cpu/1000))
EOF
