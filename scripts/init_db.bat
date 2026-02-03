@echo off
REM Navigate to project root
cd /d "%~dp0\.."

REM Run the initialization script
echo Initializing database with default data...
go run cmd/migrate/init_data.go

echo Done!
pause