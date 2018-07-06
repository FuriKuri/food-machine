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
	box = append(box, "🌽 Ear of Corn")
	box = append(box, "🌶 Hot Pepper")
	box = append(box, "🥒 Cucumber")
	box = append(box, "🥦 Broccoli")
	box = append(box, "🍄 Mushroom")
	box = append(box, "🥜 Peanuts")
	box = append(box, "🌰 Chestnut")
	box = append(box, "🍞 Bread")
	box = append(box, "🥐 Croissant")
	box = append(box, "🥖 Baguette Bread")
	box = append(box, "🥨 Pretzel")
	box = append(box, "🥞 Pancakes")
	box = append(box, "🧀 Cheese Wedge")
	box = append(box, "🍖 Meat on Bone")
	box = append(box, "🍗 Poultry Leg")
	box = append(box, "🥩 Cut of Meat")
	box = append(box, "🥓 Bacon")
	box = append(box, "🍔 Hamburger")
	box = append(box, "🍟 French Fries")
	box = append(box, "🍕 Pizza")
	box = append(box, "🌭 Hot Dog")
	box = append(box, "🥪 Sandwich")
	box = append(box, "🌮 Taco")
	box = append(box, "🌯 Burrito")
}

func pick() string {
	return box[rand.Intn(len(box))]
}
