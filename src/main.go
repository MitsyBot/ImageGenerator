package main

import (
	"trailgen/src/server"
)

func main() {
	/*d := profilegenerator.ProfileData {
	  Name: "test",
	  AboutMe: "Test ",
	  AvatarURL: "https://cdn.discordapp.com/attachments/504668288798949376/771567079538688050/bailey_hat.jpeg",
	  FavColor: "ff0000",
	  Money: "1.78m",
	  Background: "mitsy_1",
	  Married: false,
	  PartnerName: "DSHQDIQWJDIUOQW",
	}

	gen.Render()*/
	server.StartServer("2080")
}
