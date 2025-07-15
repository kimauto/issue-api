package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateIssue(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		UserID      *uint  `json:"userId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, "필수 필드 누락")
		return
	}

	issue, err := CreateNewIssue(req.Title, req.Description, req.UserID)
	if err != nil {
		respondError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, issue)
}

func GetIssues(c *gin.Context) {
	statusFilter := c.Query("status")
	result := []Issue{}

	for _, issue := range issues {
		if statusFilter == "" || issue.Status == statusFilter {
			result = append(result, *issue)
		}
	}
	c.JSON(http.StatusOK, gin.H{"issues": result})
}

func GetIssueByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, http.StatusBadRequest, "유효하지 않은 ID입니다")
		return
	}
	issue, exists := issues[uint(id)]
	if !exists {
		respondError(c, http.StatusNotFound, "이슈를 찾을 수 없습니다")
		return
	}
	c.JSON(http.StatusOK, issue)
}

func UpdateIssue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, http.StatusBadRequest, "유효하지 않은 ID입니다")
		return
	}
	issue, exists := issues[uint(id)]
	if !exists {
		respondError(c, http.StatusNotFound, "이슈가 존재하지 않습니다")
		return
	}

	if issue.Status == "COMPLETED" || issue.Status == "CANCELLED" {
		respondError(c, http.StatusBadRequest, "완료되거나 취소된 이슈는 수정할 수 없습니다")
		return
	}

	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
		UserID      *uint   `json:"userId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, "잘못된 요청 형식")
		return
	}

	if req.UserID != nil {
		if *req.UserID == 0 {
			issue.User = nil
			issue.Status = "PENDING"
		} else {
			user, ok := users[*req.UserID]
			if !ok {
				respondError(c, http.StatusBadRequest, "지정한 담당자가 존재하지 않습니다")
				return
			}
			issue.User = user
			if issue.Status == "PENDING" && req.Status == nil {
				issue.Status = "IN_PROGRESS"
			}
		}
	}

	if req.Status != nil {
		if (req.Status == strPtr("IN_PROGRESS") || req.Status == strPtr("COMPLETED")) && issue.User == nil {
			respondError(c, http.StatusBadRequest, "담당자 없이 해당 상태로 변경할 수 없습니다")
			return
		}
		issue.Status = *req.Status
	}
	if req.Title != nil {
		issue.Title = *req.Title
	}
	if req.Description != nil {
		issue.Description = *req.Description
	}
	issue.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, issue)
}

func strPtr(s string) *string {
	return &s
}
