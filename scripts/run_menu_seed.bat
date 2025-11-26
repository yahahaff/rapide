@echo off
echo Running menu seed script...

REM Get MySQL connection details from environment or use defaults
set DB_HOST=%DB_CONNECTION_HOST%
if "%DB_HOST%"=="" set DB_HOST=localhost

set DB_PORT=%DB_CONNECTION_PORT%
if "%DB_PORT%"=="" set DB_PORT=3306

set DB_USER=%DB_CONNECTION_USERNAME%
if "%DB_USER%"=="" set DB_USER=root

set DB_PASS=%DB_CONNECTION_PASSWORD%
if "%DB_PASS%"=="" set DB_PASS=password

set DB_NAME=%DB_CONNECTION_DATABASE%
if "%DB_NAME%"=="" set DB_NAME=rapide

echo Importing menu seed data to %DB_NAME% database...
mysql -h%DB_HOST% -P%DB_PORT% -u%DB_USER% -p%DB_PASS% %DB_NAME% < menu_seed_data.sql

if %ERRORLEVEL% EQU 0 (
    echo Menu seed data imported successfully.
) else (
    echo Failed to import menu seed data. Please check your MySQL connection details.
)

pause