#include "../lib/httplib.h"

void setup_post(httplib::Server& svr) {    
    svr.Post("/save_file", [](const httplib::Request& req, httplib::Response& res) {
        res.set_content("Invalid JSON", "text/plain");
    });
}
