import sys
import math
data = open(sys.argv[1]).read().split("\n")
instructions = data[0].strip()

locations = {}
ghostStarts = []
for line in data[2:]:
    if line != '':
        c = line.strip().split('=')
        code = c[0].strip()
        dirs = c[1].replace(" ", '').strip("()").split(',')
        locations[code] = [dirs[0], dirs[1]]
        if code[2] == 'A':
            ghostStarts.append(code)

allSteps = []
for g, gs in enumerate(ghostStarts):
    steps = 0
    i = 0
    while gs[2] != 'Z':
        if instructions[i] == 'L':
            gs = locations[gs][0]
        else:
            gs = locations[gs][1]
        i += 1
        if i == len(instructions):
            i = 0
        steps += 1
    allSteps.append(steps)
print("total steps taken: " + str(math.lcm(*allSteps)))
