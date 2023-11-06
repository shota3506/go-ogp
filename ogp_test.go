package ogp_test

import (
	"bytes"
	"slices"
	"strings"
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

func TestParse(t *testing.T) {
	const body = `
<!DOCTYPE html>
<html>
<head>
<title>The Rock (1996)</title>
<meta property="og:title" content="The Rock" />
<meta property="og:type" content="video.movie" />
<meta property="og:url" content="https://www.imdb.com/title/tt0117500/" />
<meta property="og:image" content="https://example.com/rock.jpg" />
<meta property="og:image:secure_url" content="https://secure.example.com/rock.jpg" />
<meta property="og:image:type" content="image/jpeg" />
<meta property="og:image:width" content="400" />
<meta property="og:image:height" content="300" />
<meta property="og:image:alt" content="A shiny red apple with a bite taken out" />
<meta property="og:image" content="https://example.com/rock2.jpg" />
<meta property="og:image" content="https://example.com/rock3.jpg" />
<meta property="og:image:height" content="1000" />
<meta property="og:audio" content="https://example.com/sound.mp3" />
<meta property="og:audio:secure_url" content="https://secure.example.com/sound.mp3" />
<meta property="og:audio:type" content="audio/mpeg" />
<meta property="og:description" 
  content="Sean Connery found fame and fortune as the suave, sophisticated British agent, James Bond." />
<meta property="og:determiner" content="the" />
<meta property="og:locale" content="en_GB" />
<meta property="og:locale:alternate" content="fr_FR" />
<meta property="og:locale:alternate" content="es_ES" />
<meta property="og:site_name" content="IMDb" />
<meta property="og:video" content="https://example.com/movie.swf" />
<meta property="og:video:secure_url" content="https://secure.example.com/movie.swf" />
<meta property="og:video:type" content="application/x-shockwave-flash" />
<meta property="og:video:width" content="400" />
<meta property="og:video:height" content="300" />
</head>
</html>
`

	obj, err := Parse(strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if obj == nil {
		t.Fatal("Object is nil")
	}
	if expected := "The Rock"; obj.Title != expected {
		t.Errorf("Object.Title is %s, expected %s", obj.Title, expected)
	}
	if expected := "video.movie"; obj.Type != expected {
		t.Errorf("Object.Type is %s, expected %s", obj.Type, expected)
	}

	// Images
	if len(obj.Images) != 3 {
		t.Fatalf("Object.Images length is %d, expected %d", len(obj.Images), 3)
	}
	for i, expected := range []*Image{
		{
			URL:       "https://example.com/rock.jpg",
			SecureURL: "https://secure.example.com/rock.jpg",
			Type:      "image/jpeg",
			Width:     400,
			Height:    300,
			Alt:       "A shiny red apple with a bite taken out",
		},
		{URL: "https://example.com/rock2.jpg"},
		{URL: "https://example.com/rock3.jpg", Height: 1000},
	} {
		image := obj.Images[i]
		if image.URL != expected.URL {
			t.Errorf("Object.Images[%d].URL is %s, expected %s", i, image.URL, expected.URL)
		}
		if image.SecureURL != expected.SecureURL {
			t.Errorf("Object.Images[%d].SecureURL is %s, expected %s", i, image.SecureURL, expected.SecureURL)
		}
		if image.Type != expected.Type {
			t.Errorf("Object.Images[%d].Type is %s, expected %s", i, image.Type, expected.Type)
		}
		if image.Width != expected.Width {
			t.Errorf("Object.Images[%d].Width is %d, expected %d", i, image.Width, expected.Width)
		}
		if image.Height != expected.Height {
			t.Errorf("Object.Images[%d].Height is %d, expected %d", i, image.Height, expected.Height)
		}
		if image.Alt != expected.Alt {
			t.Errorf("Object.Images[%d].Alt is %s, expected %s", i, image.Alt, expected.Alt)
		}
	}

	if expected := "https://www.imdb.com/title/tt0117500/"; obj.URL != expected {
		t.Errorf("Object.URL is %s, expected %s", obj.URL, expected)
	}

	// Audios
	if len(obj.Audios) != 1 {
		t.Fatalf("Object.Audios length is %d, expected %d", len(obj.Audios), 1)
	}
	for i, expected := range []*Audio{
		{
			URL:       "https://example.com/sound.mp3",
			SecureURL: "https://secure.example.com/sound.mp3",
			Type:      "audio/mpeg",
		},
	} {
		audio := obj.Audios[i]
		if audio.URL != expected.URL {
			t.Errorf("Object.Audios[%d].URL is %s, expected %s", i, audio.URL, expected.URL)
		}
		if audio.SecureURL != expected.SecureURL {
			t.Errorf("Object.Audios[%d].SecureURL is %s, expected %s", i, audio.SecureURL, expected.SecureURL)
		}
		if audio.Type != expected.Type {
			t.Errorf("Object.Audios[%d].Type is %s, expected %s", i, audio.Type, expected.Type)
		}
	}

	if expected := "Sean Connery found fame and fortune as the suave, sophisticated British agent, James Bond."; obj.Description != expected {
		t.Errorf("Object.Description is %s, expected %s", obj.Description, expected)
	}
	if expected := "the"; obj.Determiner != expected {
		t.Errorf("Object.Determiner is %s, expected %s", obj.Determiner, expected)
	}
	if expected := "en_GB"; obj.Locale.Locale != expected {
		t.Errorf("Object.Locale.Locale is %s, expected %s", obj.Locale.Locale, expected)
	}
	if expected := []string{"fr_FR", "es_ES"}; !slices.Equal(obj.Locale.Alternates, expected) {
		t.Errorf("Object.Locale.Alternate is %s, expected %s", obj.Locale.Alternates, expected)
	}
	if expected := "IMDb"; obj.SiteName != expected {
		t.Errorf("Object.SiteName is %s, expected %s", obj.SiteName, expected)
	}

	// Videos
	if len(obj.Videos) != 1 {
		t.Fatalf("Object.Videos length is %d, expected %d", len(obj.Videos), 1)
	}
	for i, expected := range []*Video{
		{
			URL:       "https://example.com/movie.swf",
			SecureURL: "https://secure.example.com/movie.swf",
			Type:      "application/x-shockwave-flash",
			Width:     400,
			Height:    300,
		},
	} {
		video := obj.Videos[i]
		if video.URL != expected.URL {
			t.Errorf("Object.Videos[%d].URL is %s, expected %s", i, video.URL, expected.URL)
		}
		if video.SecureURL != expected.SecureURL {
			t.Errorf("Object.Videos[%d].SecureURL is %s, expected %s", i, video.SecureURL, expected.SecureURL)
		}
		if video.Type != expected.Type {
			t.Errorf("Object.Videos[%d].Type is %s, expected %s", i, video.Type, expected.Type)
		}
		if video.Width != expected.Width {
			t.Errorf("Object.Videos[%d].Width is %d, expected %d", i, video.Width, expected.Width)
		}
		if video.Height != expected.Height {
			t.Errorf("Object.Videos[%d].Height is %d, expected %d", i, video.Height, expected.Height)
		}
		if video.Alt != expected.Alt {
			t.Errorf("Object.Videos[%d].Alt is %s, expected %s", i, video.Alt, expected.Alt)
		}
	}
}
