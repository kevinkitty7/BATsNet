cmake_minimum_required( VERSION 3.9 )
project(service_discovery VERSION 0.0.2 LANGUAGES C CXX)

find_package(Protobuf REQUIRED)

# EXEC_PROGRAM(protoc 
#   ARGS -I=${CMAKE_CURRENT_SOURCE_DIR}/proto
#        --cpp_out=${CMAKE_CURRENT_SOURCE_DIR}/proto
#        ${CMAKE_CURRENT_SOURCE_DIR}/proto/*.proto)

include_directories(${CMAKE_SOURCE_DIR})
include_directories(${GLOBAL_HEADER}/third_party/etcd-cpp-apiv3)

add_library(service_node SHARED src/service_node.cc)
set_target_properties(service_node PROPERTIES VERSION ${PROJECT_VERSION})
target_include_directories(service_node PUBLIC ./ ${GLOBAL_HEADER}/third_party/etcd-cpp-apiv3)
target_link_libraries(service_node etcd-cpp-api)

add_library(service_discovery SHARED src/service_discovery.cc)
set_target_properties(service_discovery PROPERTIES VERSION ${PROJECT_VERSION})
target_include_directories(service_discovery PUBLIC ./ ${GLOBAL_HEADER}/third_party/etcd-cpp-apiv3)
target_link_libraries(service_discovery etcd-cpp-api)

add_executable(test_node src/test_node.cc)
target_link_libraries(test_node service_node)

add_subdirectory(binding)

install(TARGETS service_node DESTINATION ${BATSNET_LIB_INSTALL_DIR})
install(TARGETS service_discovery DESTINATION ${BATSNET_LIB_INSTALL_DIR})
