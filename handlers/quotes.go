package handlers

import (
	"time"
	"math/rand"
	"github.com/MrBlackBlade/qotd/database"
	"github.com/MrBlackBlade/qotd/models"
	"github.com/gofiber/fiber/v2"
)

func CalcQOTD() (qotd models.Quote) {

	currentTime := time.Now()

	qotdWindow := currentTime.Add(-24 * time.Hour)
	bltime := currentTime.Add(-24 * 5 * time.Hour)

	quotes := []models.Quote{}
	database.DB.Db.Where("lastqotd > ?", qotdWindow).Find(&quotes)
	if (len(quotes) != 0) {
		qotd = quotes[0]
	} else {
		database.DB.Db.Where("lastqotd < ?", bltime).Find(&quotes)
		qotd = quotes[rand.Intn(len(quotes))]
		database.DB.Db.Model(&qotd).Update("lastqotd", currentTime)
	}

	return qotd
}

func ListQuotes(c *fiber.Ctx) error {
	quotes := []models.Quote{}
	database.DB.Db.Find(&quotes)

	return c.Status(200).JSON(quotes)
}

func QuoteOfTheDay(c *fiber.Ctx) error {
	return c.Status(200).JSON(CalcQOTD())
}

/*func CreateQuote(c *fiber.Ctx) error {
	quote := new(models.Quote)
	if err := c.BodyParser(quote); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&quote)

	return c.Status(200).JSON(quote)
}*/

func CreateQuotes(c *fiber.Ctx) error {
	quotes := new([]models.Quote)
	if err := c.BodyParser(quotes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, quote := range *quotes {
		database.DB.Db.Create(&quote)
	}
	return c.Status(200).JSON(quotes)
}