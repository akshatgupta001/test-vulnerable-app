package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/miekg/dns"
	"github.com/opencontainers/runc/libcontainer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v2"
)

func main() {
	// JWT
	token := jwt.New(jwt.SigningMethodHS256)
	_ = token

	// Gin
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "SCS test app running",
		})
	})

	// Protobuf
	_ = proto.Marshal

	// Gorilla Mux
	router := mux.NewRouter()
	_ = router

	// DNS
	msg := new(dns.Msg)
	_ = msg

	// Runc
	_ = libcontainer.Cgroupfs

	// Logrus
	logrus.Info("Initialized")

	// Viper
	viper.SetDefault("test", "value")

	// GJSON
	result := gjson.Get(`{"key":"value"}`, "key")
	_ = result

	// YAML
	data := make(map[string]string)
	_ = yaml.Unmarshal([]byte("key: value"), &data)

	fmt.Println("SCS test dependencies loaded")
	r.Run(":8080")
}
