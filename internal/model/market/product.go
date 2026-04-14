package market

type Product struct {
	Title string
}

func (sub Product) String() string {
	return sub.Title
}
