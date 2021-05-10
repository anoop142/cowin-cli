
## Command line  tool to List and Book Vaccine
cowin-cli is a simple cli tool to book vaccines and list centers using the COWIN API. It's written in go and works on linux, Windows, Mac and in Android using Termux.

- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Install via `go get`](#install-via-go-get)
  - [Download precompiled binaries](#download-precompiled-binaries)
  - [Android Termux](#android-termux)
- [Getting Started](#getting-started)
  - [List vaccine centers](#list-vaccine-centers)
  - [Book Vaccine](#book-vaccine)
  - [Displaying Captcha](#displaying-captcha)
  - [Termux Auto OTP](#termux-auto-otp)
  - [Options](#options)
    - [List Centers:](#list-center)
    - [Book Vaccine:](#book-vaccine)
- [Known issues](#known-issues)
- [License](#license)


## Installation

### Prerequisites
The following dependencies are 
required for rendering Captcha inside the terminal.This is **optional** for
**Linux and Windows**.
But **required** for android **termux**.

* **[pixterm](https://github.com/eliukblau/pixterm)**

> **Note**: Other terminal based image viewer can also be used, but you have to modify cowin/captcha.go file.
* **[imagemagick](https://imagemagick.org/)**

### Install via `go get`
```bash
$ go get -u github.com/anoop142/cowin-cli
```
### Download precompiled binaries.
Precompiled binaries are avalailable for Windows and linux.
Download them at 
**[Releases](https://github.com/anoop142/cowin-cli/releases)** page.

### Android Termux 
Follow these steps to set up in termux.
```bash
# Install packages
$ pkg i go git imagemagick
# Add go bin to PATH
$ echo 'export PATH=$HOME/go/bin/:$PATH' > ~/.bashrc
$ source ~/.bashrc
#  Install cowin-cli
$ go get -u github.com/anoop142/cowin-cli
# Install pixterm
$ go get -u github.com/eliukblau/pixterm/cmd/pixterm
```


## Getting Started
There are two modes

* List mode
* Booking mode

### **List vaccine centers**

```
cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b]  [-c dd-mm-yyyy]
```
### Example
```console
$ cowin-cli -s kerala -d alappuzha -i -m 45 

Thazhakara PHC  Free  10-05-2021  0  COVISHIELD 45+
Kayamkulam THQH  Free  10-05-2021  0  COVISHIELD 45+
```

The `-i` option displays all extra info like date, vaccine name, age...



### **Book Vaccine**

You can specify mobile number, centers to auto book, age, name etc. 
If not, you will be prompted to enter it appropriately.
```console
$  cowin-cli -sc -state -d district [-no mobileNumber] [-name Name] [-centers center1,cetner2 ] [-slot slotTime]
```
### Example 1
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780

    +----+----------------+-----------+---------+-----------+
    | ID |   CENTER       | FREE TYPE | MIN AGE |  VACCINE  |
    +----+----------------+-----------+---------+-----------+
    |  0 | Aroor FHC      | Free      |      45 | COVISHIELD|
    |  1 | Ala PHC        | Free      |      45 | COVISHIELD|
    |  2 | Chunakkara CHC | Free      |      45 | COVISHIELD|
    +----+----------------+-----------+---------+-----------+

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

     Enter Captcha :  xxxxx

    Appointment scheduled successfully!
```


you can specify most of the details for booking the vaccine

### Example 2
```console
$  cowin-cli -sc -s kerala -d alappuzha -no 9123456780 -name "John doe" -centers "Aroor FHC,Ala PHC"

Center : Aroor FHC
Enter OTP :  xxxxx
Enter Captcha :  xxxxx
```

### Displaying Captcha

* ### Windows
  We use default program to open svg files to display captcha. Usually it's **edge browser**.

* ### Linux
  If **pixterm** and **imagemaick** are installed, captcha is rendered inside the terminal using pixterm.
  if any one of them isn't available, **firefox** is used to display captcha.

* ### Termux
  Without a terminal based image viewer(pixterm) and imagemagick,
  displaying captcha isn't possible in termux.

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



### Options

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
  -p string
        pincode
  -v string
        vaccine name
  -m int
        age
```

#### Book Vaccine:

```
    -sc
            invoke schedule vaccine mode
    -name string
            registered name
    -no string
            mobile number
    -centers string seperated by ','
            centers to auto select
    -m int
            min age limit
    -slot string
            slot time (FORENOON default)
```

## Known issues
* API will throw error unauthorized access if specified vaccine is not found at the moment.
* Random Unauthorized access error for no specific reasons.

## License

GPL 3.0

Copyright (c) Anoop S
