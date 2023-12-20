import sys

file = open(sys.argv[1]).read().strip()
data = file.split("\n")
hands = []


def getValue(s):
    match s:
        case "A":
            return 14
        case "K":
            return 13
        case "Q":
            return 12
        case "T":
            return 10
        case "J":
            return 1
        case _:
            return int(s)


for line in data:
    cards, bet = line.split(" ")
    hand = {}

    for card in cards:
        if card not in hand:
            hand[card] = 1
        else:
            hand[card] += 1

    hand = {k: v for k, v in sorted(hand.items(), key=lambda item: item[1], reverse=True)}
    jokers = 0
    if "J" in hand:
        jokers = hand["J"]
        del hand["J"]

    keys = list(hand.keys())

    # in case whole hand is jokers
    if len(keys) == 0:
        hand['A'] = 0
        keys.append('A')

    hand[keys[0]] += jokers

    rank = 0
    if len(hand) == 1:
        rank = 7
    elif hand[keys[0]] == 4:
        rank = 6
    elif hand[keys[0]] == 3 and hand[keys[1]] == 2:
        rank = 5
    elif hand[keys[0]] == 3:
        rank = 4
    elif hand[keys[0]] == 2 and hand[keys[1]] == 2:
        rank = 3
    elif hand[keys[0]] == 2:
        rank = 2
    else:
        rank = 1
    hands.append((rank, cards, bet))

# process the data
hands = sorted(hands, key=lambda e: (e[0], getValue(e[1][0]),
                                     getValue(e[1][1]),
                                     getValue(e[1][2]),
                                     getValue(e[1][3]),
                                     getValue(e[1][4]),
                                     ))
total = 0
for i, h in enumerate(hands):
    total += int(i + 1) * int(h[2])

print(total)
