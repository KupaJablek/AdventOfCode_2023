#include "d6.h"

#define min_len 20

char **map;
int length = 0;
int width = 0;

int guard_x = 0;
int guard_y = 0;

int parse_file(FILE *file_pointer) {
    char *line = NULL;
    int cur_line = 0;
    ssize_t read = 0;
    size_t len = 0;

    map = (char**)malloc(sizeof (char*) * min_len);


    while((read = getline(&line, &len, file_pointer)) != -1) {
        ++cur_line;
        if (read != width) {
            width = read;
        }

        if (cur_line > length) {
            length += 10;
            map = (char**)realloc(map, sizeof(char*) * length);
        }

        // malloc a new line
        map[cur_line - 1] = (char*)malloc(sizeof (char) * width);

        for (int i = 0; i < read; ++i) {
            if (line[i] == '^') {
                // found the guard
                guard_x = cur_line - 1;
                guard_y = i;
                line[i] = '.';
            }

            map[cur_line-1][i] = line[i];
        }
    }
    return cur_line;
}

enum dir {
    UP, // towards lower index of row
    RIGHT,
    DOWN,
    LEFT,
};

void print_map() {
    for (int i = 0; i < length; ++i) {
        for (int k = 0; k < width; ++k) {
            printf("%c", map[i][k]);
        }
        printf("\n");
    }
    printf("\n");
}

int traverse(int direction, int x, int y) {
    //printf("x:[%d] y:[%d]\n", x, y);
    int x_t = x;
    int y_t = y;

    switch(direction) {
        case UP:
            x_t--;
            break;
        case RIGHT:
            y_t++;
            break;
        case DOWN:
            x_t++;
            break;
        case LEFT:
            y_t--;
            break;
    }

    if (x_t < 0 || x_t >= length) {
        return 1;
    }

    if (y_t < 0 || y_t >= width) {
        return 1;
    }

    // check for collision
    if (map[x_t][y_t] == '#') {
        direction++;
        if (direction == 4) {
            direction = UP;
        }
        return traverse(direction, x, y);
    } 

    if (map[x_t][y_t] == 'X') {
        return traverse(direction, x_t, y_t);
    }

    map[x_t][y_t] = 'X';

    return 1 + traverse(direction, x_t, y_t);
}

void d6() {
    // try to load in the file
    FILE *file_pointer;

    //file_pointer = fopen("./inputs/d6_test.txt", "r");
    file_pointer = fopen(get_path(6), "r");
    if (file_pointer == NULL) {
       printf("Error accessing file [%s]", get_path(6)); 
       return;
    } 

    length = parse_file(file_pointer);

    map[guard_x][guard_y] = 'X';
    int tot1 = traverse(UP, guard_x, guard_y);
    printf("Part 1 total: %d\n", tot1);
}