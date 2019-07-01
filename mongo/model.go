package mongo

// Link friend link model
type Link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Img  string `json:"img"`
}

// Book model
type Book struct {
	Img      string `json:"img"`
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Href     string `json:"href"`
	Code     string `json:"code"`
}

// Project model
type Project struct {
	Img      string `json:"img"`
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Href     string `json:"href"`
}
