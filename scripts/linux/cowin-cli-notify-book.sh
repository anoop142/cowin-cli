#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli
# uses script book.sh for booking

# Interval in seconds
T=15
# centers for grep matching
CENTERS='center1|center2|cnter3'

DISTRICT="district"
STATE="state"

notify(){
	notify-send -u critical -t 6000 "Center found!"&
	paplay /usr/share/sounds/freedesktop/stereo/complete.oga&
	echo "$1" $(date) >>  log.txt
}

schedule(){
	./book.sh
}


while :
do
	data=$(./cowin-cli -s $STATE  -d $DISTRICT -b  | grep -iE "$CENTERS" )
	(( $? == 0  )) && echo "$data" && notify "$data" &&  schedule
	sleep $t
done
