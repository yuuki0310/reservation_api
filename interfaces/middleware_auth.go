package interfaces

// import (
// 	"crypto/rsa"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"math/big"
// 	"net/http"
// 	"strings"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/yuuki0310/reservation_api/utils"
// )

// var (
// 	jwksCache = make(map[string]interface{}) // 公開鍵のキャッシュ
// 	mu        sync.Mutex                     // 排他制御用
// )

// // JWTMiddleware handles JWT authentication
// func jwtMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Authorizationヘッダーからトークンを取得
// 		authHeader := c.GetHeader("Authorization")
// 		if !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
// 			c.Abort()
// 			return
// 		}
// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 		// JWTを解析して検証
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// トークンの署名アルゴリズムを確認
// 			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}

// 			// Cognitoの公開鍵を取得
// 			kid := token.Header["kid"].(string)
// 			return getPublicKeyFromJWKS(kid)
// 		})

// 		if err != nil || !token.Valid {
// 			fmt.Println("err", err)
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		// JWTクレームをコンテキストに設定
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: invalid claims"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }

// // getPublicKeyFromJWKS retrieves the public key from the JWKS URL
// func getPublicKeyFromJWKS(kid string) (interface{}, error) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	// キャッシュから取得
// 	if key, found := jwksCache[kid]; found {
// 		return key, nil
// 	}

// 	// JWKSを取得
// 	resp, err := http.Get(utils.Conf.Cognito.JWKSURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read JWKS: %w", err)
// 	}

// 	var jwks struct {
// 		Keys []struct {
// 			Kid string `json:"kid"`
// 			N   string `json:"n"`
// 			E   string `json:"e"`
// 		} `json:"keys"`
// 	}

// 	if err := json.Unmarshal(body, &jwks); err != nil {
// 		return nil, fmt.Errorf("failed to parse JWKS: %w", err)
// 	}

// 	// 必要なキーをキャッシュ
// 	for _, key := range jwks.Keys {
// 		if key.Kid == kid {
// 			// モジュラスと指数をデコード
// 			nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to decode modulus: %w", err)
// 			}
// 			eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to decode exponent: %w", err)
// 			}

// 			// Exponentを整数に変換
// 			e := 0
// 			for _, b := range eBytes {
// 				e = e<<8 + int(b)
// 			}
// 			// 公開鍵を構築
// 			pubKey := &rsa.PublicKey{
// 				N: new(big.Int).SetBytes(nBytes),
// 				E: e,
// 			}
// 			jwksCache[kid] = pubKey
// 			return pubKey, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("key not found in JWKS")
// }
