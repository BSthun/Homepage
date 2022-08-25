package payload

type PhotoPerson struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}
