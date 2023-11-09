package ogp

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Parse parses the given io.Reader and returns the Open Graph protocol object.
func Parse(r io.Reader) (*Object, error) {
	tokenizer := html.NewTokenizer(r)

	var raw []Metadata
outer:
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if errors.Is(err, io.EOF) {
				break outer
			}
			return nil, err
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.DataAtom != atom.Meta {
				continue outer
			}

			var property, content string
			for _, attr := range token.Attr {
				if attr.Key == "property" {
					property = attr.Val
				}
				if attr.Key == "content" {
					content = attr.Val
				}
			}
			if !strings.HasPrefix(property, "og:") {
				continue outer
			}
			raw = append(raw, Metadata{Property: property, Content: content})
		}
	}

	obj := &Object{}
	for _, d := range raw {
		switch d.Property {
		case "og:title":
			if obj.Title == "" {
				obj.Title = d.Content
			}
		case "og:type":
			if obj.Type == "" {
				obj.Type = d.Content
			}
		case "og:image", "og:image:url": // image
			obj.Images = append(obj.Images, &Image{URL: d.Content})
		case "og:image:secure_url":
			if length := len(obj.Images); length > 0 && obj.Images[length-1].SecureURL == "" {
				obj.Images[length-1].SecureURL = d.Content
			}
		case "og:image:type":
			if length := len(obj.Images); length > 0 && obj.Images[length-1].Type == "" {
				obj.Images[length-1].Type = d.Content
			}
		case "og:image:width":
			if length := len(obj.Images); length > 0 && obj.Images[length-1].Width == 0 {
				v, _ := strconv.ParseUint(d.Content, 10, 64) // ignore error
				obj.Images[length-1].Width = v
			}
		case "og:image:height":
			if length := len(obj.Images); length > 0 && obj.Images[length-1].Height == 0 {
				v, _ := strconv.ParseUint(d.Content, 10, 64) // ignore error
				obj.Images[length-1].Height = v
			}
		case "og:image:alt":
			if length := len(obj.Images); length > 0 && obj.Images[length-1].Alt == "" {
				obj.Images[length-1].Alt = d.Content
			}
		case "og:url":
			if obj.URL == "" {
				obj.URL = d.Content
			}
		case "og:audio": // audio
			obj.Audios = append(obj.Audios, &Audio{URL: d.Content})
		case "og:audio:secure_url":
			if length := len(obj.Audios); length > 0 && obj.Audios[length-1].SecureURL == "" {
				obj.Audios[length-1].SecureURL = d.Content
			}
		case "og:audio:type":
			if length := len(obj.Audios); length > 0 && obj.Audios[length-1].Type == "" {
				obj.Audios[length-1].Type = d.Content
			}
		case "og:description":
			if obj.Description == "" {
				obj.Description = d.Content
			}
		case "og:determiner":
			if obj.Determiner == "" {
				obj.Determiner = d.Content
			}
		case "og:locale":
			if obj.Locale == nil {
				obj.Locale = &Locale{}
			}
			if obj.Locale.Locale == "" {
				obj.Locale.Locale = d.Content
			}
		case "og:locale:alternate":
			if obj.Locale == nil {
				obj.Locale = &Locale{}
			}
			obj.Locale.Alternates = append(obj.Locale.Alternates, d.Content)
		case "og:site_name":
			if obj.SiteName == "" {
				obj.SiteName = d.Content
			}
		case "og:video": // video
			obj.Videos = append(obj.Videos, &Video{URL: d.Content})
		case "og:video:secure_url":
			if length := len(obj.Videos); length > 0 && obj.Videos[length-1].SecureURL == "" {
				obj.Videos[length-1].SecureURL = d.Content
			}
		case "og:video:type":
			if length := len(obj.Videos); length > 0 && obj.Videos[length-1].Type == "" {
				obj.Videos[length-1].Type = d.Content
			}
		case "og:video:width":
			if length := len(obj.Videos); length > 0 && obj.Videos[length-1].Width == 0 {
				v, _ := strconv.ParseUint(d.Content, 10, 64) // ignore error
				obj.Videos[length-1].Width = v
			}
		case "og:video:height":
			if length := len(obj.Videos); length > 0 && obj.Videos[length-1].Height == 0 {
				v, _ := strconv.ParseUint(d.Content, 10, 64) // ignore error
				obj.Videos[length-1].Height = v
			}
		case "og:video:alt":
			if length := len(obj.Videos); length > 0 && obj.Videos[length-1].Alt == "" {
				obj.Videos[length-1].Alt = d.Content
			}
		}
	}
	return obj, nil
}
