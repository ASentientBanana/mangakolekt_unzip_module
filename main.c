#include "manga_archive.h"
#include <stdio.h>

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
  Unzip_Single_book(NULL, NULL);
  Unzip_Single_book(NULL, NULL);
  Unrar_Single_book(NULL, NULL);
}

int main() {}