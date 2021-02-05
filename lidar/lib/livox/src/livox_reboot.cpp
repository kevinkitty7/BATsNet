//
// The MIT License (MIT)
//
// Copyright (c) 2019 Livox. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

#include "cmdline.h"
#include "lds_lidar.h"
#include <chrono>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <thread>

int main(int argc, const char *argv[]) {

  std::string code(argv[1]);
  std::vector<std::string> cmdline_broadcast_code = {code};
  LdsLidar &read_lidar = LdsLidar::GetInstance();

  int ret = read_lidar.InitLdsLidar(cmdline_broadcast_code);
  if (!ret) {
    printf("Init lds lidar success!\n");
  } else {
    printf("Init lds lidar fail!\n");
  }

  printf("Start discovering device.\n");

  std::this_thread::sleep_for(std::chrono::seconds(100));

  read_lidar.RebootAllConnected();

  read_lidar.DeInitLdsLidar();
  printf("Livox lidar demo end!\n");
}