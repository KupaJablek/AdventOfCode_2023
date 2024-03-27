import itertools

buckets = [ 50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40 ]

def day17():
    tot = 0

    combo = {}

    for i in range(len(buckets)):
        stotal = 0

        for combination in itertools.combinations(buckets, i):
            if sum(combination) == 150:
                stotal += 1
                l = len(combination)
                if l not in combo:
                    combo[l] = 0
                combo[l] += 1
        tot += stotal

    return tot, combo

p1, combos = day17()
print("Part 1: " + str(p1))

smallest = -1
for key in combos:
    if smallest == -1:
        smallest = key

    if key < smallest:
        smallest = key

print("Part 2: " + str(combos[smallest]))
