package items

import (
	"net/http"
	"time"

	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	itemsStorage = map[uuid.UUID]*item{}
)

func createItemController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var newItem newItem
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := saveNewItem(claims, &newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func getAllItemsController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	onlyItems := []*item{}
	for _, v := range itemsStorage {
		if v.UserID == claims.ID || claims.IsAdmin {
			onlyItems = append(onlyItems, v)
		}
	}
	c.JSON(http.StatusOK, onlyItems)
}

func getItemController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, exists := itemsStorage[id]
	if exists && item.UserID == claims.ID || claims.IsAdmin {
		c.JSON(http.StatusOK, item)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func deleteItemController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rawID := c.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, exists := itemsStorage[id]
	if exists && (item.UserID == claims.ID || claims.IsAdmin) {
		delete(itemsStorage, id)
		now := time.Now()
		item.DeletedAt = &now
		c.JSON(http.StatusOK, item)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func saveNewItem(claims *auth.Claims, newItem *newItem) (*item, error) {
	item := item{
		ID:          uuid.New(),
		UserID:      claims.ID,
		Name:        newItem.Name,
		Description: newItem.Description,
		CreatedAt:   time.Now(),
	}
	itemsStorage[item.ID] = &item
	return &item, nil
}
