#include <stdio.h>
#include "main.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
    Unzip(NULL,NULL,NULL);
    Unzip_Single_book(NULL,NULL);
}

int main() {

}