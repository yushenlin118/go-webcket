// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/gin-gonic/gin"
)

func serveHome(c *gin.Context) {
	ans := 1
	c.HTML(200, "home.html", ans)
}
func main() {
	hub := newHub()
	go hub.run()

	r := gin.Default()
	r.LoadHTMLFiles("template/home.html")
	r.Static("/assets", "./template/assets")
	r.GET("/", serveHome)
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c)
	})
	r.Run(":8080")
}
