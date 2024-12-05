#include "util.h"

char* get_path(int day) { 
    char *path = malloc(sizeof (char) * 20);

    sprintf(path, "./inputs/d%d.txt", day);
    return path;
}