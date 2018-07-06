package main

import (
	"math/rand"
	"time"
)

var box []string

func init() {
	rand.Seed(time.Now().Unix())

	box = append(box, "🍇 Grapes")
	box = append(box, "🍈 Melon")
	box = append(box, "🍉 Watermelon")
	box = append(box, "🍊 Tangerine")
	box = append(box, "🍋 Lemon")
	box = append(box, "🍌 Banana")
	box = append(box, "🍍 Pineapple")
	box = append(box, "🍎 Red Apple")
	box = append(box, "🍏 Green Apple")
	box = append(box, "🍐 Pear")
	box = append(box, "🍑 Peach")
	box = append(box, "🍒 Cherries")
	box = append(box, "🍓 Strawberry")
	box = append(box, "🥝 Kiwi Fruit")
	box = append(box, "🍅 Tomato")
	box = append(box, "🥥 Coconut")
	box = append(box, "🥑 Avocado")
	box = append(box, "🍆 Eggplant")
	box = append(box, "🥔 Potato")
	box = append(box, "🥕 Carrot")
}

func pick() string {
	return box[rand.Intn(len(box))]
}
