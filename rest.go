package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	srv := gin.Default()
	srv.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	srv.GET("/album", GetAlbums)
	srv.GET("/album/:id", GetAlbumsByID)
	srv.PUT("/album", CreateAlbum)

	return srv
}

func GetAlbums(ctx *gin.Context) {
	albums, err := SelectAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured", "stack": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, albums)
}

func GetAlbumsByID(ctx *gin.Context) {
	id, paramErr := strconv.Atoi(ctx.Param("id"))
	if paramErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Path parameter :id error", "stack": paramErr.Error()})
		return
	}
	album, err := AlbumByID(int64(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "An error occured", "stack": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, album)
}

func CreateAlbum(ctx *gin.Context) {
	var createdAlbum Album
	// TODO: Add body content validation
	if err := ctx.BindJSON(&createdAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to create an album", "stack": err.Error()})
		return
	}

	albID, err := AddAlbum(Album{
		Title:  createdAlbum.Title,
		Artist: createdAlbum.Artist,
		Price:  createdAlbum.Price,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to create an album", "stack": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": albID})
}

