package photo

import (
	"command/types/cmm"
)

func NewPhoto() cmm.Subcommand {
	sc := &Subcommand{
		Dir: "",
	}

	return sc
}
