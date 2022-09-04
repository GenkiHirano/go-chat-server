package chat

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Run() {
	r := gin.Default()
	m := melody.New()

	r.Static("/static", "./view/static")
	r.LoadHTMLGlob("view/*.html")

	// GinでGet実装予定
}
