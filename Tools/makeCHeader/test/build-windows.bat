@echo off
rem gcc main.c -o hello
gcc -c main.c
windres appres.rc appres.o
gcc -o hello main.o appres.o -mwindows

pause