#include <cstdio>

volatile int x = 0;

int main()
{
    printf("Before infinite loop\n");
    while (!x) { }
    printf("After infinite loop\n");
    return 0;
}
