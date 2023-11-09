package ogp_test

import (
	"bytes"
	"slices"
	"testing"

	. "github.com/shota3506/go-ogp"
	"golang.org/x/net/html"
)

func TestObject_HTML(t *testing.T) {
	for tn, tc := range map[string]struct {
		obj      *Object
		expected []string
	}{
		"basic": {
			obj: &Object{
				Title:  "The Rock",
				Type:   "video.movie",
				Images: []*Image{{URL: "https://example.com/rock.jpg"}},
				URL:    "https://www.imdb.com/title/tt0117500/",
			},
			expected: []string{
				`<meta property="og:title" content="The Rock"/>`,
				`<meta property="og:type" content="video.movie"/>`,
				`<meta property="og:image" content="https://example.com/rock.jpg"/>`,
				`<meta property="og:url" content="https://www.imdb.com/title/tt0117500/"/>`,
			},
		},
		"detailed": {
			obj: &Object{
				Title: "The Rock",
				Type:  "video.movie",
				Images: []*Image{
					{
						URL:       "https://example.com/rock.jpg",
						SecureURL: "https://secure.example.com/rock.jpg",
						Type:      "image/jpeg",
						Width:     400,
						Height:    300,
						Alt:       "A shiny red apple with a bite taken out",
					},
					{
						URL:    "https://example.com/rock2.jpg",
						Height: 1000,
					},
					{
						URL: "https://example.com/rock3.jpg",
					},
				},
				URL: "https://www.imdb.com/title/tt0117500/",
				Audios: []*Audio{
					{
						URL:       "https://example.com/sound.mp3",
						SecureURL: "https://secure.example.com/sound.mp3",
						Type:      "audio/mpeg",
					},
				},
				Description: "Sean Connery found fame and fortune as the suave, sophisticated British agent, James Bond.",
				Determiner:  "the",
				Locale: &Locale{
					Locale:     "en_GB",
					Alternates: []string{"fr_FR", "es_ES"},
				},
				SiteName: "IMDb",
				Videos: []*Video{
					{
						URL:       "https://example.com/movie.swf",
						SecureURL: "https://secure.example.com/movie.swf",
						Type:      "application/x-shockwave-flash",
						Width:     400,
						Height:    300,
					},
				},
			},
			expected: []string{
				`<meta property="og:title" content="The Rock"/>`,
				`<meta property="og:type" content="video.movie"/>`,
				`<meta property="og:image" content="https://example.com/rock.jpg"/>`,
				`<meta property="og:image:secure_url" content="https://secure.example.com/rock.jpg"/>`,
				`<meta property="og:image:type" content="image/jpeg"/>`,
				`<meta property="og:image:width" content="400"/>`,
				`<meta property="og:image:height" content="300"/>`,
				`<meta property="og:image:alt" content="A shiny red apple with a bite taken out"/>`,
				`<meta property="og:image" content="https://example.com/rock2.jpg"/>`,
				`<meta property="og:image:height" content="1000"/>`,
				`<meta property="og:image" content="https://example.com/rock3.jpg"/>`,
				`<meta property="og:url" content="https://www.imdb.com/title/tt0117500/"/>`,
				`<meta property="og:audio" content="https://example.com/sound.mp3"/>`,
				`<meta property="og:audio:secure_url" content="https://secure.example.com/sound.mp3"/>`,
				`<meta property="og:audio:type" content="audio/mpeg"/>`,
				`<meta property="og:description" content="Sean Connery found fame and fortune as the suave, sophisticated British agent, James Bond."/>`,
				`<meta property="og:determiner" content="the"/>`,
				`<meta property="og:locale" content="en_GB"/>`,
				`<meta property="og:locale:alternate" content="fr_FR"/>`,
				`<meta property="og:locale:alternate" content="es_ES"/>`,
				`<meta property="og:site_name" content="IMDb"/>`,
				`<meta property="og:video" content="https://example.com/movie.swf"/>`,
				`<meta property="og:video:secure_url" content="https://secure.example.com/movie.swf"/>`,
				`<meta property="og:video:type" content="application/x-shockwave-flash"/>`,
				`<meta property="og:video:width" content="400"/>`,
				`<meta property="og:video:height" content="300"/>`,
			},
		},
	} {
		t.Run(tn, func(t *testing.T) {
			nodes := tc.obj.HTML()
			var actual []string
			for _, node := range nodes {
				var buf bytes.Buffer
				html.Render(&buf, node)
				actual = append(actual, buf.String())
			}
			if !slices.Equal(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
