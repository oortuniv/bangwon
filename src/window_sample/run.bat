@echo off

START /B CMD /C CALL "C:\dev\bangwon\sh\window_sample\test.exe" > NUL
exit %errorlevel%