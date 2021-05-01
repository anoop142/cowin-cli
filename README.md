### INTRODUCTON
cowin-cli is a simple cli tool to list centres available for scheduling vaccination  in India with their info. It uses the offical  api used by the cowin portal.

### BUILD

    go build

### USAGE

    cowin-cli -s state -d district [-v vaccine ] [-i] [-b] [-c dd-mm-yyyy]

    cowin-cli -p pincode


example :
        
        cowin-cli -s kerala -d alappuzha 

        cowin-cli -p 688003 -v covishield

By default we use tomorrow's date like the cowin portal.
Not all states are implemented yet.
You can contribute district id of other states, take states/kerala.go as template.


