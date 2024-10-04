#include "../lib/httplib.h"
#include "include/config.h"
#include <string>

void setupRoutes(httplib::Server& svr, const std::string& html_content, const std::string& launcher_page) {
    svr.Get("/", [&html_content](const httplib::Request& req, httplib::Response& res) {
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    svr.Get("/launcher", [&launcher_page](const httplib::Request& req, httplib::Response& res) {
        res.set_content(launcher_page, "text/html");
    });
}