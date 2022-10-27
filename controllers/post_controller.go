package controllers

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/baguseka01/golang_fiber_blog/database"
	"github.com/baguseka01/golang_fiber_blog/models"
	"github.com/baguseka01/golang_fiber_blog/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		fmt.Println("Tidak dapat parsing body")
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Payload gagal",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Selamat, postingan anda berhasil",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1)
	var total int64
	var getPost []models.Post
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getPost)
	database.DB.Model(&models.Post{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getPost,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var detailPost models.Post
	database.DB.Where("id=?", id).Preload("User").First(&detailPost)
	return c.JSON(fiber.Map{
		"data": detailPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	post := models.Post{
		Id: uint(id),
	}

	if err := c.BodyParser(&post); err != nil {
		fmt.Println("Tidak dapat parsing body")
	}

	database.DB.Model(&post).Updates(post)
	return c.JSON(fiber.Map{
		"message": "Berhasil mengubah post anda",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var post []models.Post
	database.DB.Model(&post).Where("user_id=?", id).Preload("User").Find(&post)

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	post := models.Post{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&post)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Tidak dapat merecord",
		})

	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus postingan",
	})
}
