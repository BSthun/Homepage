package enum

type PhotoVisibility string

const (
	PhotoVisibilityPublic  PhotoVisibility = "public"
	PhotoVisibilityPrivate PhotoVisibility = "private"
	PhotoVisibilityHidden  PhotoVisibility = "hidden"
)
