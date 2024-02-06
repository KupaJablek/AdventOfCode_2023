import sys

coords = set()
data = open(sys.argv[1]).read().strip().split('\n')

nRows = len(data)
nCols = len(data[0])


def check_char(nCol, nRow):
    # check if out of bounds
    if not (0 <= nCol < nCols and 0 <= nRow < nRows):
        return False
    return data[nCol][nRow] != "." and not data[nCol][nRow].isdigit()


def main():

    total = 0

    for i, line in enumerate(data):
        startIndex = 0

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

            result = int(num)

            search = startIndex - 1

            # check for special characters around num
            if check_char(i, search) or check_char(i, j):
                total += result
                continue

            # check surrounding rows
            for index in range(search, j + 1):
                if check_char(i-1, index) or check_char(i + 1, index):
                    total += result
                    continue
    print(str(total))


if __name__ == "__main__":
    main()
