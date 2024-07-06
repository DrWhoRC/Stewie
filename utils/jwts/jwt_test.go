package jwts_test

import (
	"fim/utils/jwts"
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	payload := jwts.JwtPayload{
		UserID:   1,
		Username: "Richard",
		Role:     1,
	}
	token, _ := jwts.GenToken(payload, "wpywatsendw0517", 8)
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	claims, _ := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImJhdGlzdGEiLCJyb2xlIjoxLCJleHAiOjE3MjAyNzMxNzN9.Rc_NufsJ4YDH15LhPL-8E7H9bklreIjLfRJTfRfvOQU", "12345")
	fmt.Println(claims)
}
