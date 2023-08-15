package handlers

import (
	"github.com/Garry028/trivia/database"
	"github.com/Garry028/trivia/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}
func ShowFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := new(models.Fact)
	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	return c.Status(200).JSON(fact)

}
func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := new(models.Fact)
	err := database.DB.Db.First(&fact, id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	database.DB.Db.Save(&fact)
	return c.Status(200).JSON(fact)
}


func DeleteFact(c *fiber.Ctx) error {
	id :=c.Params("id")
	fact:=new(models.Fact)
	err:=database.DB.Db.First(&fact,id).Error
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message":err.Error(),
			})
	}
	database.DB.Db.Delete(&fact)
	return c.Status(200).JSON(fiber.Map{
		"message":"Fact deleted successfully",
	})
}
