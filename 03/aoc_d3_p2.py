import sys

coords = set()
data = open(sys.argv[1]).read().strip().split('\n')

nRows = len(data)
nCols = len(data[0])

# array of lists to mimic data
gears = [[[] for _ in range(nCols)] for _ in range(nRows)]


def check_char(nCol, nRow, num):
    # check if out of bounds
    if not (0 <= nCol < nCols and 0 <= nRow < nRows):
        return False

    # add number to associated start location
    if data[nCol][nRow] == "*":
        gears[nCol][nRow].append(num)

    return data[nCol][nRow] != "." and not data[nCol][nRow].isdigit()


def main():
    total = 0

    for i, line in enumerate(data):
        startIndex = 0

        # python eqv of good for loop
        j = 0
        while j < nCols:
            startIndex = j
            num = ""

            while j < nCols and line[j].isdigit():
                num += line[j]
                j += 1

            if num == "":
                j += 1
                continue

            # parsed value of number
            result = int(num)

            search = startIndex - 1
            # check for special characters around num
            if check_char(i, search, result) or check_char(i, j, result):
                continue

            # check surrounding rows
            for index in range(search, j + 1):
                if check_char(i-1, index, result) or check_char(i + 1, index, result):
                    continue

    for i in range(nCols):
        for j in range(nRows):
            numbers = gears[i][j]
            if data[i][j] == "*" and len(numbers) == 2:
                total += numbers[0] * numbers[1]

    print(str(total))


if __name__ == "__main__":
    main()
