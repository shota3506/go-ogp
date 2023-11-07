package ogp

import (
	"strconv"

	"golang.org/x/net/html"
)

type Locale struct {
	Locale     string
	Alternates []string
}

// An Image is an image which represents object within the graph.
type Image struct {
	URL       string
	SecureURL string
	Type      string
	Width     uint64
	Height    uint64
	Alt       string
}

func (i *Image) html() []*html.Node {
	var nodes []*html.Node
	if i.URL != "" {
		nodes = append(nodes, metadataNode("og:image", i.URL))
	}
	if i.SecureURL != "" {
		nodes = append(nodes, metadataNode("og:image:secure_url", i.SecureURL))
	}
	if i.Type != "" {
		nodes = append(nodes, metadataNode("og:image:type", i.Type))
	}
	if i.Width != 0 {
		nodes = append(nodes, metadataNode("og:image:width", strconv.FormatUint(i.Width, 10)))
	}
	if i.Height != 0 {
		nodes = append(nodes, metadataNode("og:image:height", strconv.FormatUint(i.Height, 10)))
	}
	if i.Alt != "" {
		nodes = append(nodes, metadataNode("og:image:alt", i.Alt))
	}
	return nodes
}

// A Video is a video that complements object.
type Video struct {
	URL       string
	SecureURL string
	Type      string
	Width     uint64
	Height    uint64
	Alt       string
}

func (v *Video) html() []*html.Node {
	var nodes []*html.Node
	if v.URL != "" {
		nodes = append(nodes, metadataNode("og:video", v.URL))
	}
	if v.SecureURL != "" {
		nodes = append(nodes, metadataNode("og:video:secure_url", v.SecureURL))
	}
	if v.Type != "" {
		nodes = append(nodes, metadataNode("og:video:type", v.Type))
	}
	if v.Width != 0 {
		nodes = append(nodes, metadataNode("og:video:width", strconv.FormatUint(v.Width, 10)))
	}
	if v.Height != 0 {
		nodes = append(nodes, metadataNode("og:video:height", strconv.FormatUint(v.Height, 10)))
	}
	if v.Alt != "" {
		nodes = append(nodes, metadataNode("og:video:alt", v.Alt))
	}
	return nodes
}

// An Audio is an audio file to accompany object.
type Audio struct {
	URL       string
	SecureURL string
	Type      string
}

func (a *Audio) html() []*html.Node {
	var nodes []*html.Node
	if a.URL != "" {
		nodes = append(nodes, metadataNode("og:audio", a.URL))
	}
	if a.SecureURL != "" {
		nodes = append(nodes, metadataNode("og:audio:secure_url", a.SecureURL))
	}
	if a.Type != "" {
		nodes = append(nodes, metadataNode("og:audio:type", a.Type))
	}
	return nodes
}
