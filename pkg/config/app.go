package config 


import (

	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"


)


var db *gorm.DB


func connect (){
	d , err := gorm.Open("mysql" , "")
}