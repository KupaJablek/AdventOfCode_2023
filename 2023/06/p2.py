import sys

file = open(sys.argv[1]).read().strip()
data = file.split("\n")

times = data[0].split(":")
times = list(filter(None, times[1].split(" ")))
t = int(''.join(times))

dist = data[1].split(":")
dist = list(filter(None, dist[1].split(" ")))
d = int(''.join(dist))


def calcValue(msh, increment):
    val = 0
    while (msh < t):
        test = (t - msh) * msh
        if test > d:
            val = msh
            break
        msh += increment
    return val


first = calcValue(1, 1)
last = calcValue(t - 1, -1)

print(last - first + 1)
