### INTRODUCTON
cowin-cli is a simple cli tool to list available vaccine centres in India with their info. It uses the offical cowin api.

### BUILD

    go build

### USAGE

    cowin-cli -state state -district district [-info ] [-date dd-mm-yyyy]

example :
        
        cowin-cli -state kerala -district alappuzha 

Not all states are implemented yet.
You can contribute district id of pther states, take states/kerala.go as template/


