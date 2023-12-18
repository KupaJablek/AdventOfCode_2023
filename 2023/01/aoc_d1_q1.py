# aoc day 1 solution
import sys


def parseFile(fName):
    lines = []

    try:
        lines = open(fName).read().strip()
    except:
        return -1

    total = 0;

    for line in lines.split('\n'):
        digits = []

        for character in line:
            if character.isdigit():
                digits.append(character)

        total += int(digits[0]) * 10 + int(digits[-1])
    return total


def main():
    result = parseFile(sys.argv[1])
    print(result)


if __name__ == "__main__":
    main()
