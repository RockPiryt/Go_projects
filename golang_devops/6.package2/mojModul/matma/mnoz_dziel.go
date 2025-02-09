package matma

import "fmt"

// Multiply mnoży dwie liczby
func Multiply(a, b int) int {
    return a * b
}

// Divide dzieli dwie liczby (obsługuje dzielenie przez zero)
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("dzielenie przez zero")
    }
    return a / b, nil
}
