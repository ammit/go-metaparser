package parser_test

import (
	"io/ioutil"
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

	<!-- (new) video:actor/movie/episod -->
	<meta property="video:actor" content="http://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d" />
	<meta property="video:actor:role" content="xyz">
	<meta property="video:actor" content="http://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d" />
	<meta property="video:actor:role" content="xyz">
	<meta property="video:director" content="http://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d" />
	<meta property="video:writer" content="http://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d" />
	<meta property="video:duration" content="1236">
	<meta property="video:release_date" content="1236">
	<meta property="video:tag" content="a">
	<meta property="video:tag" content="b">
	<meta property="video:series" content="http://open.spotify.com/series/1dfeR4HaWDbWqFHLkxsg1d">

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

	<!-- article: -->
	<meta property="article:published_time" content="test" />
	<meta property="article:modified_time" content="test" />
	<meta property="article:expiration_time" content="test" />
	<meta property="article:author" content="test" />
	<meta property="article:author" content="test2" />
	<meta property="article:section" content="test" />
	<meta property="article:tag" content="test" />
	<meta property="article:tag" content="tes2" />


	<!-- book: -->
	<meta property="book:author" content="book author" />
	<meta property="book:isbn" content="isbn" />
	<meta property="book:release_date" content="test" />
	<meta property="book:tag" content="test" />
	<meta property="book:tag" content="tes2" />

	<!-- profile: -->
	<meta property="profile:first_name" content="firstname" />
	<meta property="profile:last_name" content="lastname" />
	<meta property="profile:username" content="username" />
	<meta property="profile:gender" content="male" />

	<!-- twitter: -->
	<meta property="twitter:card" content="test" />
	<meta property="twitter:site" content="test" />
	<meta property="twitter:site:id" content="test" />
	<meta property="twitter:creator" content="test" />
	<meta property="twitter:creator:id" content="test" />
	<meta property="twitter:description" content="test" />
	<meta property="twitter:title" content="test" />
	<meta property="twitter:image" content="test" />
	<meta property="twitter:image:alt" content="test" />
	<meta property="twitter:player" content="test" />
	<meta property="twitter:player:height" content="10" />
	<meta property="twitter:player:width" content="10" />
	<meta property="twitter:player:stream" content="test" />
	<meta property="twitter:app:name:iphone" content="test" />
	<meta property="twitter:app:url:iphone" content="test" />
	<meta property="twitter:app:id:iphone" content="1" />
	<meta property="twitter:app:name:ipad" content="test" />
	<meta property="twitter:app:url:ipad" content="test" />
	<meta property="twitter:app:id:ipad" content="1" />
	<meta property="twitter:app:name:googleplay" content="test" />
	<meta property="twitter:app:url:googleplay" content="test" />
	<meta property="twitter:app:id:googleplay" content="1" />
</head>
<body>
</body>
`

func TestParserParseHTML(t *testing.T) {
	p := parser.New()
	err := p.ParseHTML(ioutil.NopCloser(strings.NewReader(html)))

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

		// video:
		if len(p.Videos[0].Actors) == 0 {
			t.Error("video actors parsed incorrectly")
		} else {
			if len(p.Videos[0].Actors[0].URL) == 0 {
				t.Error("video actors url parsed incorrectly")
			}
			if len(p.Videos[0].Actors[0].Role) == 0 {
				t.Error("video actors role parsed incorrectly")
			}
		}

		if len(p.Videos[0].Director) == 0 {
			t.Error("video director parsed incorrectly")
		}

		if len(p.Videos[0].Writer) == 0 {
			t.Error("video writer parsed incorrectly")
		}

		if p.Videos[0].Duration == 0 {
			t.Error("video duration parsed incorrectly")
		}

		if len(p.Videos[0].ReleaseDate) == 0 {
			t.Error("video release_date parsed incorrectly")
		}

		if len(p.Videos[0].Series) == 0 {
			t.Error("video series parsed incorrectly")
		}

		if len(p.Videos[0].Tags) == 0 {
			t.Error("video tags parsed incorrectly")
		} else {
			if len(p.Videos[0].Tags[0]) == 0 {
				t.Error("video tags tag parsed incorrectly")
			}
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

	// article
	if len(p.Article.PublishedTime) == 0 {
		t.Error("article published_time parsed incorrectly")
	}

	if len(p.Article.ModifiedTime) == 0 {
		t.Error("article modified_time parsed incorrectly")
	}

	if len(p.Article.ExpirationTime) == 0 {
		t.Error("article expiration_time parsed incorrectly")
	}

	if len(p.Article.Section) == 0 {
		t.Error("article section parsed incorrectly")
	}

	if len(p.Article.Tags) == 0 {
		t.Error("Article tags parsed incorrectly")
	} else {
		if len(p.Article.Tags[0]) == 0 {
			t.Error("Article tags tag parsed incorrectly")
		}
	}

	if len(p.Article.Authors) == 0 {
		t.Error("Article Authors parsed incorrectly")
	} else {
		if len(p.Article.Authors[0]) == 0 {
			t.Error("Article Authors Author parsed incorrectly")
		}
	}

	// books
	if len(p.Book.ReleaseDate) == 0 {
		t.Error("article released_date parsed incorrectly")
	}

	if len(p.Book.Isbn) == 0 {
		t.Error("article isbin parsed incorrectly")
	}

	if len(p.Book.Tags) == 0 {
		t.Error("Book tags parsed incorrectly")
	} else {
		if len(p.Book.Tags[0]) == 0 {
			t.Error("Book tags tag parsed incorrectly")
		}
	}

	if len(p.Book.Authors) == 0 {
		t.Error("Book Authors parsed incorrectly")
	} else {
		if len(p.Book.Authors[0]) == 0 {
			t.Error("Book Authors Author parsed incorrectly")
		}
	}

	// profile
	if len(p.Profile.FirstName) == 0 {
		t.Error("profile first_name parsed incorrectly")
	}

	if len(p.Profile.LastName) == 0 {
		t.Error("profile last_name parsed incorrectly")
	}

	if len(p.Profile.Username) == 0 {
		t.Error("profile username parsed incorrectly")
	}

	if len(p.Profile.Gender) == 0 {
		t.Error("profile gender parsed incorrectly")
	}

	// twitter
	if len(p.Twitter.Card) == 0 {
		t.Error("twitter:card parsed incorrectly")
	}

	if len(p.Twitter.Site) == 0 {
		t.Error("twitter:site parsed incorrectly")
	}

	if len(p.Twitter.SiteID) == 0 {
		t.Error("twitter:site:id parsed incorrectly")
	}

	if len(p.Twitter.Creator) == 0 {
		t.Error("twitter:creator parsed incorrectly")
	}

	if len(p.Twitter.CreatorID) == 0 {
		t.Error("twitter:creator:id parsed incorrectly")
	}

	if len(p.Twitter.Description) == 0 {
		t.Error("twitter:description parsed incorrectly")
	}

	if len(p.Twitter.Title) == 0 {
		t.Error("twitter:title parsed incorrectly")
	}

	if len(p.Twitter.Image) == 0 {
		t.Error("twitter:image parsed incorrectly")
	}

	if len(p.Twitter.ImageAlt) == 0 {
		t.Error("twitter:image:alt parsed incorrectly")
	}

	if len(p.Twitter.Player.URL) == 0 {
		t.Error("twitter:player parsed incorrectly")
	}

	if p.Twitter.Player.Height == 0 {
		t.Error("twitter:player:height parsed incorrectly")
	}

	if p.Twitter.Player.Width == 0 {
		t.Error("twitter:player:width parsed incorrectly")
	}

	if len(p.Twitter.Player.Stream) == 0 {
		t.Error("twitter:player:stream parsed incorrectly")
	}

	// log.Fatalf("%+v", p.Twitter.Apps)

	if len(p.Twitter.Apps[0].Name) == 0 {
		t.Error("twitter:app:name:iphone parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].URL) == 0 {
		t.Error("twitter:app:url:iphone parsed incorrectly")
	}

	if p.Twitter.Apps[0].ID == 0 {
		t.Error("twitter:app:id:iphone parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].Type) == 0 {
		t.Error("twitter:app:type:iphone parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].Name) == 0 {
		t.Error("twitter:app:name:ipad parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].URL) == 0 {
		t.Error("twitter:app:url:ipad parsed incorrectly")
	}

	if p.Twitter.Apps[0].ID == 0 {
		t.Error("twitter:app:id:ipad parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].Type) == 0 {
		t.Error("twitter:app:type:ipad parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].Name) == 0 {
		t.Error("twitter:app:name:googleplay parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].URL) == 0 {
		t.Error("twitter:app:url:googleplay parsed incorrectly")
	}

	if p.Twitter.Apps[0].ID == 0 {
		t.Error("twitter:app:id:googleplay parsed incorrectly")
	}

	if len(p.Twitter.Apps[0].Type) == 0 {
		t.Error("twitter:app:type:googleplay parsed incorrectly")
	}

}

func BenchmarkParser(b *testing.B) {
	p := parser.New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = p.ParseHTML(ioutil.NopCloser(strings.NewReader(html)))
	}
}
