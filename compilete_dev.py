import os
import subprocess

let = input("Type ")

if int(let) == 0:
    os.system("git add -A")
    name = input("Name: ")
    os.system(f'git commit -m "{name}"')
    os.system("git push")
elif int(let) == 1:
    command_s = r"g++ -std=c++17 -O2 -Iinclude -o main.exe main.cpp lib/httplib.cc src/server.cpp src/config.cpp src/routes.cpp src/func.cpp -lole32 -lws2_32 -lShlwapi -lversion"
    path = r'main'

    command = [
        'powershell', 
        '-NoExit',
        '-Command', 
        f'Set-Location -Path "{path}"; {command_s}; ./main; Exit'
    ]

    subprocess.run(command, shell=True)
elif int(let) == 2:
    command_s = r"./main"
    path = r'main'

    command = [
        'powershell', 
        '-NoExit',
        '-Command', 
        f'Set-Location -Path "{path}"; {command_s}; Exit'
    ]

    subprocess.run(command, shell=True)