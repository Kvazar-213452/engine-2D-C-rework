#include "include/config.h"
#include "include/func.h"

int window_h = 1000;  
int window_w = 600;

int port = getFreePort();

const char* name_app = "Gestalt";

std::string html_content_core = R"(
<style>iframe{position: fixed;height: 100%;width: 100%;top: 0%;left: 0%;}</style><iframe src="http://127.0.0.1:)"
+ std::to_string(port) + R"("/" frameborder="0"></iframe>
)";