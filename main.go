package main

import (
	"os"

	"example.com/ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	// port := viper.GetString("server.port")
	// if port != "" {
	// 	panic(r.Run(":" + port))
	// }
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configure")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
