package product

var numberEditElement int
var EditElement bool
var Bollgame bool

var allProduct = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Product struct {
	Title string
}
