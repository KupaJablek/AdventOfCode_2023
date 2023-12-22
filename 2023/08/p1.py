import sys

data = open(sys.argv[1]).read().split("\n")
instructions = data[0].strip()

dataIndex = 2
locations = {}
while dataIndex < len(data):
    if data[dataIndex] != '':
        c = data[dataIndex].split('=')
        code = c[0].strip()
        dirs = c[1].replace(" ", '').strip("()").split(',')

        locations[code] = [dirs[0], dirs[1]]
    dataIndex += 1

i = 0
steps = 0
curKey = 'AAA'
while curKey != 'ZZZ':
    if instructions[i] == 'L':
        curKey = locations[curKey][0]
    else:
        curKey = locations[curKey][1]

    i += 1
    if i == len(instructions):
        i = 0
    steps += 1
print("total steps taken: " + str(steps))
