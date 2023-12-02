# Aoc day 2 p1 solution
import sys

# determine games only containing 12 red, 13 green, and 14 blue cubes 
# game #id: # colour

colourCounts = {
    "red": 12,
    "green": 13,
    "blue": 14
}

def parseFile(fname):
    lines = []
    try:
        lines = open(fname).read().strip().split('\n')
    except:
        return -1
    
    total = 0

    for line in lines:
        # get game ID
        data = line.split(":")
        # game id acquired
        gameId = int(data[0].strip("Game").strip())
        
        # check the data
        # each subset is seperated by a ';'
        subsets = data[1].split(';')

        validGame = True

        for i, draw in enumerate(subsets):
            colourData = draw.strip().split(',')

            for c in colourData:
                values = c.strip(' ').split(' ')
                # c[0] -> count
                # c[1] -> colour

                if int(values[0]) > colourCounts[values[1]]:
                    validGame = False
                    break

            if not validGame:
                break
        
        if validGame:
            total += gameId

    return total

def main():
    result = parseFile(sys.argv[1])
    print("the total is: " + str(result))

if __name__ == "__main__":
    main()