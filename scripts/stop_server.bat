for /f "tokens=5" %%a in ('netstat -aon ^| find ":%1" ^| find "LISTENING"') do (
    taskkill /f /pid %%a 
    goto :EOF
)