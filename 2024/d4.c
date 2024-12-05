#include "d4.h"

#define min_size 20

char **word_search;
int word_search_l = 10;
int word_search_w = 0;


int parse_input(FILE *file_pointer) {
    int cols = min_size;
    int rows = min_size;

    char *line = NULL;
    int cur_line = 0;
    ssize_t read = 0;
    size_t len = 0;

    word_search = (char**)malloc(sizeof (char*) * cols);

    while((read = getline(&line, &len, file_pointer)) != -1) {
        ++cur_line;

        if (read != rows) {
            rows = read;
            word_search_w = read;
        }

        if (cur_line > cols) {
            cols += 10;
            word_search = (char**)realloc(word_search, sizeof(char*) * cols);
        }

        // malloc a new line
        word_search[cur_line - 1] = (char*)malloc(sizeof (char) * rows);

        for (int i = 0; i < read; ++i) {
            word_search[cur_line-1][i] = line[i];
        }
    }

    return cur_line;
}


int search(int col, int row) {
    char *xmas = "XMAS";
    int tot = 0;

    // up
    bool found = true;
    for (int i = 1; i < 4 && col+i < word_search_l; ++i) {
        if (word_search[col+i][row] != xmas[i]) {
           found = false; 
           break; 
        } 
    }

    if(found) {
        ++tot;
    }
    found = true;
    // down
    for (int i = 1; i < 4 && col-i >= 0; ++i) {
        if (word_search[col-i][row] != xmas[i]) {
           found = false; 
           break; 
        } 
    }

    if(found) {
        ++tot;
    }
    found = true;
    // left
    for (int i = 1; i < 4 && row+i < word_search_w; ++i) {
        if (word_search[col][row+i] != xmas[i]) {
           found = false; 
           break; 
        } 
    }

    if(found) {
        ++tot;
    }
    found = true;
    // right
    for (int i = 1; i < 4 && row-i >= 0; ++i) {
        if (word_search[col][row-i] != xmas[i]) {
           found = false; 
           break; 
        } 
    }

    if(found) {
        ++tot;
    }

    return tot;
}

int find_xmas() {
    int tot = 0;
    for (int i = 0; i < word_search_l; ++i) {
        for (int k = 0; k < word_search_w; ++k) {
            if(word_search[i][k] == 'X') {
                // search
                tot += search(i,k);
            }
        }
    }
    return tot;
}

void d4() {
    // try to load in the file
    FILE *file_pointer;

    //file_pointer = fopen("./inputs/d4_test.txt", "r");
    file_pointer = fopen(get_path(4), "r");
    if (file_pointer == NULL) {
       printf("Error accessing file [%s]", get_path(1)); 
       return;
    } 

    word_search_l = parse_input(file_pointer);

    int tot1 = find_xmas();
    printf("P1 total: %d\n", tot1);
}