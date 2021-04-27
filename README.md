### INTRODUCTON
cowin-cli is a simple cli tool to list centres available for scheduling vaccination  in India with their info. It uses the offical  api used by the cowin portal.

### BUILD

    go build

### USAGE

    cowin-cli -state state -district district [-info ] [-date dd-mm-yyyy]

example :
        
        cowin-cli -state kerala -district alappuzha 

Not all states are implemented yet.
You can contribute district id of other states, take states/kerala.go as template.


