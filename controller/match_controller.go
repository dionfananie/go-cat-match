package controller

import (
	"fmt"
	"net/http"
	"web/go-cat-match/database"
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

	//TODO: VALIDATE THERE IS A CAT WITH ID MATCHCATID AND USERCATID
	//TODO: VALIDATE USERCATID IS BELONGS TO THE USER

	var OwnerId uint64
	var Sex string
	var HasMatched bool
	var OwnerIdUser uint64
	var SexUser string
	var HasMatchedUser bool
	userId := c.GetUint64("userId")
	// findout cat target
	baseQuery := fmt.Sprintf("SELECT sex, hasMatched, ownerId from cats WHERE id = %v", request.MatchCatId)
	println("query: ", baseQuery)
	err := database.DB.QueryRow(baseQuery).Scan(&Sex, &HasMatched, &OwnerId)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// findout cat user
	baseQuery = fmt.Sprintf("SELECT sex, hasMatched, ownerId from cats WHERE id = %v", request.UserCatId)
	err = database.DB.QueryRow(baseQuery).Scan(&SexUser, &HasMatchedUser, &OwnerIdUser)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if userId != OwnerIdUser {
		c.JSON(http.StatusNotFound, gin.H{"error": "This cat is not belong to user"})
		return
	}
	if OwnerId == OwnerIdUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The cats has the same owner"})
		return
	}
	if Sex == SexUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cat has same sex"})
		return
	}
	if HasMatched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cat is already matched"})
		return
	}

	baseQuery = fmt.Sprintf("UPDATE cats SET hasMatched = true WHERE id IN (%d , %d)", request.MatchCatId, request.UserCatId)
	res, err := database.DB.Exec(baseQuery)
	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Detail})
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusCreated, gin.H{"message": gin.H{"message": "Successfully matching your cat!"}})
}

func ListMatch(c *gin.Context) {
	userId := c.GetUint64("userId")
	query := `SELECT m.id, m.message, m.created_at,
				
				u.id as issued_user_id, u.name as issued_user_name, u.created_at as issued_user_created_at, 
				
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

	var matchs []match.MatchFullInfo
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
	// userId := c.GetUint64("userId")
	matchId := c.Param("id")

	//TODO: VALIDATE MATCHID IS BELONGS TO THE USER
	//TODO: VALIDATE MATCHID STATUS IS PENDING

	_, err := database.DB.Exec("DELETE FROM matches WHERE id = $1", matchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully"})
}
