package controller

import (
	"github.com/RangelReale/osin"
	"github.com/jinzhu/gorm"
)

var DB gorm.DB
var OAuth *osin.Server
