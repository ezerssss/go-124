package main

import "fmt"

func countLongestSubsequence(s string) int {
	ans := 0
	dp := make([]int, len(s)) // Creation of a dynamic size "array" (also known as a slice)

	var currChar byte
	var possiblePairIndex int

	// Start with 1 as that is the minimum length for a valid parenthesis pair
	for i := 1; i < len(s); i++ {
		currChar = s[i]

		if currChar == ')' {
			// Its opening parenthesis pair is directly beside it
			if s[i-1] == '(' {
				dp[i] = 2 // Count the parenthesis pair and store in dp table

				// Count the possible previous subsequence to the left of the parenthesis pair (i - 2)
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}

			} else {
				// Two meanings if this happen:
				// 1. It could possibly be a parent of a nested parenthesis, example: (()) <-- Beside a ) is also another ) so we have to find the opening parenthesis pair

				// 2. It is an invalid sequence, example: ()) <-- Although the first pair is valid, the second ) does no longer have an opening parentheses pair

				// With that in mind, we must first know where is the location of the possible opening parenthesis pair, and we can find that by leveraging the fact that we store the sequence of the "children" in the dp table!
				// Meaning we can find the location by: i - 1 - sequenceOfChildren, since sequenceOfChildren represents the maximum length of a valid child

				possiblePairIndex = i - 1 - dp[i-1] // dp[i-1] represents the ending ) of the "children" meaning it stores the longest sequence of children

				// The ) doesn't have a pair
				if possiblePairIndex < 0 || s[possiblePairIndex] != '(' {
					continue
				}

				// It does have a pair! Now we just need to add everything up:
				// Sequence of the children + Possible sequence just before the opening pair parentheses
				dp[i] = dp[i-1] + 2

				if possiblePairIndex-1 >= 0 {
					dp[i] += dp[possiblePairIndex-1]
				}
			}

			// Store answer
			ans = max(ans, dp[i])
		}
	}

	return ans
}

func main() {
	longestSubsequence := countLongestSubsequence("(())(()())()))")
	fmt.Println("Longest subsequence:")
	fmt.Println(longestSubsequence)
}
