from p1 import *
import sys
data = open(sys.argv[1]).read().strip().split('\n')
rlim = len(data[0])
blim = len(data)

vpipes = []

start = getStart()
vpipes.append(start)
# from S move iterate over the loop in one direction
s = move(start) 
c = data[s[0][0]][s[0][1]]

cpos = s[0] # current position
vpipes.append(cpos)
ppos = start # previous position
pcount = 0

while(c != "S"):
    m = move(cpos, ppos)
    ppos = cpos
    cpos = m[0]
    vpipes.append(m[0])
    pcount += 1 

    c = data[m[0][0]][m[0][1]]
fpos = int(pcount / 2) + pcount % 2

for d in data:
    print(d)
print()

totalEnc = 0
for l, line in enumerate(data):
    for p, pipe in enumerate(line):
        if [l,p] in vpipes:
            continue
        # ray cast to the right
        # FJ && L7 => one line cross L-----7 -> one line cross
        queue = [] 
        rpos = p

        zigs = 0
        while rpos < rlim:
            if [l,rpos] not in vpipes:
                rpos += 1
                continue
            c = data[l][rpos]
            if c == '-':
                rpos += 1
                continue
            if (c == 'J' and queue[-1] == 'F') or (c == '7' and queue[-1] == 'L'):
                zigs += 1
                queue.pop()
            else:
                queue.append(c)
            rpos += 1
        tot = len(queue) + zigs 
        if tot % 2 != 0:
            totalEnc += 1
print("Total enclosed: " + str(totalEnc))
