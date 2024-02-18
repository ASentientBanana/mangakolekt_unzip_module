#include "main.h"
#include <android/log.h>
#include <stdio.h>

// force gcc to link in go runtime (may be a better solution than this)
void force_link() {
  Unzip_Covers(NULL, NULL, NULL);
  Unzip_Single_book(NULL, NULL);
  Check_For_Lib_dir(NULL);
}

int main() {}