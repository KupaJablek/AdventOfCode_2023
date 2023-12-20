package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
    bid int
    hand string
    rank int
}

func getValue(s string) int {
    switch s {
        case "A":
            return 14
        case "K":
            return 13
        case "Q":
            return 12
        case "J":
            return 11
        case "T":
            return 10
    }
    val, _:= strconv.Atoi(s)
    return val
}

func toString(h Hand) {
    fmt.Printf("rank: %d, hand: %s, bid: %d", h.rank, h.hand, h.bid)
}

func lhsLessThan(lhs Hand, rhs Hand) bool {
    if lhs.rank < rhs.rank {
        return true
    }
    if lhs.rank > rhs.rank {
        return false
    }
    // settle tie by card value
    // A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
    for i := 0; i < len(lhs.hand); i++ {
        lv := getValue(string(lhs.hand[i]))
        rv := getValue(string(rhs.hand[i]))

        if lv != rv {
            return lv < rv
        }
    }
    return false
}

func sortHands(h []Hand) []Hand {
    var i, j int
    var key Hand
    for i = 1; i < len(h); i++ {
        key = h[i]
        j = i - 1
        for j >= 0 && lhsLessThan(h[j], key){
            h[j + 1] = h[j]
            j--
        }
        h[j + 1] = key;
    }
    return h
}

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Printf("No file arg\n")
        return
    }

    file, err := os.Open(args[1])
    if err != nil {
        fmt.Println("file no work" + err.Error())
        return
    }

    hands := []Hand {} 
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        values := strings.Split(scanner.Text(), " ")
    
        var h Hand
        h.hand = values[0]
        b, _ := strconv.Atoi(values[1])
        h.bid = b
        
        // determine rank
        rank := -1
        cards := map[string]int {}

        for i := 0; i < len(h.hand); i++ {
            cCard := string(h.hand[i])
            if _, found := cards[cCard]; !found {
                cards[cCard] = 1
            } else {
                cards[cCard] += 1
            }
        }

        switch len(cards) {
            case 5:
                // high card
                rank = 1
            case 4:
                // pair
                rank = 2
            case 1:
                // all of a kind
                rank = 7
        }

        if rank == -1 {
            pair := 0
            trio := 0
            quad := 0
            for _, e := range cards {
                switch e {
                    case 2:
                        pair++
                    case 3:
                        trio++
                    case 4:
                        quad++
                }
            }
            if pair == 2 {
                rank = 3
            } else if pair == 1 && trio == 1{
                //full house
                rank = 5
            } else if trio > 0{
                rank = 4
            } else if quad > 0 {
                rank = 6
            }
        }

        h.rank = rank
        hands = append(hands, h)
    }

    // sort the list
    hands = sortHands(hands)
    total := 0
    for i := len(hands) -1; i > -1 ; i-- {
        total += hands[i].bid * (len(hands) - i)
    }

    fmt.Println(total)
}
