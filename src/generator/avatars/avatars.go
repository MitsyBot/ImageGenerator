package avatarsgenerator

import (
	"image"
	"trailgen/src/generator"

	"github.com/fogleman/gg"
)

type AvatarsData struct {
	oldAvatar   string   `json:"oldAvatar" form:"oldAvatar"`
	newAvatar  string   `json:"newAvatar" form:"newAvatar"`
}

func RenderAvatars(g generator.Generator, p *AvatarsData) image.Image {
	dc := gg.NewContext(256, 128)

	oldAvatar := g.Toolbox.ReadImageFromURL(p.oldAvatar, 128, 128)
	newAvatar := g.Toolbox.ReadImageFromURL(p.newAvatar, 128, 128)

	dc.DrawImage(oldAvatar, 0, 0);
	dc.DrawImage(newAvatar, 128, 0);

	return dc.Image()
}
