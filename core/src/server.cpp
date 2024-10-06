#include "../lib/httplib.h"
#include "include/config.h"
#include "include/func.h"
#include "include/routes.h"
#include "include/post.h"

void start_server() {
    httplib::Server svr;

    std::cout << port << std::endl;

    svr.set_mount_point("/css_st", "front/static/css");
    svr.set_mount_point("/lib_st", "front/static/lib");
    svr.set_mount_point("/js_st", "front/static/js");

    std::string main_page = read_file_jsaw("front/main.html");
    std::string launcher_page = read_file_jsaw("front/launcher.html");

    setupRoutes(svr, main_page,
    launcher_page
    );

    setup_post(svr);

    svr.listen("127.0.0.1", port);
}