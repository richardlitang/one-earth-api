package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/heisenberglar/one-earth-api/database"
	"github.com/heisenberglar/one-earth-api/models"
)

func handleBadRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func GetQuests(c *fiber.Ctx) error {
	q := []models.Quest{}
	database.DB.Db.Find(&q)

	return c.Status(200).JSON(q)
}

func GetQuest(c *fiber.Ctx) error {
	id := c.Params("id")
	qid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	q := models.Quest{}
	err = database.DB.Db.Debug().Model(models.Quest{}).Where("id = ?", qid).Take(&q).Error
	if err != nil {
		return handleBadRequest(c, err)
	}

	return c.Status(200).JSON(q)
}

func CreateQuest(c *fiber.Ctx) error {
	quest := new(models.Quest)

	if err := c.BodyParser(quest); err != nil {
		return handleBadRequest(c, err)
	}

	database.DB.Db.Create(&quest)

	return c.Status(http.StatusOK).JSON(quest)
}

func UpdateQuest(c *fiber.Ctx) error {
	id := c.Params("id")
	qid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	oq := models.Quest{}
	err = database.DB.Db.Debug().Model(models.Quest{}).Where("id = ?", qid).Take(&oq).Error
	if err != nil {
		return handleBadRequest(c, err)
	}

	q := new(models.Quest)
	if err := c.BodyParser(q); err != nil {
		return handleBadRequest(c, err)
	}

	err = database.DB.Db.Debug().Where("id = ?", qid).Updates(&q).Error
	if err != nil {
		return handleBadRequest(c, err)
	}

	return c.Status(http.StatusOK).JSON(q)
}

func DeleteQuest(c *fiber.Ctx) error {
	id := c.Params("id")
	qid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return handleBadRequest(c, err)
	}

	oq := models.Quest{}
	err = database.DB.Db.Debug().Model(models.Quest{}).Where("id = ?", qid).Take(&oq).Error
	if err != nil {
		return handleBadRequest(c, err)
	}

	err = database.DB.Db.Debug().Where("id = ?", qid).Take(&oq).Delete(&oq).Error
	if err != nil {
		return handleBadRequest(c, err)
	}

	return c.Status(http.StatusOK).JSON(oq)
}
