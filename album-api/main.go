package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// السماح للويب + الموبايل يتصلون على الـ API
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // بعدين تقدر تضبطها على الدومينات حقك
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// اختبار بسيط
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// مثال مبدئي لـ حساب سعر ألبوم (بشكل ثابت الآن)
	app.Post("/pricing", func(c *fiber.Ctx) error {
		type AlbumConfig struct {
			Size        string `json:"size"`
			Cover       string `json:"cover"`
			ImagesCount int    `json:"imagesCount"`
		}

		var cfg AlbumConfig
		if err := c.BodyParser(&cfg); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid body",
			})
		}

		price := 0.0

		// السعر الأساسي حسب الحجم
		switch cfg.Size {
		case "30x30":
			price += 20
		case "20x20":
			price += 12
		default:
			price += 15 // افتراضي مثلاً
		}

		// عدد الصور الإضافية
		baseImages := 30
		if cfg.ImagesCount > baseImages {
			extra := cfg.ImagesCount - baseImages
			price += float64(extra) * 0.3
		}

		// الغلاف
		if cfg.Cover == "جلد" {
			price += 8
		}

		return c.JSON(fiber.Map{
			"price": price,
		})
	})

	log.Println("Server running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
