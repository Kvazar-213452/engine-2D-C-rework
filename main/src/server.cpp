#include "../lib/httplib.h"
#include "include/config.h"
#include "include/html.h"
#include "include/func.h"
#include <iostream>
#include <Windows.h>

typedef int (*FindFreePortFunc)();

int getFreePort() {
    HMODULE hDll = LoadLibrary("dynamics_lib/find_free_port.dll");
    if (hDll == nullptr) {
        std::cerr << "Could not load DLL" << std::endl;
        return -1;
    }

    FindFreePortFunc findFreePort = (FindFreePortFunc)GetProcAddress(hDll, "FindFreePort");
    if (findFreePort == nullptr) {
        std::cerr << "Could not locate the function" << std::endl;
        FreeLibrary(hDll);
        return -1;
    }
    
    int freePort = findFreePort();
    FreeLibrary(hDll);

    return freePort; 
}

void start_server() {
    httplib::Server svr;

    int port = getFreePort();
    std::cout << port << std::endl;

    svr.set_mount_point("/css_st", "front/static/css");
    svr.set_mount_point("/lib_st", "front/static/lib");
    svr.set_mount_point("/js_st", "front/static/js");

    std::string html_content = read_file_jsaw("front/main.html");

    svr.Get("/", [&html_content](const httplib::Request& req, httplib::Response& res) {
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    svr.listen("127.0.0.1", port);
}