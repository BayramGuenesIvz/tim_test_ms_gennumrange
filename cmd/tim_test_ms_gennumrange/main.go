package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//boot "mdh.koeln.ivz.cn.ard.de/bitbucket/scm/mdhpres/tim_test_ms_gennumrange.git/cmd"
	//web "mdh.koeln.ivz.cn.ard.de/bitbucket/scm/mdhpres/tim_test_ms_gennumrange.git/web"

	"github.com/BayramGuenesIvz/tim_test_ms_gennumrange/internal/data"
	config "github.com/BayramGuenesIvz/tim_test_ms_gennumrange/internal/readConfig"
	settings "github.com/BayramGuenesIvz/tim_test_ms_gennumrange/internal/readSettings"
	"github.com/BayramGuenesIvz/tim_test_ms_gennumrange/pkg/monitor"
)

var ()

func main() {
	println()
	routerGinGonic()
}

func routerGinGonic() {
	settings.LoadExternalSettings()
	_, err := config.GetApplConf()
	if err != nil {
		return
	}

	router := gin.New()

	router.GET("/", amAliveHandler)

	// -----------------------------------------------------------------//
	router.LoadHTMLGlob("web/templates/*")
	//router.Static("/static", "./public")
	//router.Static("/static", "web/public")
	router.Static("/static", "web/static")

	router.GET("/NumRangeServices", monitor.GotoMonitor)
	router.POST("/performOp", monitor.ServicesPerformOp)
	// -----------------------------------------------------------------//

	log.Print("Starting Server on Port", ":"+data.ThisPort) //setImageServiceRoutes(router)

	http.ListenAndServe(":"+data.ThisPort, router)
}
func amAliveHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "tim_test_ms_gennumrange (micro-)service is alive")
}
