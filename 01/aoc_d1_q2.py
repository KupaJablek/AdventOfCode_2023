## aoc day 1 solution p2
import sys

validNums = ['one', 'two', 'three','four','five','six','seven','eight','nine']

def parseFile(fName):
    lines = ''
    try:
        lines = open(fName).read().strip()
    except:
        return -1

    total = 0

    for line in lines.split('\n'):
        digits = [] ## list of all digits found

        for i, char in enumerate(line):
            if char.isdigit():
                digits.append(char)
            for index, num in enumerate(validNums):
                if line[i:].startswith(num):
                    digits.append(str(index+1))
        total += int(digits[0]) * 10 + int(digits[-1]) ## list[-1] -> access last element of list

    return total

def main():
    result = parseFile(sys.argv[1])
    print(result)

if __name__ == "__main__":
    main()