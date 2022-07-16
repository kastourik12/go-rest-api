package song

import "github.com/gin-gonic/gin"

type Controller struct {
	SongService Service
}

func NewSongController(songService Service) Controller {
	return Controller{
		SongService: songService,
	}
}
func (c *Controller) CreateSong(ctx *gin.Context) {
	var song DTO
	ctx.BindJSON(&song)
	err := c.SongService.CreateSong(&song)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, song)
}
func (c *Controller) GetSong(ctx *gin.Context) {
	id := ctx.Param("id")
	song, err := c.SongService.GetSong(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, song)
}
func (c *Controller) GetSongs(ctx *gin.Context) {
	songs, err := c.SongService.GetSongs()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(songs) == 0 {
		ctx.JSON(404, gin.H{"message": "No songs found"})
		return
	}
	ctx.JSON(200, songs)
}
func (c *Controller) UpdateSong(ctx *gin.Context) {
	var song Song
	ctx.BindJSON(&song)
	err := c.SongService.UpdateSong(song.Id.Hex(), &song)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, song)
}
func (c *Controller) DeleteSong(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.SongService.DeleteSong(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Song deleted"})
}
func (c *Controller) GetSongsByAlbum(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	songs, err := c.SongService.GetSongsByAlbum(albumId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(songs) == 0 {
		ctx.JSON(200, gin.H{"message": "No songs found"})
		return
	}
	ctx.JSON(200, songs)
}
func (c *Controller) GetSongsByArtist(ctx *gin.Context) {
	artistId := ctx.Param("artistId")
	songs, err := c.SongService.GetSongsByArtist(artistId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(songs) == 0 {
		ctx.JSON(200, gin.H{"message": "No songs found"})
		return
	}
	ctx.JSON(200, songs)
}
func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	songRoute := rg.Group("/song")
	songRoute.POST("/add", c.CreateSong)
	songRoute.GET("/:id", c.GetSong)
	songRoute.GET("/all", c.GetSongs)
	songRoute.PUT("/update/:id", c.UpdateSong)
	songRoute.DELETE("/delete/:id", c.DeleteSong)
	songRoute.GET("/album/:albumId", c.GetSongsByAlbum)
	songRoute.GET("/artist/:artistId", c.GetSongsByArtist)

}
