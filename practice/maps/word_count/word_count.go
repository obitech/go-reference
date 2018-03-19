/*
* You have 50 bitcoins to distribute to 10 users:
* Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth.
* A vowel equals one coin and a user can’t get more than 10 coins.
* Print a map with each user’s name and the amount of coins distributed.
 */
package word_count

import (
    "strings"
)

func WordCount(words []string, coins int) (distribution map[string]int, excess int) {
    vowels := [5]string{"a", "e", "i", "o", "u"}
    distribution = make(map[string]int, len(words))

    // Iterating over every word
    for _, word := range words {
        excess = coins
        lower := strings.ToLower(word)
        count := 0
        // Iterating over every vowel
        for _, vowel := range vowels {
            // Checking vowel count, adding to count
            if c := strings.Count(lower, vowel); c > 0 {
                count += c
            }
        }

        // People can't get more than 10 coins
        if count > 10 {
            count = 10
        }

        payout := count
        excess = coins - payout

        if excess < 0 && coins > 0 {
            payout = -payout
        } else if excess < 0 && coins == 0 {
            payout = 0
        }

        if excess < 0 {
            excess = 0
        }

        distribution[word] = payout
        coins = excess
    }
    return
}
