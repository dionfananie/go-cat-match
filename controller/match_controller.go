package controller

import (
	"fmt"
	"net/http"
	"web/go-cat-match/database"
	"web/go-cat-match/model/match"

	"github.com/gin-gonic/gin"
)

func MatchCat(c *gin.Context) {
	var request match.MatchCatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: VALIDATE THERE IS A CAT WITH ID MATCHCATID AND USERCATID
	//TODO: VALIDATE USERCATID IS BELONGS TO THE USER

	userId := c.GetUint64("userId")
	matchStatus := "pending"
	_, err := database.DB.Exec("INSERT INTO matches (match_cat_id, user_cat_id, issued_user_id, message, status) VALUES ($1, $2, $3, $4)", request.MatchCatId, request.UserCatId, userId, request.Message, matchStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request)

	c.JSON(http.StatusCreated, gin.H{"message": "Matched successfully"})
}

func ListMatch(c *gin.Context) {
	userId := c.GetUint64("userId")
	rows, err := database.DB.Query("SELECT id, match_cat_id, user_cat_id, message, created_at from matches WHERE issued_user_id = $1", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var matchs []match.Match
	for rows.Next() {
		var match match.Match
		if err := rows.Scan(&match.Id, &match.MatchCatId, &match.UserCatId, &match.Message, &match.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		matchs = append(matchs, match)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//TODO: ADD ISSUEDUSER, CATMATCHDETAIL, AND CATUSERDETAIL TO THE RESPONSE

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
