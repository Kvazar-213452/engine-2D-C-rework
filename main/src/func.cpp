#include <fstream>
#include <sstream>
#include <iostream>
#include <regex>

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