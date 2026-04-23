package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/TaisukeFujise/blog_api/api/common"
	"github.com/TaisukeFujise/blog_api/apperrors"
)

type userNameKey struct{}

const (
	googleClientID = "390957602026-e7nhsbm36gpjsv92f8nrd4c2mf1847k7.apps.googleusercontent.com"
)

// コンテキストからnameフィールドの値を取り出す関数
// func GetUserName(ctx context.Context) string {
// 	id := ctx.Value(userNameKey{})

// 	if usernameStr, ok := id.(string); ok {
// 		return usernameStr
// 	}
// 	return ""
// }

// コンテキストにnameフィールドをセットする関数
// func SetUserName(req *http.Request, name string) *http.Request {
// 	ctx := req.Context()

// 	ctx = context.WithValue(ctx, userNameKey{}, name)
// 	req = req.WithContext(ctx)

// 	return req
// }

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダからAuthorizationフィールド抜き出し
		authorization := req.Header.Get("Authorization")

		// Authorizationフィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// Bearerで、2つ目がからではないか
		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// IDトークン検証
		tokenValidator, err := idtoken.NewValidator(nil)
		if err != nil {
			err := apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// nameフィールドをpayloadから抜き出す
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
