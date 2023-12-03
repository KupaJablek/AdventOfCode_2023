import sys

def validateNum(data, dIndex, numEndIndex, num):
    value = 0
    validNum = False

    nStart = numEndIndex - len(num)

    line = data[dIndex]

    if nStart > 0:
        nStart -= 1

    if ((line[nStart] != '.' and not line[nStart].isdigit()) 
        or (line[numEndIndex] != '.' and not line[numEndIndex].isdigit())):
        # number is valid
        validNum = True

    numEndIndex +=1

    print("Start: " + str(nStart))
    print("End: " + str(numEndIndex))

    # check for adjacent character in surrounding lines
    if not validNum:
        if dIndex > 0:
            # check line above
            subset = data[dIndex - 1][nStart:numEndIndex]

            print("Data[" + str(dIndex -1) + "]: " + subset)
            for char in subset:
                if (not char.isdigit() and char != '.'):
                    validNum = True
                    break

        if dIndex < len(data) - 1:
            # check line below
            subset = data[dIndex + 1][nStart:numEndIndex]
            print("Data[" + str(dIndex + 1) + "]: " + subset)
            for char in subset:
                if (not char.isdigit() and char != '.'):
                    validNum = True
                    break

    if validNum:
        tempString = ""
        for c in num:
            tempString += c
        
        value += int(tempString)
        print(tempString + " is a valid num")

    return value 


def main():
    data = open(sys.argv[1]).read().strip().split('\n')

    total = 0

    for i, line in enumerate(data):
        print(line)
        tempNum = []

        for j, char in enumerate(line):
            if (char.isdigit()):
                tempNum.append(char)
            else:
                if (len(tempNum) > 0):
                    value = validateNum(data, i, j, tempNum)
                    total += value
                    tempNum = []
    
        if len(tempNum) > 0:
            value = validateNum(data, i, j, tempNum)
            total += value
            tempNum = []
    
    print()
    print("TOTAL: " + str(total))

if __name__ == "__main__":
    main()