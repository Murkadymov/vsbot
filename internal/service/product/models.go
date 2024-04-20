package product

var AllProducts []*Product = []*Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
	{Title: "six"},
	{Title: "seven"},
}

type Product struct {
	Title string
}
