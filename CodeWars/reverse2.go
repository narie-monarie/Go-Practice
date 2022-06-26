package kata

import "fmt"

func Solution(word string) string {
  var result string
  fmt.Scanf("%s", word)
    for _, v := range word {
        result = string(v) + result
    }
    return result
}
