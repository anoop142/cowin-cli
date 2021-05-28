
## Command line  tool to List and Book Vaccine
cowin-cli is a simple cli tool to book vaccines and list centers using the COWIN API. It also supports **auto captcha completion** .


>Note: By default cowin-cli will not run continoulsy and monitor slot changes, use bash / batch scripts for that purpose, which can be found [here](#scripts).


## Features
* **Zero dependency** : No neeed to install anything, download precompiled binary and run.
* **Automatic captcha support**: credits to https://github.com/ayushchd
* **Scripting support** : scripts are available for all platforms providing additional features.
* **Reuse OTP** : session token is written to a text file to reuse it later.
* **Advanced Filters**: built-in filter by age, dose, vaccines..etc.
* **Cross platform** : Windows, Linux, macOS, Termux.
* **Automatic OTP support for Termux** 

https://user-images.githubusercontent.com/40671157/119601114-1b2dce00-be06-11eb-923b-3889a24c6ba2.mp4

- [Installation](#installation)
  - [Install via `go get`](#install-via-go-get)
  - [Download precompiled binaries](#download-precompiled-binaries)
  - [Android Termux](#android-termux)
- [Getting Started](#getting-started)
  - [List vaccine centers](#list-vaccine-centers)
  - [Book Vaccine](#book-vaccine)
- [Scripts](#scripts)
- [Advanced](#advanced)  
  - [Termux Auto OTP](#termux-auto-otp)
  - [Generating Token](#generating-token)

- [Options](#options)
    - [List Centers:](#list-center)
    - [Book Vaccine:](#book-vaccine)
- [Known issues](#known-issues)
- [License](#license)


## Installation

### Install via `go get`
```bash
$ go get -u github.com/anoop142/cowin-cli
```
> **Note** : go version 1.16+ is required.

**OR**

### Download precompiled binaries.
Precompiled binaries are avalailable for Windows, macOS and linux.
Download them at 
**[Releases](https://github.com/anoop142/cowin-cli/releases)** page.

### Android Termux 
Follow these steps to set up in termux.
```bash
# Install packages
$ pkg i golang git
# Add go bin to PATH
$ echo -e "export GOPATH=$HOME/go\nexport PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.bashrc
$ source ~/.bashrc
#  Install cowin-cli
$ go get -u github.com/anoop142/cowin-cli
```


## Getting Started
There are two main modes

* List mode
* Booking mode

### **List vaccine centers**

```
cowin-cli -s state -d district [-v vaccine1,vaccine2] [-m age] [-i] [-b]  [-c dd-mm-yyyy] [-dose dose] [-t freeType]
```
### Example 1
```console
$ cowin-cli -s kerala -d alappuzha 

Thazhakara PHC
Kayamkulam THQH  
```

### Example 2
```console
$ cowin-cli -s kerala -d alappuzha -i -m 45 -v "covaxin,covishield" -b -dose 1 -t free

Kalavoor PHC  Free  18-05-2021  11 COVAXIN 45 Dose-1
Vandanam MCH  Free  18-05-2021  4 COVISHIELD 45 Dose-1
Mannanchery PHC  Free  18-05-2021  7 COVISHIELD 45 Dose-1
```

The `-i` option displays all extra info like date, vaccine name, age...
`-b'` prints only bookable centers.



### **Book Vaccine**

You can specify mobile number, centers to auto book, age, name etc. 
If not, you will be prompted to enter it appropriately.
```
$  cowin-cli -sc -state -d district [-no mobileNumber] [-v vaccine1,vaccine2] [-names name1,name2] [-centers center1,cetner2 ] [-slot slotTime] [-ntok]  [-dose dose] [-t freeType]
```
### Example 1
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780

+----+---------------+-----------+---------+-----------+------+
| ID | CENTER        | FREE TYPE | MIN AGE | VACCINE   | DOSE |
+----+---------------+-----------+---------+-----------+------+
|  0 | Aroor FHC     | Free      |      45 | COVISHIELD|   1  |
|  1 | Ala PHC       | Free      |      45 | COVISHIELD|   1  |
|  2 | Kalavoor PHC  | Free      |      45 | COVISHIELD|   2  |
+----+---------------+-----------+---------+-----------+------+

Enter Center ID : 1
Enter OTP : xxxxx

+----+---------------+
| ID |     NAME      |
+----+---------------+
|  0 | John doe      |
|  1 | Jane doe      |
|  2 | Somebody
|  3 | All           |
+----+---------------+

Enter name ID : 1,2

Appointment scheduled successfully!
```
>Note: By default cowin-cli will reuse token, so until the token expires , you don't need to enter otp again.

you can specify most of the details for booking the vaccine

### Example 2
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780 -names "John doe, Jane doe" -centers "Aroor FHC,Ala PHC" -v "covaxin,sputnik v" -dose 2 -t free

Center : Aroor FHC COVAXIN Dose-2
Enter OTP :  xxxxx
```
>**Note**: -centers "any" to auto select any center.
>-name "all" to book for all under same mobile no.


## Scripts
Scripts are available for notifying and booking using cowin-cli [here](scripts). You need to edit the vaules of the script like district name, mobile number etc..


## Advanced
 >**Note**: This is meant for advanced users.

### Termux Auto OTP
It's possible to detect OTP message and get OTP in Termux without user input. use -aotp flag to invoke this feature.

You need to first setup termux to read sms.

  1.Install Termux API apk from Fdroid

  2.Install termux-api package 

  ```bash
  # Install termux-api package
  $ pkg i termux-api
  # To give permisiion
  $ termux-list-sms
  # Example
  $ cowin-cli -s kerala -d alappuzha -sc -no 9123456789 -aotp
  ```

 ### Generating Token

 Tokens are always written to "token.txt" after successfully validating OTP while booking to re-use later.

You can generate token manually using the token generation mode.

```
cowin-cli -gen [-no MobileNumber] [-token tokenFile]
```
### Example
```console
$ cowin-cli -gen -no 9123456789

Enter OTP :xxxxxx
Written to token.txt
```


### **Options**

```
  -s	state State Name
  -d  district District name
  -version	Show version
  -h  Show Help
```

#### List Center:

```
  -b	
        show bookable only
  -c string
        date dd-mm-yyyy (default tomorrow's date)
  -i	
        full info
  -v string
        vaccine names separated by ','
  -m int
        age
  -dose int
            dose type
  -t string
            free type
```

#### Book Vaccine:

```
    -sc
            invoke schedule vaccine mode
    -names string separated by ','
            beneficiaries name           
    -no string
            mobile number
    -centers string separated by ','
            centers to auto select
    -m int
            min age limit
    -slot string
            slot time 
    -v string
            vaccine names separated by ','
    -ntok
            don't reuse token
    -dose int
            dose type
    -token string
            file to write token (default "token.txt")
    -t string
            free type

```

### Generate Token:

```
     -gen
        invoke token generation mode
    -token string
            file to write token (default "token.txt")
     -no int       
            mobile number
```

## Known issues
* Random Unauthenticated access error for no reasons.

## License

GPL 3.0

Copyright (c) Anoop S
