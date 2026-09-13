package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	// (フィールド名を省略した場合。型名がそのままフィールド名になる)
	ErrCode //レスポンスとログに表示するエラーコード

	// string型のMessageフィールド
	Message string //レスポンスに表示するエラーメッセージ

	Err error `json:"-"` //エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// 返り値として、入れ子にして含んでいる内部エラーを返す
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
