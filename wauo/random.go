package wauo

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Randint python的random.randint
func Randint(min, max int) int {
	if min > max {
		panic("min > max")
	}
	iv := rand.Intn(max-min+1) + min
	return iv
}

// RandomSleep 随机睡眠
func RandomSleep(min, max int) {
	if min > max {
		panic("min > max")
	}
	delay := time.Duration(rand.Intn(max-min+1)+min) * time.Second
	time.Sleep(delay)

}
