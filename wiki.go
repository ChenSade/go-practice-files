package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gowiki/service"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// handle files, save & load

	p1 := &Page{Title: "Test Page", Body: []byte("our first test page1")}
	err := p1.save()
	if err != nil {
		return
	}

	p2, err := loadPage(p1.Title)
	if err != nil {
		return
	}
	fmt.Println(string(p2.Body))

	// handle http req/res

	album := service.Album{Title: "new ALBUM", ID: "album1", Price: 50.90}
	fmt.Println(album)

	router := gin.Default()
	router.GET("/api/v1/albums", service.GetAlbums)
	router.GET("/api/v1/albums/:id", service.GetAlbum)
	router.POST("/api/v1/albums", service.AddAlbum)
	router.DELETE("/api/v1/albums/:id", service.DeleteAlbum)
	router.PATCH("/api/v1/albums", service.UpdateAlbum)

	err = router.Run("localhost:8080")
	if err != nil {
		return
	}
}
