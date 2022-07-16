package artist

import (
	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	ArtistService ArtistService
}

func NewArtistController(userService ArtistService) ArtistController {
	return ArtistController{
		ArtistService: userService,
	}
}

func (c *ArtistController) GetArtists(ctx *gin.Context) {
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
func (c *ArtistController) GetArtist(ctx *gin.Context) {
	id := ctx.Param("id")
	artist, err := c.ArtistService.GetArtist(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}

func (c *ArtistController) CreateArtist(ctx *gin.Context) {
	var artist ArtistDTO
	ctx.BindJSON(&artist)
	err := c.ArtistService.CreateArtist(&artist)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}
func (c *ArtistController) UpdateArtist(ctx *gin.Context) {
	var artist Artist
	ctx.BindJSON(&artist)
	err := c.ArtistService.UpdateArtist(artist.Id.Hex(), &artist)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, artist)
}
func (c *ArtistController) DeleteArtist(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.ArtistService.DeleteArtist(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Artist deleted"})
}
func (c *ArtistController) RegisterRoutes(rg *gin.RouterGroup) {
	artistRoute := rg.Group("/artist")
	artistRoute.GET("/all", c.GetArtists)
	artistRoute.GET(":id", c.GetArtist)
	artistRoute.POST("/add", c.CreateArtist)
	artistRoute.PUT("/update/:id", c.UpdateArtist)
	artistRoute.DELETE("/delete/:id", c.DeleteArtist)
}
