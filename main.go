package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/jdkato/prose.v2"
)

type Entity struct {
	Label string `json:"label"`
	Text  string `json:"text"`
}
type Data struct {
	Text string `json:"text" binding:"required"`
}

func prose_ner(c *gin.Context) {
	var req Data

	var err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	log.Print(req.Text)
	doc, _ := prose.NewDocument(req.Text)

	var Entities []Entity

	for _, ent := range doc.Entities() {

		fmt.Println(ent.Text, ent.Label)
		Entities = append(Entities, Entity{ent.Label, ent.Text})

	}
	log.Print(Entities)
	c.JSON(200, Entities)
}

func main() {
	router := gin.Default()

	router.POST("/prose_ner/", prose_ner)

	log.Fatal(http.ListenAndServe(":8080", router))
}
