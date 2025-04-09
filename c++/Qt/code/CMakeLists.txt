cmake_minimum_required(VERSION 3.10)
project(QtStudy)

set(CMAKE_CXX_STANDARD 11)

include_directories(/usr/local/bin/Qt/include)

#TARGET_LINK_LIBRARIES (QtStudy ${QT_LIBRARIES})

add_executable(QtStudy main.cpp myWidget.h myWidget.cpp)