#ifndef ROUTES_H
#define ROUTES_H

#include "../../lib/httplib.h"
#include <string>

void setupRoutes(httplib::Server& svr, const std::string& html_content, const std::string& launcher_page);

#endif