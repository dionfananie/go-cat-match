package controller

import (
	"fmt"
	"net/http"
	"strings"
	"web/go-cat-match/database"
	"web/go-cat-match/helper"
	"web/go-cat-match/model/cat"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func RegisterCat(c *gin.Context) {
	var req cat.RegisterRequest

	userId := c.GetUint64("userId")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var CatId uint64
	var CreatedAt string
	err := database.DB.QueryRow("INSERT INTO cats (name, race, sex, ageInMonth, description, imageUrls, ownerId) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, createdat", req.Name, req.Race, req.Sex, req.AgeInMonth, req.Description, pq.Array(req.ImageUrls), userId).Scan(&CatId, &CreatedAt)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "successfully add cat", "data": gin.H{
		"id":        CatId,
		"createdAt": CreatedAt,
	}})

}

func ListCat(c *gin.Context) {
	baseQuery := "SELECT name, race, sex, ageInMonth, description, imageUrls, createdAt, hasMatched, id, ownerId from cats WHERE TRUE"
	var params []interface{}
	var conditions []string
	var limitQuery, offsetQuery string

	if id := c.Query("id"); id != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", len(params)+1))
		params = append(params, id)
	}
	if sex := c.Query("sex"); sex != "" {
		err := helper.ValidateSex(sex)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		conditions = append(conditions, fmt.Sprintf("sex = $%d", len(params)+1))
		params = append(params, sex)
	}
	if race := c.Query("race"); race != "" {
		err := helper.ValidateRace(race)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		conditions = append(conditions, fmt.Sprintf("race = $%d", len(params)+1))
		params = append(params, race)
	}
	if hasMatched := c.Query("hasMatched"); hasMatched != "" {
		conditions = append(conditions, fmt.Sprintf("hasmatched = $%d", len(params)+1))
		fmt.Println("hasMatched", hasMatched)
		params = append(params, hasMatched)
	}
	if ageInMonth := c.Query("ageInMonth"); ageInMonth != "" {
		if ageInMonth == ">4" {
			conditions = append(conditions, "ageinmonth > 4")
		} else if ageInMonth == "<4" {
			conditions = append(conditions, "ageinmonth < 4")
		} else if ageInMonth == "4" {
			conditions = append(conditions, "ageinmonth =- 4")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ageInMonth must be >4, <4 or 4"})
			return
		}
	}
	if owned := c.Query("owned"); owned != "" {
		userId := c.GetUint64("userId")
		if owned == "true" {
			conditions = append(conditions, fmt.Sprintf("ownerid = $%d", userId))
		} else {
			conditions = append(conditions, fmt.Sprintf("ownerid = $%d", userId))
		}

		fmt.Println("conditions", conditions)
	}
	if search := c.Query("search"); search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(params)+1))
		params = append(params, "%"+search+"%")
	}
	if limit := c.Query("limit"); limit != "" {
		limitQuery = fmt.Sprintf("LIMIT $%d", len(params)+1)
		params = append(params, limit)
	}
	if offset := c.Query("offset"); offset != "" {
		offsetQuery = fmt.Sprintf("OFFSET $%d", len(params)+1)
		params = append(params, offset)
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if limitQuery != "" {
		baseQuery += " " + limitQuery
	}

	if offsetQuery != "" {
		baseQuery += " " + offsetQuery
	}
	println(baseQuery)
	rows, err := database.DB.Query(baseQuery, params...)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err.Code)
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
			return
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
		cats = append(cats, catItem)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": cats})
}

func EditCat(c *gin.Context) {
	id := c.Param("id")
	var req cat.ListCat
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter required"})
	}

	baseQuery := "UPDATE cats SET "
	var params []interface{}
	var conditions []string
	if req.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(params)+1))
		params = append(params, req.Name)

	}
	if req.AgeInMonth != 0 {
		conditions = append(conditions, fmt.Sprintf("ageInMonth = $%d", len(params)+1))
		params = append(params, req.AgeInMonth)

	}
	if req.Description != "" {
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(params)+1))
		params = append(params, req.Description)

	}
	if len(req.ImageUrls) > 0 {
		conditions = append(conditions, fmt.Sprintf("imageUrls = $%d", len(params)+1))
		params = append(params, pq.Array(req.ImageUrls))

	}
	if req.Race != "" {
		if err := helper.ValidateRace(req.Race); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		conditions = append(conditions, fmt.Sprintf("race = $%d", len(params)+1))
		params = append(params, req.Race)

	}
	if req.Sex != "" {
		if err := helper.ValidateSex(req.Sex); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		conditions = append(conditions, fmt.Sprintf("sex = $%d", len(params)+1))
		params = append(params, req.Sex)
	}
	if req.HasMatched {
		conditions = append(conditions, fmt.Sprintf("hasmatched = $%d", len(params)+1))
		fmt.Println("hasMatched", req.HasMatched)
		params = append(params, req.HasMatched)
	}

	if len(conditions) > 0 {
		baseQuery = baseQuery + strings.Join(conditions, ", ")
	}

	baseQuery = baseQuery + " WHERE id = " + id

	println(baseQuery)
	res, err := database.DB.Exec(baseQuery, params...)

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

	c.JSON(http.StatusOK, gin.H{"message": "successfully update cat"})
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")

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
