package middlewares

import (
	"backendtickitz/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middlewares struct{}

func InitMiddleware() *Middlewares {
	return &Middlewares{}

}

func (m *Middlewares) VerifyToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login terlebih dahulu",
		})
		return
	}

	// verifikasi bearer token
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login terlebih dahulu",
		})
		return
	}
	// pisahkan token dari bearer

	token := strings.Split(bearerToken, " ")[1]
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "silahkan login terlebih dahulu",
		})
		return
	}

	// verifikasi
	claims := &pkg.Claims{}
	log.Println("ini claims 1 verifikasi", claims)

	if err := claims.VerifyToken(token); err != nil {
		log.Println("ini error verifikasi", err)

		if strings.Contains(err.Error(), "expired") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Sesi anda berakhir, Silahkan login kembali",
			})
			return
		}
		if strings.Contains(err.Error(), "malformed") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Identitas login anda rusak, Silahkan login kembali",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "terjadi kesalah server",
		})
		return
	}
	log.Println(claims)
	ctx.Set("Payload", claims)
	ctx.Next()
}
