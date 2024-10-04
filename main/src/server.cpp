#include <fstream>
#include <sstream>
#include "../lib/httplib.h"
#include <iostream>
#include "include/config.h"
#include "include/html.h"
#include <filesystem>
#include <locale>
#include <codecvt>

std::string read_file(const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) {
        std::cerr << "Error: Could not open file " << filename << std::endl;
        return "";
    }

    std::stringstream buffer;
    buffer << file.rdbuf();
    std::string content = buffer.str();

    // Шаблон для пошуку тегів <% ('файл') %>
    std::regex re(R"(<%\s*\(\s*'([^']*)'\s*\)\s*%>)");
    std::smatch match;

    // Шукати всі збіги
    while (std::regex_search(content, match, re)) {
        // Отримати шлях до файлу із шаблону
        std::string include_filename = match[1].str();

        // Зчитати вміст файлу
        std::string include_content = read_file(include_filename);

        // Замінити шаблон на вміст файлу
        content.replace(match.position(0), match.length(0), include_content);

        // Оновити пошук після заміни
    }

    return content;
}

void start_server() {
    httplib::Server svr;

    // Встановлення статичної папки для css, js, зображень і т.д.
    svr.set_mount_point("/css_st", "front/static/css");
    svr.set_mount_point("/lib_st", "front/static/lib");
    svr.set_mount_point("/js_st", "front/static/js");  // Додаємо мапінг для статичних файлів

    // Читання index.html
    std::string html_content = read_file("front/index.html");

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