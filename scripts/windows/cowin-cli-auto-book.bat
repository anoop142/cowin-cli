:: Auto covid vaccine booking script using cowin-cli
:: Anoop S

@echo off

set COWIN-CLI=.\cowin-cli.exe 

:: EDIT VAULES 

:: time interval to check in seconds ,DO NOT  SET IT LESS THAN 10S
set /A INTERVAL=11
:: Age limit of centers
set /A AGE=45
:: State name
set STATE="tamil nadu"
:: District name
set DISTRICT="chennai"
:: Name should be accurate , cases should be correct
set NAME=""
:: Mobile number
set NO=""
:: 	centers name pattern seperated by space, name is not required to be accurate, eg: ala kayam 
set CENTERS_MATCH=""
:: centers name to auto select, should be accurate, seperated by ',' . 
set CENTERS=""
:: vaccines seperated by ','
set VACCINE=""

::END OF EDITABLE VALUES



:: Main looping  function
:loop
echo looking for centers...
call:list && call:book
timeout /t %INTERVAL% >nul
goto loop


:: Booking function
:book
%COWIN-CLI% -s %STATE% -d %DISTRICT% -sc -no %NO% -name %NAME% -centers %CENTERS% -v %VACCINE%
pause
exit

:: Listing function
:list
IF [%CENTERS_MATCH%]==[""] (
   %COWIN-CLI% -s %STATE% -d %DISTRICT% -m %AGE% -b -v %VACCINE%
 ) ELSE ( 
    %COWIN-CLI% -s %STATE% -d %DISTRICT% -m %AGE% -b -v %VACCINE% | findstr /I %CENTERS_MATCH%
 )
goto:eof


