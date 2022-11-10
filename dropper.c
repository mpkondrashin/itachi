#include <stdio.h>
#include <stdlib.h>

const char *eicar = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*";
const char *filename = "eicar.com";
int main() {
    FILE *f = fopen(filename, "w");
    if (f == NULL) {
        perror(filename);
        exit(1);
    }
    fputs(eicar, f);
    return fclose(f);
}