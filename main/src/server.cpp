#include "../lib/httplib.h"
#include "include/config.h"
#include "include/html.h"
#include "include/func.h"

void start_server() {
    httplib::Server svr;

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