#include "days.h"
#define min_numbers 5

char* get_path(int day) { char *path = malloc(sizeof (char) * 20);

    sprintf(path, "./inputs/d%d.txt", day);
    return path;
}
int num_numbers = 5;
int *list1;
int *list2;

int parse(FILE* file_p) {
    char *line = NULL;
    int cur_line = 0;
    ssize_t read = 0;
    size_t len = 0;

    while ((read = getline(&line, &len, file_p)) != -1) {
        ++cur_line;

        int num1 = 0;
        int num2 = 0;
        int idx = 0;

        // first num is up to double space
        for (idx = 0; idx < read && line[idx] != ' '; ++idx) {
            num1 = num1 * 10 + line[idx] - '0';
        }

        // skip over whitespace
        idx+=2;

        // second num is from whitespace to end of string
        char * second_int = line;
        second_int += idx;
        num2 = atoi(second_int);

        // make space if needed
        if (cur_line > num_numbers) {
            num_numbers += 100;
            list1 = realloc(list1, sizeof (int) * (num_numbers));
            list2 = realloc(list2, sizeof (int) * (num_numbers));
        }

        list1[cur_line - 1] = num1;
        list2[cur_line - 1] = num2;
    }

    // trim unused space
    list1 = realloc(list1, sizeof(int) * (cur_line));
    list2 = realloc(list2, sizeof(int) * (cur_line));

    return cur_line;
}

int compare(const void* a, const void* b) {
    return (*(int*)a - *(int*)b);
}

unsigned int sort_and_compare() {
    qsort(list1, num_numbers, sizeof(int), compare);
    qsort(list2, num_numbers, sizeof(int), compare);

    unsigned int total = 0;

    for (int i = 0; i < num_numbers; ++i) {
        total += abs(list1[i] - list2[i]);
    }

    return total;
}

// O(n^2)... would have been O(n) with hashmap...
unsigned int similarity_score() {
    unsigned int total = 0;
    for (int i = 0; i < num_numbers; i++) {
        int target = list1[i];
        int count = 0;

        for (int k = 0; k < num_numbers; k++) {
           if (list2[k] == target) {
            count++;
           } 
        }
        total += target * count;
    }
    return total;
}

void d1() {
    FILE *file_pointer;

    file_pointer = fopen(get_path(1), "r");
    if (file_pointer == NULL) {
       printf("Error accessing file [%s]", get_path(1)); 
       return;
    } 

    list1 = malloc(sizeof (int) * min_numbers);
    list2 = malloc(sizeof (int) * min_numbers);

    // read in the input file
    num_numbers = parse(file_pointer);

    unsigned int tot1 = sort_and_compare();
    printf("Total for p1: %d\n", tot1);
    printf("Total for p2: %d\n", similarity_score());
}
