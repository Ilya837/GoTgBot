package domain

type Subdomain struct {
	Title string
}

func (sub Subdomain) String() string {
	return sub.Title
}
