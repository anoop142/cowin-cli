### INTRODUCTON
cowin-cli is a simple cli tool to book vaccines as well as list centres available for scheduling vaccination  in India with their info. It uses the offical  api used by the cowin portal.

### BUILD

    go build

### USAGE

There are two modes of operation

1. ### List vaccine centers

    cowin-cli -s state -d district [-v vaccine ] [-i] [-b] [-c dd-mm-yyyy]

    cowin-cli -p pincode


### Example :
        
        cowin-cli -s kerala -d alappuzha 

        cowin-cli -p 688003 -v covishield
### Output
        
    Aroor FHC
        03-05-2021 - 0   45+
    Ala PHC
        03-05-2021 - 0   45+
    
 2. ### Book vaccine
 
    
    cowin-cli -sc -s state -d district [-no mobileNo] [-name Name] 

### Example:

    cowin-cli -sc -s kerala -d alappuzha -no 9123456780
### Output
    +----+----------------+-----------+---------+-----------+
    | ID |   CENTER       | FREE TYPE | MIN AGE |    DATE   |
    +----+----------------+-----------+---------+-----------+
    |  0 | Aroor FHC      | Free      |      45 | 03-05-2021|
    |  1 | Ala PHC        | Free      |      45 | 03-05-2021|
    |  2 | Chunakkara CHC | Free      |      45 | 03-05-2021|
    +----+-----------------------------+-----------+--------+

        Enter Center ID : 1
        Enter OTP : xxxxx

    +----+---------------+
    | ID |     NAME      |
    +----+---------------+
    |  0 | John doe      |
    |  1 | Jane doe      |
    |  2 | All           |
    +----+---------------+

        Enter name ID : 1

        Appointment scheduled successfully!


### Options:
    -b	bookable only
    -c string
            date dd-mm-yyyy (default "03-05-2021")
    -d string
            district
    -i	full info
    -name string
            registered name
    -no string
            mobile number
    -p string
            pincode
    -s string
            state
    -sc
            schedule vaccine
    -v string
            vaccine name
    -version
            version

you can pass -name with "all" to book all registered under same number.
if name is not passed user will prompted to select one.

### Why cowin-cli generates  OTP first and prompt to input after selecting the Center?
This is done to prevent waiting for OTP, it  may take some time to generate otp and receive it, thit time can be saved by sending otp first and entering it after selecting Center.

### Note
By default we use tomorrow's date like the cowin portal.
Not all states are implemented yet.
You can contribute district id of other states, take states/kerala.go as template.


