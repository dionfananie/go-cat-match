package controller

import (
	"fmt"
	"net/http"
	"web/go-cat-match/database"
	"web/go-cat-match/helper"
	"web/go-cat-match/model/cat"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func RegisterCat(c *gin.Context) {
	var req cat.RegisterRequest

	// debugging
	// &userId string := c.GetString("userId")
	// if !exists {
	// 	return
	// }
	// println("UserId", &userId)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println(req.Name)
	// var CatId uint64
	// var CreatedAt string
	// err := database.DB.QueryRow("INSERT INTO cats (name, race, sex, ageInMonth, description, imageUrls) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at", req.Name, req.Race, req.Sex, req.AgeInMonth, req.Description, pq.Array(req.ImageUrls)).Scan(&CatId, &CreatedAt)
	// if err != nil {

	// 	if err, ok := err.(*pq.Error); ok {
	// 		fmt.Println("pq error:", err.Code)
	// 		c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
	// 	}
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"message": "successfully add cat", "data": gin.H{
	// 	"id":         CatId,
	// 	"created_at": CreatedAt,
	// }})

}

func ListCat(c *gin.Context) {

	rows, err := database.DB.Query("SELECT name, race, sex, ageInMonth, description, imageUrls, created_at, hasMatched, id, ownerId from cats")

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var cats []cat.ListCat
	for rows.Next() {
		var catItem cat.ListCat
		if err := rows.Scan(&catItem.Name, &catItem.Race, &catItem.Sex, &catItem.AgeInMonth, &catItem.Description, &catItem.ImageUrls, &catItem.CreatedAt,
			&catItem.HasMatched, &catItem.Id, &catItem.OwnerId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		println(&catItem)
		cats = append(cats, catItem)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": cats})
}
func EditCat(c *gin.Context) {
	id := c.Param("id")
	var req cat.ListCat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// query := `UPDATE cats SET name = $1`
	if err := helper.ValidateSex(req.Sex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := helper.ValidateRace(req.Race); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := database.DB.Exec("UPDATE cats SET name = $1, sex = $2, race = $3, description = $4, ageInMonth = $5, imageUrls = $6 WHERE id = $7", req.Name, req.Sex, req.Race, req.Description, req.AgeInMonth, req.ImageUrls, id)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " successfully update cat"})
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")
	println("id", id)

	rows, err := database.DB.Exec("DELETE FROM cats WHERE id = $1", id)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": " successfully delete cat"})

}
