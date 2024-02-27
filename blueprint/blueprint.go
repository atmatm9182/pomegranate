package blueprint

type FileSpec struct {
	Type    string
	Raw     string
	Src     string
	Entries map[string]FileSpec
}

type Blueprint struct {
	absolutePath string
	Project struct {
		Name  string
		Files map[string]FileSpec
	}
}
