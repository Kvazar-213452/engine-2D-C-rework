#include <fstream>
#include <sstream>
#include <iostream>
#include <regex>
#include <iostream>
#include <Windows.h>

typedef int (*FindFreePortFunc)();

std::string read_file_html(const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) {
        std::cerr << "Error: Could not open file " << filename << std::endl;
        return "";
    }

    std::stringstream buffer;
    buffer << file.rdbuf();
    return buffer.str();
}



std::string read_file_jsaw(const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) {
        std::cerr << "Error: Could not open file " << filename << std::endl;
        return "";
    }

    std::stringstream buffer;
    buffer << file.rdbuf();
    std::string content = buffer.str();

    std::regex re(R"(<%\s*\(\s*'([^']*)'\s*\)\s*%>)");
    std::smatch match;

    while (std::regex_search(content, match, re)) {
        std::string include_filename = match[1].str();
        std::string include_content = read_file_jsaw(include_filename);
        content.replace(match.position(0), match.length(0), include_content);
    }

    return content;
}



int getFreePort() {
    HMODULE hDll = LoadLibraryA("dynamics_lib/find_free_port.dll");
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