#include "../lib/httplib.h"
#include "include/func.h"

void setup_routes(httplib::Server& svr) {
    std::string html_content = read_file_jsaw("front/main.html");
    std::string launcher = read_file_jsaw("front/launcher.html");

    svr.Get("/", [&html_content](const httplib::Request& req, httplib::Response& res) {
        if (html_content.empty()) {
            res.status = 500;
            res.set_content("Error: Could not read index.html", "text/plain");
        } else {
            res.set_content(html_content, "text/html");
        }
    });

    svr.Get("/main", [&launcher](const httplib::Request& req, httplib::Response& res) {
        res.set_content(launcher, "text/html");
    });
}