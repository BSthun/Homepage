package photo

import (
	"command/types/common"
)

func NewPhoto() common.Subcommand {
	sc := &Subcommand{
		Dir: "",
	}

	return sc
}
