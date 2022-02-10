package server

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"strconv"
	"time"
	"trailgen/src/generator"
	profilegenerator "trailgen/src/generator/profile"
	"trailgen/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/nfnt/resize"
)

func StartServer(port string) {
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
		ReduceMemoryUsage: true,
		Prefork:           true,
	})
	app.Get("/dashboard", monitor.New(monitor.Config{
		APIOnly: true,
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Add("Cache-Time", "90000")
		return c.Next()
	})

	app.Use(pprof.New())
	app.Use(compress.New(compress.Config{
		Next:  nil,
		Level: compress.LevelBestSpeed, // 1
	}))

	// app.Use(logger.New())
	gen := generator.NewGenerator()
	encoder := png.Encoder{
		CompressionLevel: -2, // no compression.
	}

	app.Static("/static", "./images")
	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString("trailgen v1.1-dev1 (fasthttp/fiber; golang)")
	})
	app.Get("/get_asset_list/:image", func(c *fiber.Ctx) error {
		if _, found := utils.Find([]string{"badges", "stickers", "bgs"}, c.Params("image")); !found {
			return c.SendString("não!")
		}

		files, err := ioutil.ReadDir("./images/" + c.Params("image"))
		if err != nil {
			panic(err)
			return c.SendString("[]")
		}

		return c.JSON(utils.FilterFileList(files))
	})

	app.Post("/render/profile", func(c *fiber.Ctx) error {
		p := new(profilegenerator.ProfileData)

		if err := c.BodyParser(p); err != nil {
			fmt.Println(err)
			return err
		}

		var img image.Image

		switch p.Type {
		case "default":
			c.Set("Content-Type", "image/png")
			img = profilegenerator.RenderProfile(gen, p)
		default:
			c.Set("Content-Type", "image/png")
			img = profilegenerator.RenderProfile(gen, p)
		}

		if c.Query("w", "0") != "0" {
			if c.Query("h", "0") != "0" {
				parseW, err := strconv.ParseUint(c.Query("w", "0"), 10, 32)
				if err != nil {
					fmt.Println(err)
					return err
				}
				parseH, err := strconv.ParseUint(c.Query("h", "0"), 10, 32)
				if err != nil {
					fmt.Println(err)
					return err
				}

				if c.Query("type", "0") != "0" {
					switch c.Query("type", "nil") {
					case "thumb":
						newImage := resize.Thumbnail(uint(parseW), uint(parseH), img, resize.Lanczos3)
						img = newImage

					case "resize":
						newImage := resize.Resize(uint(parseW), uint(parseH), img, resize.Lanczos3)
						img = newImage
					default:
						c.Set("Content-Type", "text/plain; charset=utf-8")
						return c.SendString("😳 Você esqueceu alguma coisa.")
					}
				}

			}
		}

		return encoder.Encode(c.Context(), img)
	})

	app.Listen(":" + port)
}
