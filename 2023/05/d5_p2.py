import sys
file = open(sys.argv[1]).read().strip().split("\n")
seeds = [int(s) for s in file[0].split(':')[1].strip().split(" ")]
sp = []  # seed pairs
for i in range(0, len(seeds), 2):
    sp.append([seeds[i], (seeds[i] + seeds[i + 1] - 1)])


def updateSeed(locations, s):
    updated = []
    extras = []
    seedStart, seedEnd = s[0], s[1]

    for l in locations:
        dst, srcS, srcE = l[0], l[1], l[1] + l[2] - 1

        if srcS <= seedStart and srcE >= seedEnd:
            updated = [seedStart - srcS + dst, seedEnd - srcS + dst]
            break
        elif srcS > seedStart and srcS <= seedEnd and srcE > seedEnd:
            updated = [dst, seedEnd - srcS + dst]
            extras.append([seedStart, srcS - 1])
        elif srcE < seedEnd and srcE >= seedStart and srcS < seedStart:
            updated = [seedStart - srcS + dst, srcE - seedStart + seedStart - srcS + dst]
            extras.append([srcE + 1, seedEnd])
        else:
            updated = [dst, srcE - seedStart + seedStart - srcS + dst]
            extras.append([seedStart, srcS - 1])
            extras.append([srcE + 1, seedEnd])
    if updated == []:
        return s, []
    return updated, extras


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

for s, seed in enumerate(sp):
    for d, loc in enumerate(ranges):
        ret, ex = updateSeed(loc, sp[s])

        sp[s] = ret
        if ex != []:
            for e in ex:
                sp.append(e)

sp = sorted(sp, key = lambda e: e[0])
print(sp)
