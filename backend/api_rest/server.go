package apirest

import (
	"log"
	"net/http"
	"os"

	"github.com/JuanMillan7818/FUT21/api_rest/controllers"
	"github.com/JuanMillan7818/FUT21/rest_fut21/util"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	environment()
	util.Start()

	instanceEcho := echo.New()

	instanceEcho.Use(middleware.Logger())
	instanceEcho.Use(middleware.Recover())
	instanceEcho.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	instanceEcho.Use(verifyKeyApi)

	controllers.EndPointMain(instanceEcho)

	instanceEcho.Logger.Fatal(instanceEcho.Start(":8000"))
}

func verifyKeyApi(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		keyRequest := c.Request().Header.Get("x-api-key")
		keyResponse := os.Getenv("X_API_KEY")
		if keyRequest == keyResponse {
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "No tiene la autorizacion necesaria")
		}
	}
}

func environment() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
}
