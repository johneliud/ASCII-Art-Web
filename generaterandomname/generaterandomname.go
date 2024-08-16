package generaterandomname

import "math/rand"

func GenerateRandomName() string {
	characterSet := "abcdefghijklmnopqrstuvwxyz0123456789"

	name := make([]byte, 5)
	for i := range name {
		name[i] = characterSet[rand.Intn(len(characterSet))]
	}
	return string(name)
}
