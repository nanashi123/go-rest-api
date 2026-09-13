package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Nanashi123/go-api-practice/apperrors"
	"github.com/Nanashi123/go-api-practice/common"
	"google.golang.org/api/idtoken"
)

var googleClientID = os.Getenv("GOOGLE_CLIENTID")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ヘッダからAuthzationフィールドを抜き出す (2.5参照)
		authorization := req.Header.Get("Authorization")

		// Authorizationフィールドが"Bearer [IDトークン]"の形になっているか検証
		authHeaders := strings.Split(authorization, " ") //空白区切りで2つに分かれる
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1] //空白区切りで分けた1つ目がBearerで、2つ目が空ではないか
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// IDトークン検証 (12.3参照)
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err := apperrors.CannotMakeValidator.Wrap(err, "invalid auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err := apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// nameフィールドをpayloadから抜き出す (12.3参照)
		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// context にユーザー名をセット
		req = common.SetUserName(req, name.(string))

		// 本物のハンドラへ
		next.ServeHTTP(w, req)
	})
}
