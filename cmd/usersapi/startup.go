package usersapi

// swag init -g ./cmd/usersapi/startup.go --output ./api/usersapi
// PATH=$(go env GOPATH)/bin:$PATH
import (
	"context"
	"fmt"
	"net/http"

	usersapicmd "faceit_users/internal/usersapi"
	"faceit_users/internal/usersapi/models/config"
	"faceit_users/internal/usersapi/models/response"
	confpkg "faceit_users/pkg/config"
	mid "faceit_users/pkg/middleware"
	"faceit_users/pkg/mongo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func Execute(args []string) error {

	e := echo.New()
	var conf config.Config
	err := confpkg.LoadYMLConfig("configs/api.yaml", &conf)
	if err != nil {
		panic(err)
	}
	conf.SetMongoHost()
	client := mongo.NewMongoClient(conf.ConnectionString)
	userRep := mongo.NewUserRepository(*client, conf.DBName)
	eventRep := mongo.NewEventsRepository(*client, conf.DBName)
	eventHistroyRep := mongo.NewEventsHistoryRepository(*client, conf.DBName)
	userController := usersapicmd.NewUserController(userRep, eventRep, eventHistroyRep)

	e.Use(middleware.Recover())
	e.Use(mid.Logger())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.File("/swagger/doc.json", "./internal/usersapi/docs/swagger.json")
	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetUsers)
	e.PATCH("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
	e.GET("/hc", func(c echo.Context) error {
		err := client.Ping(context.TODO(), nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewHealthCheckResponse("FAIL", conf.DBName, conf.Port))
		}
		return c.JSON(http.StatusOK, response.NewHealthCheckResponse("OK", conf.DBName, conf.Port))
	})
	fmt.Println("api started")

	return e.Start(":" + conf.Port)
}
