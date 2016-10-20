package controllers

import (
	"net/http"
	common "github.com/Zhanat87/stack-auth/common"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	common.respondJson("All good. You don't need to be authenticated to call this 2", w)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	common.respondJson("All good. You only get this message if you're authenticated 2", w)
}
