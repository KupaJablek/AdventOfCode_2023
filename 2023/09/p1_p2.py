import sys
data = open(sys.argv[1]).read().strip().split('\n')


def f(values):
    temp = []
    for i in range(len(values) - 1):
        temp.append(values[i+1] - values[i])

    if all(v == 0 for v in values):
        return values[-1]
    else:
        # recursive call of f() with remaining bits
        return values[-1] + f(temp)


total = 0
total2 = 0
for line in data:
    vals = [int(v) for v in line.split(' ')]
    total += f(vals)

    vals2 = [int(v) for v in reversed(line.split(' '))]
    total2 += f(vals2)

print("Total #1: " + str(total))
print("Total #2: " + str(total2))
