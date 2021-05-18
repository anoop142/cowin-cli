#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli
# Install notify-send for notifications


COWIN_CLI="./cowin-cli"

# Interval in seconds
T=15
# centers for grep matching
CENTERS_MATCH='center1|center2|cnter3'
# centers to auto select seperated by ','
CENTERS=""
DISTRICT="district"
STATE="state"
AGE=45
NAME=""
NO=""
# vaccines seperated by ','
VACCINE=""





notify(){
	notify-send -u critical -t 6000 "Center found!"&
	paplay /usr/share/sounds/freedesktop/stereo/complete.oga&
}

schedule(){
	"$COWIN_CLI" -s "$STATE" -d "$DISTRICT" -sc -no "$NO" -name "$NAME" -centers "$CENTERS" -v "$VACCINE" && exit 0 
}


while :
do
	echo "looking for centers.."

	"$COWIN_CLI" -s "$STATE"  -d "$DISTRICT" -m "$AGE" -b -v "$VACCINE"

	if (( $? == 0  )) 
	then
		notify
		schedule
	fi

	sleep $T
done
