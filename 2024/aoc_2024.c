#include "aoc_2024.h"

int main(int argc, char *argv[]) {
    
    if (argc != 1 && argc != 2) {
        printf("Expected usage: 2024.bin [day]\n");
        return 1;
    }
    int day = 0;

    // prompt for user input
    if (argc == 1) {
    }

    // otherwise just run the day
    day = atoi(argv[1]);

    switch (day) {
    case 1:
        d1();
        break;
    
    default:
        printf("Day %d has not been implemented yet\n", day); 
        break;
    }

    return 0;
}
