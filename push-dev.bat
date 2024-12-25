@echo off

:retry

echo pushing...

git push --thin --progress "origin" dev

if %errorlevel% NEQ 0 (
  goto retry
)
pause