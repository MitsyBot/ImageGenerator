package profilegenerator

import (
	"image"
	"generator/src/generator"
	"github.com/fogleman/gg"
)

type ProfileData struct {
	AvatarURL   string   `json:"avatarUrl" form:"avatarUrl"`
	Background  string   `json:"bgId" form:"bgId"`
	Reps        string   `json:"reps" form:"reps"`           // New field
	Married     bool     `json:"married" form:"married"`
	Name        string   `json:"name" form:"name"`
	AboutMe     string   `json:"aboutMe" form:"aboutMe"`
	FavColor    string   `json:"favColor" form:"favColor"`
	Money       string   `json:"money" form:"money"`
	Type        string   `json:"type" form:"type"`
	Badges      []string `json:"badges" form:"badges"`
	PartnerName string   `json:"partnerName" form:"partnerName"`
	AvatarIcon  string   `json:"avatarIcon" form:"avatarIcon"`
}


func RenderDefaulProfile(g generator.Generator, p *ProfileData) image.Image {
	dc := gg.NewContext(1100, 720)

	img := g.Toolbox.GetAsset("bgs/" + p.Background)

	dc.DrawImage(img, 260, 0)
	dc.LoadFontFace("../assets/fonts/Ghost/iknowaghost.ttf", 50)
	img = g.Toolbox.GetAsset("template/profile.png")
	dc.DrawImage(img, 0, 0)

	dc.SetHexColor(p.FavColor)
	dc.SetLineWidth(6)
	dc.DrawLine(-250, 260, float64(1120), float64(260))
	dc.Stroke()

	avatarSize := 240
	circleSize := float64(116)
	x := 155
	y := 126

	dc.SetHexColor("ffffff")
	dc.DrawCircle(float64(x), float64(y), 77)
	dc.Fill()

	dc.SetHexColor(p.FavColor)
	dc.DrawCircle(float64(x), float64(y), 125)
	dc.Fill()

	avatar := g.Toolbox.ReadImageFromURL(p.AvatarURL, avatarSize, avatarSize)
	dc.DrawCircle(float64(x), float64(y), circleSize)
	dc.Clip()
	dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
	dc.ResetClip()

	if !(p.AvatarIcon == "") {
		xIcon := 50
		yIcon := 200
		avatarIcon := g.Toolbox.ReadImageFromURL(p.AvatarIcon, 100, 100)
		dc.DrawCircle(float64(xIcon), float64(yIcon), float64(40))
		dc.Clip()
		dc.DrawImageAnchored(avatarIcon, xIcon, yIcon, 0.5, 0.5)
		dc.ResetClip()
	}

	dc.LoadFontFace("../assets/fonts/Montserrat/Montserrat-ExtraLight.ttf", 30)
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
	dc.DrawString(p.Name, float64(25), float64(295))

	dc.LoadFontFace("../assets/fonts/Montserrat/Montserrat-ExtraLight.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, p.Money, 520, 306, 208, 408, 13)

	dc.LoadFontFace("../assets/fonts/Montserrat/Montserrat-Regular.ttf", 23)
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
	g.Toolbox.DrawTextWrapped(dc, p.AboutMe, 25, 377, 500, 250, 28)

	if p.Married {

		img := g.Toolbox.GetAsset("/template/profile_ship")
		dc.DrawImage(img, 0, 0)
		dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
		dc.DrawString(p.PartnerName, 659, 380)

	}

	bx := 47.0
	by := 479.0
	badgesizex := 35.0
	badgesizey := 30.0
	badgespacing := 65.0 // <=
	spacebtwedge := 5.0
	recsizex := 450.0

	recsizey := badgesizey*2 + badgespacing

	cxpos := bx + spacebtwedge
	cypos := by + spacebtwedge
	nb := 0

	for _, b := range p.Badges {
		nb++
		if nb == 100 {
		} else {
			if b != "" {
				dc.DrawImage(g.Toolbox.GetAsset("badges/"+b), int(cxpos), int(cypos))
			}

			cxpos = cxpos + badgesizex + badgespacing
			if cxpos > bx+spacebtwedge+recsizex {
				cypos = cypos + badgesizey + badgespacing
				cxpos = bx + spacebtwedge
				if cypos > by+spacebtwedge+recsizey {
					break
				}
			}
		}

	}

	return dc.Image()
}
