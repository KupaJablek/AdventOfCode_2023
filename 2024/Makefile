CC = gcc
BIN_NAME = aoc_2024.bin

BIN_PATH = ./bin/

## will define the dependencies
deps = util.o d4.o d6.o days.o aoc_2024.o

util.o: util.c util.h
	$(CC) -c util.c

aoc_2024.o: aoc_2024.c aoc_2024.h
	$(CC) -c aoc_2024.c

day4.o: d4.c d4.h
	$(CC) -c d4.c

day6.o: d6.c d6.h
	$(CC) -c d6.c

days.o: days.c days.h
	$(CC) -c days.c

build: $(deps)
	gcc -o ${BIN_PATH}${BIN_NAME} $(deps)

.PHONY : clean
clean:
	-rm ${BIN_PATH}${BIN_NAME} $(deps)
	-rm $(deps)