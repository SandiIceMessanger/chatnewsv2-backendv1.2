package controllers

import (
	"GolangMongo/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"strconv"
)

type (
	// PostController represents the controller for operating on the Post resource
	PostController struct {
		session *mgo.Session
	}
)

// NewPostController provides a reference to a PostController with provided mongo session
func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

// Get all Posts
func (uc PostController) PostsList(c *gin.Context) {

	var results []models.Post
	err := uc.session.DB("chatnews_cron").C("posts").Find(nil).All(&results)
	if err != nil {
		checkErrTypeOne(err, "Posts doesn't exist", "404", c)
		return
	}

	c.JSON(200, results)
}

// GetPost retrieves an individual post resource
func (uc PostController) GetPost(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	//Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub post
	u := models.Post{}
	err := uc.session.DB("chatnews_cron").C("posts").FindId(oid).One(&u)
	// Fetch post
	if err != nil {
		checkErrTypeTwo("Posts doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// GetPost retrieves an individual post resource
func (uc PostController) GetPostBySlug(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	//Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub post
	u := models.Post{}
	err := uc.session.DB("chatnews_cron").C("posts").FindId(oid).One(&u)
	// Fetch post
	if err != nil {
		checkErrTypeTwo("Posts doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// CreatePost creates a new post resource
func (uc PostController) CreatePost(c *gin.Context) {

	var json models.Post

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	u := uc.create_post(json.Title, json.Link, json.Pubdate, c)
	if u.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"Title":   u.Title,
			"Link":    u.Link,
			"Pubdate": u.Pubdate,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

// RemovePost removes an existing post resource
func (uc PostController) RemovePost(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove post
	if err := uc.session.DB("chatnews_cron").C("posts").RemoveId(oid); err != nil {
		checkErrTypeOne(err, "Fail to Remove", "404", c)
		return
	}

	messageTypeDefault("Success", c)

}

// RemovePost removes an existing post resource
func (uc PostController) UpdatePost(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var json models.Post

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	u := uc.update_post(id, json.Title, json.Link, json.Pubdate, c)

	if u.Title == json.Title {
		content := gin.H{
			"result": "Success",
			"Name":   u.Title,
			"Gender": u.Link,
			"Age":    u.Pubdate,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

	// Write status
	//c.AbortWithStatus(200)
}

func (uc PostController) create_post(Title string, Link string, Pubdate string, c *gin.Context) models.Post {
	post := models.Post{
		Title:   Title,
		Link:    Link,
		Pubdate: Pubdate,
	}
	// Write the post to mongo
	err := uc.session.DB("chatnews_cron").C("posts").Insert(&post)
	checkErrTypeOne(err, "Insert failed", "403", c)
	return post
}

func (uc PostController) update_post(Id string, Title string, Link string, Pubdate string, c *gin.Context) models.Post {

	post := models.Post{
		Title:   Title,
		Link:    Link,
		Pubdate: Pubdate,
	}

	oid := bson.ObjectIdHex(Id)
	// Write the post to mongo
	if err := uc.session.DB("chatnews_cron").C("posts").UpdateId(oid, &post); err != nil {
		checkErrTypeOne(err, "Update failed", "403", c)

	}

	return post
}
