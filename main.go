package main

import "gin-gonic-demo/router"

// @title API 演示 gin-gonic-demo
// @version 1.0
// @description Gin-gonic 學習筆記.
// @termsOfService https://ruisi

// @contact.name Ruisi
// @contact.url https://ruisi
// @contact.email lhc03952@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath /

func main() {
    router.Init()
}