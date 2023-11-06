package ogp

type Locale struct {
	Locale     string
	Alternates []string
}

type Image struct {
	URL       string
	SecureURL string
	Type      string
	Width     uint64
	Height    uint64
	Alt       string
}

type Video struct {
	URL       string
	SecureURL string
	Type      string
	Width     uint64
	Height    uint64
	Alt       string
}

type Audio struct {
	URL       string
	SecureURL string
	Type      string
}
