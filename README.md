### INTRODUCTON
cowin-cli is a simple cli tool to list centres available for scheduling vaccination  in India with their info. It uses the offical  api used by the cowin portal.

### BUILD

    go build

### USAGE

    cowin-cli -state state -district district [--vaccine vaccine name] [-info ] [-date dd-mm-yyyy]

    cowin-cli -pin pincode --vaccine covishield


example :
        
        cowin-cli -state kerala -district alappuzha 

        cowin-cli -pin 688003 --vaccine covaxin

Not all states are implemented yet.
You can contribute district id of other states, take states/kerala.go as template.


