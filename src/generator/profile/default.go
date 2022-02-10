package profilegenerator

import (
	"image"
	"trailgen/src/generator"

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
	Level 		string   `json:"level" form:"level"`
	Xp 			string   `json:"xp" form:"xp"`
	Rep 		string   `json:"rep" form:"rep"`
	Money       string   `json:"money" form:"money"`
	Type        string   `json:"type" form:"type"`
	Badges      []string `json:"badges" form:"badges"`
	PartnerName string   `json:"partnerName" form:"partnerName"`
	AvatarIcon  string   `json:"avatarIcon" form:"avatarIcon"`
}

func RenderProfile(g generator.Generator, p *ProfileData) image.Image {
	dc := gg.NewContext(1100, 720)

	/* BACKGROUND */
	img := g.Toolbox.GetAsset("bgs/profile/" + p.Background)

	dc.DrawImage(img, 260, 0)
	dc.LoadFontFace("./fonts/Ghost/iknowaghost.ttf", 50)
	img = g.Toolbox.GetAsset("template/profile")
	dc.DrawImage(img, 0, 0)

	/* AVATAR DRAWING */
	avatarSize := 240
	circleSize := float64(116)
	x := 197
	y := 143

	/* Outline of avatar	*/
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

	/* NickName	*/
	dc.LoadFontFace("./fonts/Inter/semibold.ttf", 30)
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
	dc.DrawString(p.Name, float64(120), float64(620))


	/* Dinheiros */
	dc.LoadFontFace("./fonts/Monstserrat/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, "Cometinhas", 378, 320, 208, 408, 13)

	dc.LoadFontFace("./fonts/Inter/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, p.Money, 640, 320, 208, 408, 13)


	/* Rep */
	dc.LoadFontFace("./fonts/Monstserrat/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, "Reputação", 378, 370, 208, 408, 13)
	
	dc.LoadFontFace("./fonts/Inter/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, p.Rep, 640, 370, 208, 408, 13)
	
	/* ProgressBar */

	dc.LoadFontFace("./fonts/Monstserrat/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, "Experiencia", 378, 466, 208, 408, 13)
	
	dc.LoadFontFace("./fonts/Monstserrat/Montserrat-Medium.ttf", 24) // Default is 40px
	g.Toolbox.DrawTextWrapped(dc, p.Xp, 681, 466, 208, 408, 13)
	
	dc.SetHexColor("#E8E8E8")
	dc.DrawRoundedRectangle(378, 493, 366, 9, float64(3))
	dc.Fill()

	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(378, 493, 200, 9, float64(3))
	dc.Fill()

	/* Nivel */
	dc.SetHexColor("#FFFFFF")
	dc.LoadFontFace("../assets/fonts/Montserrat/Montserrat-Regular.ttf", 25)
	dc.DrawString("Nivel", float64(990), float64(322))

	dc.LoadFontFace("./fonts/Monstserrat/Montserrat-Regular.ttf", 26) // Default is 40px
	dc.DrawString(p.Level, float64(1010), float64(366))

		
	/* About Me	*/
	dc.LoadFontFace("./fonts/Inter/medium.ttf", 23)
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
	g.Toolbox.DrawTextWrapped(dc, p.AboutMe, 25, 671, 500, 250, 28)

	/* Married	*/
	if p.Married {

		/* Load Bar Float */
		img := g.Toolbox.GetAsset("/template/profile_ship")
		dc.DrawImage(img, 0, 0)

		/* Partner */
		dc.SetHexColor(g.Toolbox.GetCompatibleFontColor("#ffff"))
		dc.DrawString(p.PartnerName, 51, 402)

	}

	/* BADGES */

	bx := 1000.0
	by := 390.0
	badgesizex := 35.0
	badgesizey := 30.0
	badgespacing := 65.0 // <=
	spacebtwedge := 10.0
	recsizex := 450.0

	recsizey := badgesizey*2 + badgespacing

	cxpos := bx + spacebtwedge
	cypos := by
	nb := 0

	for _, b := range p.Badges {
		nb++
		if nb == 7 {
		} else {
			if b != "" {
				dc.DrawImage(g.Toolbox.GetAsset("badges/profile/"+b), int(cxpos), int(cypos))
			}

			cypos = cypos + badgesizey + badgespacing
			if cypos > by+spacebtwedge+recsizey {
				cxpos = cxpos - badgesizex - badgespacing
				cypos = 390
				if cxpos > bx+spacebtwedge+recsizex {
					break // if we run out of space, break the loop
				}
			}
		}

	}

	return dc.Image()
}
