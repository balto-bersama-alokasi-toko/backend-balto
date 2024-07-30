package server

import (
	"backend-balto/handler/usecase/merchant"
	"backend-balto/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
	"strings"
)

type Svr struct {
	merchantUc merchant.Handler
}

func NewServer(merchantUc merchant.Handler) *Svr {
	return &Svr{
		merchantUc: merchantUc,
	}
}

func (s *Svr) StartListening(port int) {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:               5 * 1024 * 1024,
		EnableTrustedProxyCheck: true,
	})

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: strings.Join([]string{
			http.MethodPost,
			http.MethodGet,
		}, ","),
		AllowHeaders: strings.Join([]string{
			"Content-Length",
			"Content-Type",
			"Accept",
		}, ","),
		AllowCredentials: true,
		ExposeHeaders: strings.Join([]string{
			"Content-Length",
			"Content-Type",
			"Accept",
		}, ","),
	}))

	fiberApp.Use(recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			return false
		},
		EnableStackTrace: false,
		StackTraceHandler: func(c *fiber.Ctx, r interface{}) {
			fmt.Println("Recovered. Error:\n", r)
		},
	}))

	fiberApp.Post("/merchant-category", s.GetMerchantPerCategoryHandler)
	fiberApp.Get("/merchant-category-list", s.GetMerchantCategoriesHandler)
	fiberApp.Get("/public-category-list", s.GetPublicCategoriesHandler)
	fiberApp.Post("/location-around", s.GetLocationAroundHandler)
	fiberApp.Post("/kelurahan-detail", s.GetKelurahanDetailHandler)
	fiberApp.Post("/merchant-detail", s.GetMerchantDetailHandler)
	fiberApp.Post("/location-prediction", s.LocationPrediction)

	err := fiberApp.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}

}

func (s *Svr) GetMerchantPerCategoryHandler(c *fiber.Ctx) error {
	var queryBody map[string]string

	err := c.BodyParser(&queryBody)
	if err != nil {
		return err
	}
	merchants, err := s.merchantUc.FindByCategory(queryBody["merchant_category"])
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(merchants)
}

func (s *Svr) GetMerchantCategoriesHandler(c *fiber.Ctx) error {
	categories, err := s.merchantUc.GetMerchantCategories()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

func (s *Svr) GetPublicCategoriesHandler(c *fiber.Ctx) error {
	categories, err := s.merchantUc.GetPublicCategories()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

func (s *Svr) GetLocationAroundHandler(c *fiber.Ctx) error {
	var queryBody map[string]string

	err := c.BodyParser(&queryBody)
	if err != nil {
		return err
	}
	kelurahans, err := s.merchantUc.GetPublicPlaces(queryBody["public_place"])
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(kelurahans)
}

func (s *Svr) GetKelurahanDetailHandler(c *fiber.Ctx) error {
	var queryBody map[string]int

	err := c.BodyParser(&queryBody)
	if err != nil {
		return err
	}
	kelurahans, err := s.merchantUc.GetKelurahanDetails(queryBody["kelurahan_id"])
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(kelurahans)

}

func (s *Svr) GetMerchantDetailHandler(c *fiber.Ctx) error {
	var queryBody map[string]int

	err := c.BodyParser(&queryBody)
	if err != nil {
		return err
	}
	details, err := s.merchantUc.GetMerchantDetails(queryBody["merchant_id"])
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(details)
}

func (s *Svr) LocationPrediction(c *fiber.Ctx) error {
	var queryBody map[string]string

	err := c.BodyParser(&queryBody)
	if err != nil {
		return err
	}
	prediction, err := s.merchantUc.PredictPotentialMerchantLocation(queryBody["kategori"])
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(prediction)
}

func (s *Svr) ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Message

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	response := models.GeneralResponse{
		Message: message,
	}

	return c.Status(code).JSON(response)
}
