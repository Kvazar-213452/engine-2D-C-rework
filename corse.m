g++ -std=c++17 -O2 -Iinclude -o main.exe main.cpp src/httplib.cc src/server.cpp src/config.cpp src/html.cpp -lshell32 -lole32 -lws2_32 -lversion -lshlwapi