@echo off
REM Build script for Windows - builds Go binaries and copies test files

REM Create build directory if it doesn't exist
if not exist build mkdir build

REM Build main Go binary
cd src
call go build -o ..\build\serpent .

REM Build Windows executable
set GOOS=windows
set GOARCH=amd64
call go build -o ..\build\serpent.exe .

cd ..

REM Copy test files to build directory
xcopy test build /E /I /Y /Q

echo Build complete.
