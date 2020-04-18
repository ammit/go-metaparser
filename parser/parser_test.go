package parser_test

import (
	"strings"
	"testing"

	"github.com/ammit/go-metaparser/parser"
)

const html = `
<!doctype html>
<html class="no-js" lang="">

<head>
	<meta charset="utf-8">
	<title>Go Meta Parser</title>
	<meta name="description" content="Go Meta Description">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="manifest" href="site.webmanifest">
	<link rel="apple-touch-icon" href="icon.png">
	<link rel="stylesheet" href="css/normalize.css">
	<link rel="stylesheet" href="css/main.css">

	<!-- og:basic -->
	<meta property="og:title" content="sample title" />
	<meta property="og:description" content="sample description" />
	<meta property="og:determiner" content="determiner" />
	<meta property="og:type" content="demo" />
	<meta property="og:url" content="http://example.com" />
	<meta property="og:locale" content="en_GB" />
	<meta property="og:locale:alternate" content="fr_FR" />
	<meta property="og:locale:alternate" content="es_ES" />
	<meta property="og:site_name" content="sample site_name" />

	<!-- og:image -->
	<meta property="og:image" content="http://example.com/ogp.jpg" />
	<meta property="og:image:secure_url" content="https://secure.example.com/ogp.jpg" />
	<meta property="og:image:type" content="image/jpeg" />
	<meta property="og:image:width" content="400" />
	<meta property="og:image:height" content="300" />
	<meta property="og:image:alt" content="A shiny red apple with a bite taken out" />

	<!-- og:video -->
	<meta property="og:video" content="http://example.com/movie.swf" />
	<meta property="og:video:secure_url" content="https://secure.example.com/movie.swf" />
	<meta property="og:video:type" content="application/x-shockwave-flash" />
	<meta property="og:video:width" content="400" />
	<meta property="og:video:height" content="300" />

	<!-- og:audio -->
	<meta property="og:audio" content="http://example.com/sound.mp3" />
	<meta property="og:audio:secure_url" content="https://secure.example.com/sound.mp3" />
	<meta property="og:audio:type" content="audio/mpeg" />

	<!-- (new) music:music/album/musician/playlist/radio_station -->
	<meta property="music:musician" content="http://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d">
	<meta property="music:musician" content="http://open.spotify.com/artist/0oSGxfWSnnOXhD2fKuz2Gy">
	<meta property="music:album" content="http://open.spotify.com/album/7rq68qYz66mNdPfidhIEFa">
	<meta property="music:album:track" content="2">
	<meta property="music:duration" content="236">
	<meta property="music:release_date" content="2011-04-19">     
	<meta property="music:song" content="http://open.spotify.com/track/2aSFLiDPreOVP6KHiWk4lF">
	<meta property="music:song:disc" content="1">
	<meta property="music:song:track" content="2">
	<meta property="music:creator" content="http://open.spotify.com/user/austinhaugen"/>
</head>
<body>
</body>
`

func TestParserParseHTML(t *testing.T) {
	p := parser.New("")
	err := p.ParseHTML(strings.NewReader(html))

	if err != nil {
		t.Fatal(err)
	}

	if p.Title != "sample title" {
		t.Error("title parsed incorrectly")
	}

	if p.Description != "sample description" {
		t.Error("description parsed incorrectly")
	}

	if p.Determiner != "determiner" {
		t.Error("determiner parsed incorrectly")
	}

	if p.Type != "demo" {
		t.Error("type parsed incorrectly")
	}

	if p.URL != "http://example.com" {
		t.Error("url parsed incorrectly")
	}

	if p.Locale != "en_GB" {
		t.Error("locale parsed incorrectly")
	}

	if len(p.LocalesAlternate) == 0 {
		t.Error("locales_alternate parsed incorrectly")
	}

	if p.SiteName != "sample site_name" {
		t.Error("site_name parsed incorrectly")
	}

	// og:image
	if len(p.Images) == 0 {
		t.Error("images parsed incorrectly")
	} else {
		if len(p.Images[0].URL) == 0 {
			t.Error("image url parsed incorrectly")
		}

		if len(p.Images[0].SecureURL) == 0 {
			t.Error("image secure_url parsed incorrectly")
		}

		if len(p.Images[0].Type) == 0 {
			t.Error("image type parsed incorrectly")
		}

		if p.Images[0].Width == 0 {
			t.Error("image width parsed incorrectly")
		}

		if p.Images[0].Height == 0 {
			t.Error("image height parsed incorrectly")
		}

		if len(p.Images[0].Alt) == 0 {
			t.Error("image alt parsed incorrectly")
		}
	}

	// og:video
	if len(p.Videos) == 0 {
		t.Error("videos parsed incorrectly")
	} else {
		if len(p.Videos[0].URL) == 0 {
			t.Error("video url parsed incorrectly")
		}

		if len(p.Videos[0].SecureURL) == 0 {
			t.Error("video secure_url parsed incorrectly")
		}

		if len(p.Videos[0].Type) == 0 {
			t.Error("video type parsed incorrectly")
		}

		if p.Videos[0].Width == 0 {
			t.Error("video width parsed incorrectly")
		}

		if p.Videos[0].Height == 0 {
			t.Error("video height parsed incorrectly")
		}
	}

	// og:audio
	if len(p.Audios) == 0 {
		t.Error("audios parsed incorrectly")
	} else {
		if len(p.Audios[0].URL) == 0 {
			t.Error("audio url parsed incorrectly")
		}

		if len(p.Audios[0].SecureURL) == 0 {
			t.Error("audio secure_url parsed incorrectly")
		}

		if len(p.Audios[0].Type) == 0 {
			t.Error("audio type parsed incorrectly")
		}
	}

	// music
	if len(p.Music.Musicians) == 0 {
		t.Error("musicians parsed incorrectly")
	}

	// music.album
	if len(p.Music.Album.URL) == 0 {
		t.Error("Music Album URL parsed incorrectly")
	}

	if p.Music.Album.Track == 0 {
		t.Error("Music Album Track parsed incorrectly")
	}

	if p.Music.Duration == 0 {
		t.Error("Music duration parsed incorrectly")
	}

	if len(p.Music.ReleaseDate) == 0 {
		t.Error("Music Release Date parsed incorrectly")
	}

	if len(p.Music.Creator) == 0 {
		t.Error("Music creator parsed incorrectly")
	}

	if len(p.Music.Songs) == 0 {
		t.Error("Music songs parsed incorrectly")
	} else {
		if len(p.Music.Songs[0].URL) == 0 {
			t.Error("Music song url parsed incorrectly")
		}

		if p.Music.Songs[0].Disc == 0 {
			t.Error("Music song disc parsed incorrectly")
		}

		if p.Music.Songs[0].Track == 0 {
			t.Error("Music song track parsed incorrectly")
		}
	}
}
