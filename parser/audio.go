package parser

// Audio type in Open Graph
type Audio struct {
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Type      string `json:"type"`
}

func (p *Parser) ensureAudio() {
	if len(p.Audios) > 0 {
		return
	}
	p.Audios = append(p.Audios, &Audio{})
}

func (p *Parser) parseAudioMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "og:audio":
		if len(p.Audios) > 0 && len(p.Audios[len(p.Audios)-1].URL) == 0 {
			p.Audios[len(p.Audios)-1].URL = attrs["content"]
		} else {
			p.Audios = append(p.Audios, &Audio{URL: attrs["content"]})
		}
	case "og:audio:secure_url":
		p.ensureAudio()
		p.Audios[len(p.Audios)-1].SecureURL = attrs["content"]
	case "og:audio:type":
		p.ensureAudio()
		p.Audios[len(p.Audios)-1].Type = attrs["content"]
	}
}
