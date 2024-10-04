#ifndef FUNC_H
#define FUNC_H

#include <string>

std::string read_file_html(const std::string& filename);
std::string read_file_jsaw(const std::string& filename);
int getFreePort();

#endif 