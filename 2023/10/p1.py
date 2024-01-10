import sys
data = open(sys.argv[1]).read().strip().split('\n')

# pipes = ['|', '-', 'L', 'J', '7', 'F']
# S - start of pipe maze
# . - empty tile

dl = len(data) - 1
ll = len(data[0]) - 1

def getStart():
    for l, line in enumerate(data):
        for c, char in enumerate(line):
            if char == "S":
                return [l, c]
    return [-1, -1] 

def isValid(sPipe, nPipe):
    ydiff = sPipe[0] - nPipe[0] 
    xdiff = sPipe[1] - nPipe[1] 
    c = data[nPipe[0]][nPipe[1]]
    s = data[sPipe[0]][sPipe[1]]

    if c == '.':
        return False

    if c == "S":
        return True

    if ydiff > 0:
        if s != "|" and s != "L" and s != "J" and s != "S":
            # pipe can't go up
            return False
        if c == "7" or c == "|" or c == "F":
            return True
    if ydiff < 0:
        if s != "|" and s != "7" and s != "F" and s != "S":
            # pipe can't go down
            return False
        if c == "L" or c == "|" or c == "J":
            return True

    if xdiff > 0:
        if s != "-" and s != "7" and s != "J" and s != "S":
            # pipe cant go left
            return False
        if c == "L" or c == "-" or c == "F":
            return True
    if xdiff < 0:
        if s != "-" and s != "F" and s != "L" and s != "S":
            # pipe cant go right
            return False
        if c == "J" or c == "-" or c == "7":
            return True
    return False


def move(curr, prev = None):
    next = []
    if curr[0] > 0 and isValid(curr, [curr[0] - 1, curr[1]]):
        n = [curr[0] - 1, curr[1]]
        if n != prev:
            next.append(n)

    if curr[0] < dl and isValid(curr, [curr[0] + 1, curr[1]]):
        n = [curr[0] + 1, curr[1]]
        if n != prev:
            next.append(n)

    if curr[1] > 0 and isValid(curr, [curr[0], curr[1] - 1]):
        n = [curr[0], curr[1] - 1]
        if n != prev:
            next.append(n)

    if curr[1] < ll and isValid(curr, [curr[0], curr[1] + 1]):
        n = [curr[0], curr[1] + 1]
        if n != prev:
            next.append(n)
    return next

start = getStart()

# from S move iterate over the loop in one direction
s = move(start) 
c = data[s[0][0]][s[0][1]]
cpos = s[0] # current position
ppos = start # previous position
pcount = 0
while(c != "S"):
    m = move(cpos, ppos)
    ppos = cpos
    cpos = m[0]
    pcount += 1 

    c = data[m[0][0]][m[0][1]]

fpos = int(pcount / 2) + pcount % 2
print(fpos)
