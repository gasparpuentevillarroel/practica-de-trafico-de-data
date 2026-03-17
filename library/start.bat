@echo off
cd /d "%~dp0"
start cmd /k "cd back && go run ."
start cmd /k "cd front && npm run dev"
