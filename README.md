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
        
        cowin-cli -s kerala -d alappuzha -i

        cowin-cli -p 688003 -v covishield -i
### Output
        
    Aroor FHC
        03-05-2021 - 0   45+
    Ala PHC
        03-05-2021 - 0   45+
    
2. ### Book vaccine
 
    
     cowin-cli -sc -state -d district [-no mobileNumber] [-name Name] [-centers center1,cetner2 ]

### Example 1:

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

### Example 2:

        cowin-cli -sc -s kerala -d alappuzha -no 9123456780 -name "John doe" -centers "Aroor FHC,Ala PHC"

### Options:
    -b	print bookable only
    -c string
            date dd-mm-yyyy (default tomorrow's date)
    -centers string
    	centers to auto book seperated by ','
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
This is done to prevent waiting for OTP. It  may take some time to generate OTP and receive it, this time can be used for selecting the center and enter OTP after it.
### Note
By default we use tomorrow's date like the cowin portal.
Not all states are implemented yet.
You can contribute district id of other states, take states/kerala.go as template.


