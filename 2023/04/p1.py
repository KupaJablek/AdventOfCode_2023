import sys

data = open(sys.argv[1]).read().strip()
cards = data.split('\n')

# total number of cards
maxCards = len(cards)

# list to hold card counts
cardUses = [0 for _ in range(maxCards)]

# total number of cards earned
totalCards = 0

for i, card in enumerate(cards):
    cardUses[i] += 1

    # get the data
    cData = card.split(":")
    nums = cData[1].split("|")

    wNums = set(filter(None, nums[0].split(' ')))
    myNums = list(filter(None, nums[1].split(' ')))

    matches = 0

    for n in myNums:
        if n in wNums:
            matches += 1
    if matches > 0:
        uses = 0
        while uses < cardUses[i]:
            j = 0
            while j < matches and i + j < maxCards:
                cardUses[j + i + 1] += 1
                j += 1
            uses += 1

for tot in cardUses:
    totalCards += tot

print(str(totalCards))
