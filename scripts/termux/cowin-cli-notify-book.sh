#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli
# Install notify-send for notifications


COWIN_CLI="./cowin-cli"

# Interval in seconds
T=15
# centers for grep matching
CENTERS_MATCH='center1|center2|cnter3'
# centers to auto select
CENTERS=""
DISTRICT="district"
STATE="state"
AGE=45
NAME=""
NO=""
# vaccines seperated by ','
VACCINE=""



schedule(){
	"$COWIN_CLI" -s "$STATE" -d "$DISTRICT" -sc -no "$NO" -name "$NAME" -centers "$CENTERS" -v "$VACCINE" -aotp && exit 0 
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
