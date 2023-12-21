import sys

file = open(sys.argv[1]).read().strip()
data = file.split("\n")

seedData = data[0].split(':')
seeds = seedData[1].strip().split(" ")
for i, seed in enumerate(seeds):
    seeds[i] = int(seed)
currMap = []


def updateSeeds():
    for s, seed in enumerate(seeds):
        for i, item in enumerate(currMap):
            if seed >= item[0] and seed <= (item[0] + item[1]):
                seeds[s] = seed + item[2]
                break
    print(seeds)


for i, line in enumerate(data):
    if "map" in line:
        if currMap != []:
            updateSeeds()
            currMap = []
    elif "seeds" not in line and line != "":
        d = line.split(" ")
        # range
        r = int(d[2])
        # source
        source = int(d[1])
        step = int(d[0]) - int(d[1])
        currMap.append((source, r, step))

if currMap != []:
    updateSeeds()

seeds.sort()
print("Final order: " + str(seeds))
print("\nLowest location #: " + str(seeds[0]))
