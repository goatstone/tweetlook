package gostart

import (
	"net/http"
	"github.com/goatstone/form/admin"
 )

func init() {
	http.HandleFunc("/admin", admin.HandleTemplate)
}
