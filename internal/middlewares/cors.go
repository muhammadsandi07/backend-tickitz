package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) CORSMiddleware(ctx *gin.Context) {
	// setup whitelist origin
	whitelistOrigin := []string{"http://localhost:5173"}
	origin := ctx.GetHeader("Origin")

	if slices.Contains(whitelistOrigin, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, content-type")
	// handle preflight
	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()

}
