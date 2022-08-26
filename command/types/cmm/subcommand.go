package cmm

type Subcommand interface {
	Parse([]string) error
	Run() error
	Name() string
}
