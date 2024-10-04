#include "../lib/httplib.h"
#include "include/config.h"
#include "include/routes.h"

void start_server() {
    httplib::Server svr;

    svr.set_mount_point("/css_st", "front/static/css");
    svr.set_mount_point("/lib_st", "front/static/lib");
    svr.set_mount_point("/js_st", "front/static/js");

    setup_routes(svr);

    svr.listen("127.0.0.1", port);
}