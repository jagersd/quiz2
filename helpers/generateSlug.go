package helpers

import (
	"math/rand"
	"strings"
	"time"
)

const charSet string = "abcdefghjkmnopqrstuvwxyz" + "123456789"

func GenerateSlug() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder
	length := 6
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}
