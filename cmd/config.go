package cmd

type Config struct {
	HistoryFile  string
	BookMarkFile string
	Command      string
	NoOutput     bool
	UseSSH       bool
	Make         bool
	CustomFile   []CustomFile
	FuzzyFinder  FuzzyFinder
}

type FuzzyFinder struct {
	CommandPath string
	Option      string
}

type CustomFile struct {
	Name   string
	Path   string
	Format string
}

type Record struct {
	Number int
	Path   string
}
