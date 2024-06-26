package controller

import (
	"fmt"
	"net/http"
	"web/go-cat-match/database"
	"web/go-cat-match/model/cat"
	"web/go-cat-match/model/match"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func MatchCat(c *gin.Context) {
	var request match.MatchCatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetUint64("userId")

	baseQuery := fmt.Sprintf("SELECT sex, hasMatched, ownerId from cats WHERE id in (%d, %d)", request.MatchCatId, request.UserCatId)
	rows, err := database.DB.Query(baseQuery)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var catUser cat.ListCat
	var catMatcher cat.ListCat

	for rows.Next() {
		var catItem cat.ListCat
		if err := rows.Scan(&catItem.Sex, &catItem.HasMatched, &catItem.OwnerId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if uint64(catItem.OwnerId) == userId {
			catUser = catItem
		} else {
			catMatcher = catItem
		}
	}

	if uint64(catUser.OwnerId) == userId && uint64(catMatcher.OwnerId) == userId {
		c.JSON(http.StatusNotFound, gin.H{"error": "This cat is not belong to user"})
		return
	}
	if catUser.OwnerId == catMatcher.OwnerId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The cats has the same owner"})
		return
	}
	if catUser.Sex == catMatcher.Sex {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cat has same sex"})
		return
	}
	if catUser.HasMatched || catMatcher.HasMatched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cat is already matched"})
		return
	}

	var matchesId uint64
	err = database.DB.QueryRow("INSERT INTO matches (issued_user_id, user_cat_id, match_cat_id, status, message) VALUES ($1, $2, $3, $4, $5) RETURNING id", userId, request.UserCatId, request.MatchCatId, "pending", request.Message).Scan(&matchesId)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": gin.H{"message": "Successfully matching your cat!"}})
}
func ListMatch(c *gin.Context) {
	userId := c.GetUint64("userId")
	query := `SELECT m.id, m.message, m.createdAt,
				
				u.id as issued_user_id, u.name as issued_user_name, u.createdAt as issued_user_createdAt, 
				
				uc.id as user_cat_id, uc.name as user_cat_name, uc.race as user_cat_race, uc.sex as user_cat_sex,
				uc.description as user_cat_description, uc.ageinmonth as user_cat_age_in_month, uc.imageurls as user_cat_image_urls,
				
				mc.id as user_cat_id, mc.name as user_cat_name, mc.race as user_cat_race, mc.sex as user_cat_sex,
				mc.description as user_cat_description, mc.ageinmonth as user_cat_age_in_month, mc.imageurls as user_cat_image_urls
				
				FROM matches m
				JOIN users u ON m.issued_user_id = u.id
				JOIN cats uc ON m.user_cat_id = uc.id
				JOIN cats mc ON m.match_cat_id = mc.id
				WHERE m.issued_user_id = $1`
	rows, err := database.DB.Query(query, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	matchs := make([]match.MatchFullInfo, 0)
	for rows.Next() {
		var match match.MatchFullInfo
		if err := rows.Scan(
			&match.Id,
			&match.Message,
			&match.CreatedAt,

			&match.IssuedBy.Id,
			&match.IssuedBy.Name,
			&match.IssuedBy.CreatedAt,

			&match.UserCatDetail.Id,
			&match.UserCatDetail.Name,
			&match.UserCatDetail.Race,
			&match.UserCatDetail.Sex,
			&match.UserCatDetail.Description,
			&match.UserCatDetail.AgeInMonth,
			&match.UserCatDetail.ImageUrls,

			&match.MatchCatDetail.Id,
			&match.MatchCatDetail.Name,
			&match.MatchCatDetail.Race,
			&match.MatchCatDetail.Sex,
			&match.MatchCatDetail.Description,
			&match.MatchCatDetail.AgeInMonth,
			&match.MatchCatDetail.ImageUrls,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		matchs = append(matchs, match)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": matchs})

}

type ApproveMatchRequest struct {
	MatchId int `json:"matchId" binding:"required,numeric"`
}

func ApproveMatch(c *gin.Context) {
	// userId := c.GetUint64("userId")
	var req ApproveMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: VALIDATE MATCHID IS BELONGS TO THE USER
	//TODO: VALIDATE MATCHID STATUS IS PENDING

	_, err := database.DB.Exec("UPDATE matches SET status = 'approved' WHERE id = $1", req.MatchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var matchCatId uint64
	var userCatId uint64
	err = database.DB.QueryRow("SELECT match_cat_id, user_cat_id WHERE id = $1", req.MatchId).Scan(&matchCatId, &userCatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("UPDATE cats SET hasmatched = true WHERE id = $1", matchCatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("UPDATE cats SET hasmatched = true WHERE id = $1", userCatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match approved successfully"})
}

type RejectMatchRequest struct {
	MatchId int `json:"matchId" binding:"required,numeric"`
}

func RejectMatch(c *gin.Context) {
	// userId := c.GetUint64("userId")
	var req RejectMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: VALIDATE MATCHID IS BELONGS TO THE USER
	//TODO: VALIDATE MATCHID STATUS IS PENDING

	_, err := database.DB.Exec("UPDATE matches SET status = 'rejected' WHERE id = $1", req.MatchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match rejected successfully"})
}

func DeleteMatch(c *gin.Context) {
	userId := c.GetUint64("userId")
	matchId := c.Param("id")
	var issuedUserId uint64

	query := fmt.Sprintf("SELECT issued_user_id FROM matches WHERE id = %s", matchId)
	err := database.DB.QueryRow(query).Scan(&issuedUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if issuedUserId != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not the creator of this matching"})
		return
	}
	rows, err := database.DB.Exec("DELETE FROM matches WHERE id = $1", matchId)
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

	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully"})
}
