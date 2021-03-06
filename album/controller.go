package album

import "github.com/gin-gonic/gin"

type Controller struct {
	AlbumService Service
}

func NewAlbumController(albumService Service) Controller {
	return Controller{
		AlbumService: albumService,
	}
}

func (c *Controller) CreateAlbum(ctx *gin.Context) {
	var album DTO
	ctx.BindJSON(&album)
	err := c.AlbumService.CreateAlbum(&album)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, album)
}
func (c *Controller) GetAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	album, err := c.AlbumService.GetAlbum(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, album)
}
func (c *Controller) GetAlbums(ctx *gin.Context) {
	albums, err := c.AlbumService.GetAlbums()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(albums) == 0 {
		ctx.JSON(200, gin.H{"message": "No albums found"})
		return
	}
	ctx.JSON(200, albums)
}
func (c *Controller) UpdateAlbum(ctx *gin.Context) {
	var album Album
	ctx.BindJSON(&album)
	err := c.AlbumService.UpdateAlbum(album.Id.Hex(), &album)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, album)
}
func (c *Controller) DeleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.AlbumService.DeleteAlbum(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Album deleted"})
}
func (c *Controller) GetAlbumsByArtist(ctx *gin.Context) {
	id := ctx.Param("id")
	albums, err := c.AlbumService.GetAlbumsByArtist(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(albums) == 0 {
		ctx.JSON(200, gin.H{"message": "No albums found"})
		return
	}
	ctx.JSON(200, albums)
}
func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	albumRouter := rg.Group("/album")
	albumRouter.POST("/add", c.CreateAlbum)
	albumRouter.GET("/all", c.GetAlbums)
	albumRouter.GET(":id", c.GetAlbum)
	albumRouter.PUT("update/:id", c.UpdateAlbum)
	albumRouter.DELETE("/delete/:id", c.DeleteAlbum)
	albumRouter.GET("/artist/:id", c.GetAlbumsByArtist)
}
