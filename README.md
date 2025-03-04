# What is it and why
Tired of Roblox sneaking onto your PC? This program will help! Simply download it and add it to startup, and it will prevent Roblox from launching on your computer.
Go_RobloxKiller is a program that starts and runs in the background, preventing you from launching Roblox on your PC!
( I was asked to make this programm )
## Installation
You can build the .exe by yourself or install a [ready-to-use version](https://github.com/Wrtgvr2/Go_RobloxKiller/releases/tag/v1)
## Build
Do it, if you don't want to install ready-to-use version
```bash
# Navigate to the folder where you want to clone the repository
cd [folder where to clone]
# Clone the repository
git clone https://github.com/Wrtgvr2/Go_RobloxKiller
```
Now you need to install Go to build .exe file.
You can install Go via [official website](https://go.dev/dl/)

**Important!**
This command will build .exe file, and you can change name for the file, but if you do, you will need to remember name you give it!
Otherwise, you may not be able to find the file in startup :X
Name the file whatever you want. If you want to name it "CoolProgramm123", then write "CoolProgramm123.exe" instead of "roblox_killer.exe"
```bash
# Build the executable file
# The flag "-H windowsgui" ensures that the program runs without showing a console window
go build -ldflags "-H windowsgui" -o roblox_killer.exe
```
After successful build you will be able to find .exe file in the folder with main.go file
## How to put it in startup
To make the program run automatically when you start your PC, follow these steps:
1. Press `Win + R`
2. Type `shell:startup`
3. Put `.exe` file into the folder
Now, the program will start in the background each time your PC boots up!
## How to uninstall?
Uninstalling the program is simple:
1. Press `Ctrl + Shift + Esc` to open Task Manager.
2. Find the running `.exe` process and end it.
3. Press `Win + R` to open the Run dialog.
4. Type `shell:startup` and press Enter to open the Startup folder.
5. Remove the `.exe` file from the folder.
