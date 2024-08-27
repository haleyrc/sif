package pages

type MenuItem struct {
	Label string
	Path  string
}

type PageData struct {
	Name string
}

type SiteData struct {
	Author             string
	CopyrightStartYear int
	Description        string
	Menu               []MenuItem
	Title              string
}
