package groups

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	d "github.com/zigamedved/user-management-api/defaults"
	h "github.com/zigamedved/user-management-api/helpers"
	m "github.com/zigamedved/user-management-api/models"
)

type GroupsAPI struct {
}

// Get /groups
// Get all groups
func (api *GroupsAPI) GroupsGet(c *gin.Context) {

	var groups []m.Group
	err := d.GetAll[m.Group](m.Groups, c, &groups)
	if err != nil {
		log.Printf("error while getting groups")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// Get /groups/:id
// Get a group by ID
func (api *GroupsAPI) GroupsIdGet(c *gin.Context) {

	var group m.Group
	err := d.GetSingle[m.Group](m.Groups, c, &group)
	if err != nil {
		log.Printf("error while getting requested group")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if group.ID.IsZero() {
		c.JSON(http.StatusNotFound, gin.H{"message": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// Post /groups
// Add a new group
func (api *GroupsAPI) GroupsPost(c *gin.Context) {

	var document m.Group
	if err1 := h.RequestToStruct(c.Request, &document); err1 != nil {
		log.Printf("error while parsing POST group request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
		return
	}

	if err2 := h.ValidateStruct(document); err2 != nil {
		log.Printf("error while validating POST group request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}

	ID, err3 := d.Create[m.Group](m.Groups, c, document)
	if err3 != nil {
		log.Printf("error while creating group")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err3.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": ID})
}

// Put /groups/:id
// Update a group by ID
func (api *GroupsAPI) GroupsIdPut(c *gin.Context) {

	var document m.Group
	if err1 := h.RequestToStruct(c.Request, &document); err1 != nil {
		log.Printf("error while parsing PUT group request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
		return
	}

	if err2 := h.ValidateStruct(document); err2 != nil {
		log.Printf("error while validating PUT group request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}

	err3 := d.Update[m.Group](m.Groups, c, document)
	if err3 != nil {
		log.Printf("error while updating group")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err3.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Delete /groups/:id
// Delete a group by ID
func (api *GroupsAPI) GroupsIdDelete(c *gin.Context) {

	err := d.Delete[m.Group](m.Groups, c)
	if err != nil {
		log.Printf("error while deleting group")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
