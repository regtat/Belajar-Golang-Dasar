package main

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	Jari   int `json:"jari-jari-lingkaran"`
	Sisi   int `json:"sisi-persegi"`
	Alas   int `json:"alas-segitiga"`
	Tinggi int `json:"tinggi-segitiga"`
}

type ResponseData struct {
	LuasLingkaran     float64 `json:"luas-Lingkaran"`
	LuasPersegi       float64 `json:"luas-Persegi"`
	LuasSegitiga      float64 `json:"luas-Segitiga"`
	KelilingLingkaran float64 `json:"keliling-Lingkaran"`
	KelilingPersegi   float64 `json:"keliling-Persegi"`
	KelilingSegitiga  float64 `json:"keliling-Segitiga"`
}

func main() {

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(RequestData)
		if err := c.BodyParser(request); err != nil {
			return err
		}
		response := ResponseData{}
		cSegitiga := math.Sqrt(math.Pow(float64(request.Alas), 2) + math.Pow(float64(request.Tinggi), 2))
		Ks := float64(request.Alas + request.Tinggi + int(cSegitiga))

		response.LuasLingkaran = 3.14 * float64(request.Jari*request.Jari)
		response.LuasPersegi = float64(request.Sisi * request.Sisi)
		response.LuasSegitiga = 0.5 * float64(request.Alas*request.Tinggi)
		response.KelilingLingkaran = 2 * 3.14 * float64(request.Jari)
		response.KelilingPersegi = 4 * float64(request.Sisi)
		response.KelilingSegitiga = Ks

		return c.JSON(response)
	})
	app.Listen(":3000")
}
