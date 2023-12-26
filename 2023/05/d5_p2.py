import sys

file = open(sys.argv[1]).read().strip().split("\n")
seeds = [int(s) for s in file[0].split(':')[1].strip().split(" ")]
sp = []  # seed pairs
for i in range(0, len(seeds), 2):
    sp.append([seeds[i], seeds[i + 1]])


def updateSeed(locations, sPair):
    # parse seed range
    return sPair


ranges = []
d = file[2:] 

cMap = []
for line in d:
    if line == "" or line == None:
        continue
    if "map" in line:
        if cMap:
            ranges.append(cMap)
            cMap = []
    else:
        bits = [int(x) for x in line.strip().split(" ")]
        # dest_rng, src_start, src_rng
        cMap.append(bits) 
ranges.append(cMap)

#print(ranges)
for destination in ranges:
    tSeeds = []
    for s in sp:
        us = updateSeed(destination, s)
        tSeeds.append(us)
    sp = tSeeds

sp = sorted(sp, key = lambda e: e[0])
print(sp[0])
