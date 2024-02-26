package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	d "github.com/zigamedved/user-management-api/defaults"
	h "github.com/zigamedved/user-management-api/helpers"
	m "github.com/zigamedved/user-management-api/models"
)

type UsersAPI struct {
}

// Get /users
// Get all users
func (api *UsersAPI) UsersGet(c *gin.Context) {

	var users []m.User
	err := d.GetAll[m.User](m.Users, c, &users)
	if err != nil {
		log.Printf("error while getting users")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Get /users/:id
// Get a user by ID
func (api *UsersAPI) UsersIdGet(c *gin.Context) {

	var user m.User
	err := d.GetSingle[m.User](m.Users, c, &user)
	if err != nil {
		log.Printf("error while getting requested user")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if user.ID.IsZero() {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Post /users
// Add a new user
func (api *UsersAPI) UsersPost(c *gin.Context) {

	var document m.User
	if err1 := h.RequestToStruct(c.Request, &document); err1 != nil {
		log.Printf("error while parsing POST user request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
		return
	}

	if err2 := h.ValidateStruct(document); err2 != nil {
		log.Printf("error while validating POST user request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}

	ID, err3 := d.Create[m.User](m.Users, c, document)
	if err3 != nil {
		log.Printf("error while creating user")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err3.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"_id": ID})
}

// Put /users/:id
// Update a user by ID
func (api *UsersAPI) UsersIdPut(c *gin.Context) {

	var document m.User
	if err1 := h.RequestToStruct(c.Request, &document); err1 != nil {
		log.Printf("error while parsing PUT user request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
		return
	}

	if err2 := h.ValidateStruct(document); err2 != nil {
		log.Printf("error while validating PUT user request")
		c.JSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
		return
	}

	err3 := d.Update[m.User](m.Users, c, document)
	if err3 != nil {
		log.Printf("error while updating user")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err3.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Delete /users/:id
// Delete a user by ID
func (api *UsersAPI) UsersIdDelete(c *gin.Context) {

	err := d.Delete[m.User](m.Users, c)
	if err != nil {
		log.Printf("error while deleting user")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
