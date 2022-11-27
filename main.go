package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(30)
	fmt.Println(arr)
	fmt.Println(Concsort(arr))
}
