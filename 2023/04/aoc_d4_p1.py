import sys

data = open(sys.argv[1]).read().strip()
cards = data.split('\n')

total = 0

for card in cards:
    # print(card)
    # get the data
    cData = card.split(":")
    nums = cData[1].split("|")

    wNums = set(filter(None, nums[0].split(' ')))
    myNums = list(filter(None, nums[1].split(' ')))

    multiplier = 0
    for n in myNums:
        if n in wNums:
            if multiplier == 0:
                multiplier = 1
            else:
                multiplier *= 2

    total += (1 * multiplier)

print(str(total))
