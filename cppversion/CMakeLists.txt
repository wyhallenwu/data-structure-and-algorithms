cmake_minimum_required(VERSION 3.21)
project(cppversion)

set(CMAKE_CXX_STANDARD 14)

include_directories(stack
        queue)

add_executable(cppversion
#        stack/main.cpp
#        stack/stack.h
        queue/queue.h
        queue/main.cpp)
