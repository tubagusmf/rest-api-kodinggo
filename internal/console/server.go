package console

import (
	"golang-rest-api-articles/db"
	"golang-rest-api-articles/internal/config"
	handlerHttp "golang-rest-api-articles/internal/delivery/http"
	"golang-rest-api-articles/internal/repository"
	"golang-rest-api-articles/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start the HTTP server",
	Long:  `Start server`,
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	config.LoadWithViper()

	db := db.NewMysql()
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)
	userRepo := repository.NewUserRepository(db)

	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	//Create a new echo instance
	e := echo.New()

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewArticleHandler(routeGroup, articleUsecase)
	handlerHttp.NewUserHandler(routeGroup, userUsecase)

	e.Logger.Fatal(e.Start(":3200"))

}
