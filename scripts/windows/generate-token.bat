:: Auto covid vaccine booking script using cowin-cli
:: Anoop S

@echo off

set COWIN-CLI=.\cowin-cli.exe
:: mobile number
set NO=""


%COWIN-CLI% -gen -no %NO%
pause
exit

