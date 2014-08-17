package helpers

type Assets struct {
	Normal         string
	NonStaticfile  string
	AlternateRoot  string
	BasicAuth      string
	DirectoryIndex string
}

func NewAssets() Assets {
	return Assets{
		Normal:         "../assets/normal_staticfile",
		NonStaticfile:  "../assets/non_staticfile_app",
		AlternateRoot:  "../assets/alternate_root",
		BasicAuth:      "../assets/basic_auth",
		DirectoryIndex: "../assets/directory_index",
	}
}
