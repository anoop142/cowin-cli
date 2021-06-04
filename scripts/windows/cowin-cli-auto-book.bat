:: Auto covid vaccine booking script using cowin-cli
:: Anoop S

@echo off

set COWIN-CLI=.\cowin-cli.exe 

:: EDIT VAULES 

:: time interval to check in seconds ,DO NOT  SET IT LESS THAN 10S
set /A INTERVAL=15
:: Age limit of centers
set /A AGE=45
:: State name
set STATE="tamil nadu"
:: District name
set DISTRICT="chennai"
:: beneficiaries names separated by ','
set NAMES=""
:: Mobile number
set NO=""
:: centers name pattern seperated by space, name is not required to be accurate, eg: ala kayam 
set CENTERS_MATCH=""
:: centers name to auto select, should be accurate, seperated by ',' . 
set CENTERS=""
:: vaccines seperated by ','
set VACCINE=""
:: dose, 0 means all
set /A DOSE=0
:: date, no need to edit this by default, dd-mm-yyyy format
set DATE=""
:: center type, free or paid, empty means all
set TYPE=""
::END OF EDITABLE VALUES



:: Main looping  function
:loop
echo looking for centers...
call:book
timeout /t %INTERVAL% >nul
goto loop


:: Booking function
:book
%COWIN-CLI% -s %STATE% -d %DISTRICT% -sc -no %NO% -names %NAMES% -centers %CENTERS% -v %VACCINE% -dose %DOSE% -c %DATE% -t %TYPE%
pause
exit


