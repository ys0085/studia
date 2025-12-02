#!/bin/bash

let old_rec=0
let old_tra=0

let curr_rec_sp=0
let curr_tra_sp=0

let prev_rec_sp=0
let prev_tra_sp=0

let pprev_rec_sp=0
let pprev_tra_sp=0

let ppprev_rec_sp=0
let ppprev_tra_sp=0

update_network_data() {
    if [ ${old_rec} -eq 0 ]; then
        old_rec=$(echo $(grep $"enp0s3" /proc/net/dev) | awk '{print $2}')
        old_tra=$(echo $(grep $"enp0s3" /proc/net/dev) | awk '{print $10}')
    fi

    local new_rec=$(echo $(grep $"enp0s3" /proc/net/dev) | awk '{print $2}')
    local new_tra=$(echo $(grep $"enp0s3" /proc/net/dev) | awk '{print $10}')


    ppprev_rec_sp=${pprev_rec_sp}
    ppprev_tra_sp=${pprev_tra_sp}

    pprev_rec_sp=${prev_rec_sp}
    pprev_tra_sp=${prev_tra_sp}

    prev_rec_sp=${curr_rec_sp}
    prev_tra_sp=${curr_tra_sp}

    curr_rec_sp=$((new_rec - old_rec))
    curr_tra_sp=$((new_tra - old_tra))

    old_rec=${new_rec}
    old_tra=${new_tra}
}

let MAX_BAR=3000000

let BAR_INTERVAL=50000

draw_bar() {
    echo -n "| "
    for ((i = 0; i < MAX_BAR; i += BAR_INTERVAL)); do
        if (($1 > i)); then
            echo -n "█"
        else
            echo -n " "
        fi
    done
    echo ""
}

graph_newtork_data() {
    echo "| Download speed"
    draw_bar "$ppprev_rec_sp"
    draw_bar "$pprev_rec_sp"
    draw_bar "$prev_rec_sp"
    draw_bar "$curr_rec_sp"
    echo "| Current: $((curr_rec_sp / 1000)) Mb/s"
    echo ""
    echo "| Upload speed"
    draw_bar "$ppprev_tra_sp"   
    draw_bar "$pprev_tra_sp"
    draw_bar "$prev_tra_sp"
    draw_bar "$curr_tra_sp"
    echo "| Current: $((curr_tra_sp / 1000)) Mb/s"
}

# The VM I'm using doesn't provide info on max CPU speed. 
# The suggested directory "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq" doesn't exist either.
print_cpu_data() {
    local current_speed=$(echo $(grep 'cpu MHz' /proc/cpuinfo) | awk '{print $4}') 
    echo "CPU0 Speed: ${current_speed} MHz"
}

print_uptime() {
    local uptime=$(cat /proc/uptime | awk '{print $1}')
    echo "Uptime: $(date --date=@${uptime} -u "+%H:%M:%S")"
}

print_loadavg() {
    local loadavg1=$(cat /proc/loadavg | awk '{print $1}')
    local loadavg5=$(cat /proc/loadavg | awk '{print $2}')
    local loadavg15=$(cat /proc/loadavg | awk '{print $3}')
    echo "System load average: "
    echo " - over 1 minute: ${loadavg1}"
    echo " - over 5 minutes: ${loadavg5}"
    echo " - over 15 minutes: ${loadavg15}"
}

print_memory_info() {
    local memfree=$(echo $(grep 'MemFree:' /proc/meminfo) | awk '{print $2}')
    local memtotal=$(echo $(grep 'MemTotal:' /proc/meminfo) | awk '{print $2}')
    local memused=0
    local perc=0
    (( memused = memtotal - memfree ))
    (( perc = 100 * memused / memtotal ))
    echo "Memory usage: ${memused} kB out of ${memtotal} kB used (${perc}%)"
}

main() {
    while true; do
        clear
        update_network_data
        graph_newtork_data
        echo ""
        print_cpu_data
        print_uptime
        echo ""
        print_loadavg
        echo ""
        print_memory_info
        sleep 1
    done
}

main