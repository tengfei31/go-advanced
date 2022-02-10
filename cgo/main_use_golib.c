#include "main.h"
#include <stdio.h>

int number_add_mod(int, int, int);

int UseGoLibcMain() {
    int a = 10;
    int b = 5;
    int c = 12;

    int x = number_add_mod(a, b, c);
    printf("(%d+%d)%%%d = %d\n", a, b, c, x);

    goPrintln("done");
}