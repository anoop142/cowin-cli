:: Auto covid vaccine booking script using cowin-cli
:: Anoop S

@echo off

set COWIN-CLI=.\cowin-cli.exe 

:: set this to 1 to use protected API
set /A PROTECTED_API=0


::------------------------------------------------
:: NEED EDIT
::------------------------------------------------
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

:: dose, 0 means both doses 1 & 2
set /A DOSE=0
::------------------------------------------------



::------------------------------------------------
:: Optionaly Edit
::------------------------------------------------

:: centers name to auto select, should be accurate, seperated by ',' . 
set CENTERS=""

:: vaccines seperated by ','
set VACCINE=""

:: no need to edit this, auto defaults to tomorrow's date
set DATE=""

:: center type, free or paid, empty means all
set TYPE=""

:: minimum slots 
set /A MIN_SLOT=1

::END OF EDITABLE VALUES
::------------------------------------------------



:: Main looping  function
:loop
echo looking for centers...
IF NOT %PROTECTED_API%==0 (call:book) ELSE (call:list && call:book)
timeout /t %INTERVAL% >nul
goto loop


:: Booking function
:book
%COWIN-CLI% -s %STATE% -d %DISTRICT% -sc -no %NO% -names %NAMES% -centers %CENTERS% -v %VACCINE% -dose %DOSE% -c %DATE% -t %TYPE% -ms %MIN_SLOT%
goto:eof

:: Listing function
:list
%COWIN-CLI% -s %STATE% -d %DISTRICT% -m %AGE% -b -v %VACCINE% -dose %DOSE% -c %DATE% -t %TYPE% -ms %MIN_SLOT%
goto:eof


