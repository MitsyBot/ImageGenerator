package miscgenerator

import (
	"image"
	"trailgen/src/generator"

	"github.com/fogleman/gg"
)

type ShipData struct {
	LinkOne   string   `json:"linkOne" form:"linkOne"`
	LinkTwo   string   `json:"linkTwo" form:"linkTwo"`
	ShipValue float64  `json:"value" form:"value"`
	ShipPercentage string  `json:"percentage" form:"percentage"`
}

func RenderShip(g generator.Generator, p *ShipData) image.Image {
	dc := gg.NewContext(730, 346)

	dc.SetHexColor("#ae37ee")
	dc.DrawRoundedRectangle(0, 0, 730, 346, float64(3))
	dc.Fill()

	avatarSize := 189
	circleSize := float64(90)

	dc.SetHexColor("#F895FB")
	dc.DrawCircle(float64(143), float64(143), 96)
	dc.Fill()

	avatarOne := g.Toolbox.ReadImageFromURL(p.LinkOne, avatarSize, avatarSize)
	dc.DrawCircle(float64(143), float64(143), circleSize)
	dc.Clip()
	dc.DrawImageAnchored(avatarOne, 143, 143, 0.5, 0.5)
	dc.ResetClip()


	Heart := g.Toolbox.ReadImageFromURL("https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/twitter/322/red-heart_2764-fe0f.png", 189, 189)
	dc.DrawImage(Heart, 265, 45);

	dc.SetHexColor("#F895FB")
	dc.DrawRoundedRectangle(0, 0, 730, 9, float64(3))
	dc.Fill()

	dc.SetHexColor("#bc5bf1")
	dc.DrawRoundedRectangle(96, 276, 528, 48, float64(3))
	dc.Fill()


	dc.SetHexColor("#E8E8E8")
	dc.DrawRoundedRectangle(100, 280, 520, 40, float64(3))
	dc.Fill()

	dc.SetHexColor("#F895FB")
	dc.DrawRoundedRectangle(100, 280, p.ShipValue, 40, float64(3))
	dc.Fill()

	dc.SetHexColor("#F895FB")
	dc.DrawCircle(float64(580), float64(143), 96)
	dc.Fill()

	dc.LoadFontFace("./fonts/Inter/semibold.ttf", 20)
	dc.SetHexColor("#3e4451")
	dc.DrawString(p.ShipPercentage, float64(338), float64(307))

	avatarTwo := g.Toolbox.ReadImageFromURL(p.LinkTwo, avatarSize, avatarSize)
	dc.DrawCircle(float64(580), float64(143), circleSize)
	dc.Clip()
	dc.DrawImageAnchored(avatarTwo, 580, 143, 0.5, 0.5)
	dc.ResetClip()

	return dc.Image()
}
