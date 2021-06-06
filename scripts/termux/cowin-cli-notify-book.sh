#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli
# Install notify-send for notifications


COWIN_CLI="cowin-cli"

# Interval in seconds
T=15
# centers for grep matching
CENTERS_MATCH='center1|center2|cnter3'
# centers to auto select seperated by ','
CENTERS=""
DISTRICT="district"
STATE="state"
AGE=45
# beneficiaries names seperated by ','
NAMES=""
NO=""
# vaccines seperated by ','
VACCINE=""
DOSE=0
DATE=""
# free type, free or paid, default all
TYPE=""



schedule(){
	# pass -aotp if you have setup auto OTP  
	"$COWIN_CLI" -s "$STATE" -d "$DISTRICT" -sc -no "$NO" -names "$NAMES" -m "$AGE" \
	-centers "$CENTERS" -v "$VACCINE" -dose $DOSE   -c "$DATE" -t "$TYPE"  && exit 0 
}


while :
do
	echo "looking for centers.."

	"$COWIN_CLI" -s "$STATE"  -d "$DISTRICT" -m "$AGE" -b -v "$VACCINE" -dose $DOSE -c "$DATE" -t "$TYPE" -public

	if (( $? == 0  )) 
	then
		schedule
	fi

	sleep $T
done
