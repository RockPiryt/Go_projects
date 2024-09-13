package prose

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"05", "12", "2018"}, "/"))
}

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ",")
	result += " i "
	result += phrases[len(phrases)-1]
	return result
}
