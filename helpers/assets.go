package helpers

type Assets struct {
	Normal string
}

func NewAssets() Assets {
	return Assets{
		Normal: "../assets/normal_staticfile",
	}
}
