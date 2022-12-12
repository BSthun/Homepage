package compare

import "command/types/common"

func NewPhotoCompare() common.Subcommand {
	sc := &Subcommand{
		Dir1: "",
		Dir2: "",
	}

	return sc
}
