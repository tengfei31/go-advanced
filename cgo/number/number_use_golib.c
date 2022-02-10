#include "number.h"
#include <stdio.h>

int main() {
    int a = 1;
    int b = 3;
    int mod = 3;
    int x = number_add_mod(a, b, mod);
    printf("(%d+%d)%%%d = %d\n", a, b, mod, x);
}