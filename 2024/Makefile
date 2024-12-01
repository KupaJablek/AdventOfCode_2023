CC = gcc
BIN_NAME = aoc_2024.bin

BIN_PATH = ./bin/

## will define the dependencies
deps = aoc_2024.o days.o

aoc_2024.o: aoc_2024.c aoc_2024.h
	$(CC) -c aoc_2024.c

days.o: days.c days.h
	$(CC) -c days.c

build: $(deps)
	gcc -o ${BIN_PATH}${BIN_NAME} $(deps)

.PHONY : clean
clean:
	-rm ${BIN_PATH}${BIN_NAME} $(deps)
	-rm $(deps)