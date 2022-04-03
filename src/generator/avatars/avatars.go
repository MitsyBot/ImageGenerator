package avatarsgenerator

import (
	"image"
	"trailgen/src/generator"

	"github.com/fogleman/gg"
)

type ProfileData struct {
	OldAvatar   string   `json:"oldAvatar" form:"oldAvatar"`
	NewAvatar   string   `json:"newAvatar" form:"newAvatar"`
}

func RenderAvatars(g generator.Generator, p *AvatarsData) image.Image {
	dc := gg.NewContext(256, 128)

	oldAvatar := g.Toolbox.ReadImageFromURL(p.OldAvatar, 128, 128)
	newAvatar := g.Toolbox.ReadImageFromURL(p.NewAvatar, 128, 128)

	dc.DrawImage(oldAvatar, 0, 0);
	dc.DrawImage(newAvatar, 128, 0);

	return dc.Image()
}
