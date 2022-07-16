package artist

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ArtistService Service
}

func NewArtistController(userService Service) Controller {
	return Controller{
		ArtistService: userService,
	}
}

func (c *Controller) GetArtists(ctx *gin.Context) {
	artists, err := c.ArtistService.GetArtists()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(artists) == 0 {
		ctx.JSON(200, gin.H{"message": "No artists found"})
		return
	}
	ctx.JSON(200, artists)
}
func (c *Controller) GetArtist(ctx *gin.Context) {
	id := ctx.Param("id")
	artist, err := c.ArtistService.GetArtist(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}

func (c *Controller) CreateArtist(ctx *gin.Context) {
	var artist DTO
	ctx.BindJSON(&artist)
	err := c.ArtistService.CreateArtist(&artist)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}
func (c *Controller) UpdateArtist(ctx *gin.Context) {
	var artist Artist
	ctx.BindJSON(&artist)
	err := c.ArtistService.UpdateArtist(artist.Id.Hex(), &artist)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}
func (c *Controller) DeleteArtist(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.ArtistService.DeleteArtist(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Artist deleted"})
}
func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	artistRoute := rg.Group("/artist")
	artistRoute.GET("/all", c.GetArtists)
	artistRoute.GET(":id", c.GetArtist)
	artistRoute.POST("/add", c.CreateArtist)
	artistRoute.PUT("/update/:id", c.UpdateArtist)
	artistRoute.DELETE("/delete/:id", c.DeleteArtist)
}
