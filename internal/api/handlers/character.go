package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
)

// ListCharacters 角色列表
func ListCharacters(c *gin.Context) {
	limit := 20
	offset := 0

	if l, ok := c.GetQuery("limit"); ok {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		}
	}
	if o, ok := c.GetQuery("offset"); ok {
		if val, err := strconv.Atoi(o); err == nil {
			offset = val
		}
	}

	characters, err := repository.NewCharacterRepository().List(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"characters": characters,
		"limit":      limit,
		"offset":     offset,
	})
}

// GetCharacter 角色详情
func GetCharacter(c *gin.Context) {
	id := c.Param("id")

	character, err := repository.NewCharacterRepository().FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, character)
}

// CreateCharacter 创建角色
func CreateCharacter(c *gin.Context) {
	userID := c.GetString("userId")

	var req struct {
		Name         string   `json:"name" binding:"required"`
		Age          int      `json:"age"`
		Avatar       string   `json:"avatar"`
		Banner       string   `json:"banner"`
		Bio          string   `json:"bio"`
		Personality  []string  `json:"personality"`
		Relationship string   `json:"relationship"`
		Profession   string   `json:"profession"`
		Interests    []string  `json:"interests"`
		Voice        string   `json:"voice"`
		Style        string   `json:"style"`
		Gender       string   `json:"gender"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character := &models.Character{
		UserID:       userID,
		Name:         req.Name,
		Age:          req.Age,
		Avatar:       req.Avatar,
		Banner:       req.Banner,
		Bio:          req.Bio,
		Personality:  req.Personality,
		Relationship: req.Relationship,
		Profession:   req.Profession,
		Interests:    req.Interests,
		Voice:        req.Voice,
		Style:        req.Style,
		Gender:       req.Gender,
		IsOfficial:   false,
	}

	if err := repository.NewCharacterRepository().Create(character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, character)
}

// UpdateCharacter 更新角色
func UpdateCharacter(c *gin.Context) {
	userID := c.GetString("userId")
	id := c.Param("id")

	character, err := repository.NewCharacterRepository().FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	// 验证所有权
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	var req struct {
		Name         *string  `json:"name"`
		Age          *int     `json:"age"`
		Avatar       *string  `json:"avatar"`
		Banner       *string  `json:"banner"`
		Bio          *string  `json:"bio"`
		Personality  []string `json:"personality"`
		Relationship *string  `json:"relationship"`
		Profession   *string  `json:"profession"`
		Interests    []string `json:"interests"`
		Voice        *string  `json:"voice"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新非空字段
	if req.Name != nil {
		character.Name = *req.Name
	}
	if req.Age != nil {
		character.Age = *req.Age
	}
	if req.Avatar != nil {
		character.Avatar = *req.Avatar
	}
	if req.Banner != nil {
		character.Banner = *req.Banner
	}
	if req.Bio != nil {
		character.Bio = *req.Bio
	}
	if req.Personality != nil {
		character.Personality = req.Personality
	}
	if req.Relationship != nil {
		character.Relationship = *req.Relationship
	}
	if req.Profession != nil {
		character.Profession = *req.Profession
	}
	if req.Interests != nil {
		character.Interests = req.Interests
	}
	if req.Voice != nil {
		character.Voice = *req.Voice
	}

	if err := repository.NewCharacterRepository().Update(character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}

// DeleteCharacter 删除角色
func DeleteCharacter(c *gin.Context) {
	userID := c.GetString("userId")
	id := c.Param("id")

	character, err := repository.NewCharacterRepository().FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	// 验证所有权
	if character.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	if err := repository.NewCharacterRepository().Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character deleted successfully"})
}

// GetMyCharacters 获取我的角色列表
func GetMyCharacters(c *gin.Context) {
	userID := c.GetString("userId")

	characters, err := repository.NewCharacterRepository().FindByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"characters": characters,
	})
}
