#!/bin/bash
# shell script to book vaccine using cowin-cli
# Anoop S


# put name to all if booking for all at once
NAME="name"
STATE="state"
DISTRICT="district"
NO="mobile no"
# centers for cowin-cli to auto select
CENTERS="Center1,Center2"

./cowin-cli -sc -s $STATE -d $DISTRICT -no $NO -name $NAME -centers $CENTERS
