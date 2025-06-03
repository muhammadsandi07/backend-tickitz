package middlewares

import (
	"backendtickitz/pkg"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// func (m *Middlewares) AccessGateAdmin(ctx *gin.Context) {
// 	// ambil payload/claims dari context gin
// 	claims, exist := ctx.Get("Payload")
// 	if !exist {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"message": "Silahkan login terlebih dahulu",
// 		})
// 		return
// 	}
// 	// type assertion claims menjadi pkg.claims
// 	userClaims, ok := claims.(*pkg.Claims)
// 	// log.Println(userClaims)
// 	if !ok {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{

// 			"message": "Identitas login anda rusak, Silahkan login kembali",
// 		})
// 		return
// 	}
// 	// cek role yang ada di claims
// 	if userClaims.Role != "admin" {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"message": "anda tidak dapat mengakses sumber ini",
// 		})
// 		return
// 	}
// 	ctx.Next()
// }

func (m *Middlewares) AccessGate(allowedRole ...string) func(*gin.Context) {
	// ambil payload/claims dari context gin
	return func(ctx *gin.Context) {
		claims, exist := ctx.Get("Payload")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Silahkan login terlebih dahulu",
			})
			return
		}
		// type assertion claims menjadi pkg.claims
		userClaims, ok := claims.(*pkg.Claims)
		// log.Println(userClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "silahkan login dahulu",
			})
			return

		}
		// cek role yang ada di claims
		if !slices.Contains(allowedRole, userClaims.Role) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "anda tidak dapat mengakses sumber ini",
			})
			return
		}
		ctx.Next()
	}
}
