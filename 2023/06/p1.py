import sys

file = open(sys.argv[1]).read().strip()
data = file.split("\n")

times = data[0].split(":")
times = times[1].split(" ")
times = list(filter(None, times))

dist = data[1].split(":")
dist = dist[1].split(" ")
dist = list(filter(None, dist))

index = 0
total = 0

while index < len(times):
    # calculate shortest time to hole the button
    d = int(dist[index])
    t = int(times[index])

    possibleWaysToBeat = 0
    msh = 1  # miliseconds held
    while (msh < t):
        test = (t - msh) * msh
        if test > d:
            possibleWaysToBeat += 1
        msh += 1

    if total == 0:
        total = possibleWaysToBeat
    else:
        total = total * possibleWaysToBeat
    index += 1

print(total)
