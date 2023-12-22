import sys

file = open(sys.argv[1]).read().strip()
data = file.split("\n")

seedData = data[0].split(':')
seeds = seedData[1].strip().split(" ")

seedRanges = []
cpair = []
for i, seed in enumerate(seeds):
    if len(cpair) < 1:
        cpair.append(int(seed))
    elif len(cpair) == 1:
        cpair.append(int(seed))
        seedRanges.append(cpair)
        cpair = []

    seeds[i] = int(seed)
currMap = []


def updateRangedSeeds():
    tempSeeds = []
    print(seedRanges)
    for p, pair in enumerate(seedRanges):
        seedStart = pair[0]
        seedRange = pair[1]
        for i, item in enumerate(currMap):
            sourceStart, sourceRange = item[0], item[1]
            sourceEnd = sourceStart + sourceRange - 1
            step = item[2]

            if sourceStart > seedStart and sourceStart > seedStart + seedRange:
                # no overlap, seed range is prior to source start
                tempSeeds.append(pair)
                continue
            if sourceEnd < seedStart and sourceEnd < seedStart:
                # no overlap, seed range is post source start
                tempSeeds.append(pair)
                continue

            # split portions of seed range that do not overlap
            startOffset = sourceStart - seedStart
            endOffset = (sourceEnd) - (seedStart + seedRange - 1)

            if startOffset <= 0 and endOffset >= 0:
                # everything is in range
                tempSeeds.append([seedStart + step, seedRange])
                continue

            newSRPair = pair
            # check for start offset
            if startOffset > 0:
                sr = (seedStart, startOffset)
                tempSeeds.append(sr)
                newSRPair = (seedStart + startOffset, seedRange - startOffset)
            if endOffset < 0:
                newSRPair = [newSRPair[0], newSRPair[1] + endOffset]
                sr = [newSRPair[0] + newSRPair[1], (-1 * endOffset)]
                tempSeeds.append(sr)

            newSRPair = [newSRPair[0] + step, newSRPair[1]]
            tempSeeds.append(newSRPair)
    return tempSeeds


def updateSeeds():
    for s, seed in enumerate(seeds):
        for i, item in enumerate(currMap):
            if seed >= item[0] and seed <= (item[0] + item[1]):
                seeds[s] = seed + item[2]
                break


for i, line in enumerate(data):
    if "map" in line:
        if currMap != []:
            updateSeeds()
            seedRanges = updateRangedSeeds()
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
print(seedRanges)
seedRanges = sorted(seedRanges, key=lambda e: e[0])
print("\nLowest location #: " + str(seeds[0]))
print("\nLowest ranged location #: " + str(seedRanges[0]))
