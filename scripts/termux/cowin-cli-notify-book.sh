#!/bin/bash
# Anoop S
# shell script to notify and book using cowin-cli


COWIN_CLI="cowin-cli"

# set this to 1 for protected API
PROTECTED_API=0

#--------------------------------------------
# NEED EDITING
#--------------------------------------------
STATE="tamil nadu"

DISTRICT="chennai"

AGE=45

# 0 means both dose 1 and 2
DOSE=0

# mobile number
NO=""

# beneficiaries names seperated by ','
NAMES=""
#--------------------------------------------


#-------------------------------------------
# OPTIONAL
#-------------------------------------------
# Loop Interval in seconds
T=15

# centers to auto select seperated by ','
CENTERS=""

# vaccines seperated by ',', default all
VACCINE=""

# no need to edit this, defaults to tomorrow's date 
DATE=""

# free type, free or paid, default all
TYPE=""
#-------------------------------------------



list(){
	"$COWIN_CLI" -s "$STATE"  -d "$DISTRICT" -m "$AGE" -b \
	 -v "$VACCINE" -dose "$DOSE" -c "$DATE" -t "$TYPE"

}


schedule(){
	"$COWIN_CLI" -s "$STATE" -d "$DISTRICT" -sc -no "$NO" -names "$NAMES" -m "$AGE" \
	-centers "$CENTERS" -v "$VACCINE" -dose "$DOSE"   -c "$DATE" -t "$TYPE" -aotp
}



# main function
while :
do
	if [[ $PROTECTED_API  -ne 0 ]]
	then
		schedule
	else
		echo "looking for centers.."
		list && schedule && exit
	fi

	sleep $T
done
