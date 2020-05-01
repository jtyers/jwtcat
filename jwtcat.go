package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		processJwt(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}

// JWT is a triplet of Base64-encoded strings separated by dots.
// We simply decode each in turn.
func processJwt(jwtStr string) string {
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(jwtStr, &jwt.MapClaims{})
	if err != nil {
		panic(fmt.Sprintf("error parsing token: %v", err))
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		panic(fmt.Sprintf("claims was %T, wanted *MapClaims: %v", token.Claims))
	}

	fmt.Printf("Signed using %q\n", token.Method.Alg())

	fmt.Printf("\n")
	fmt.Printf("Header\n")
	for k, v := range token.Header {
		fmt.Printf("HH  [%s] %v\n", k, v)
	}

	fmt.Printf("\n")
	fmt.Printf("Claims\n")
	for k, v := range *claims {
		fmt.Printf("CC  [%s] %v\n", k, v)
	}

	return ""
}
