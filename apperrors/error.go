package apperrors

type MyAppError struct {
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  `json: "-"` // エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	if myErr == nil {
		return "<nil>"
	}
	if myErr.Err != nil {
		return myErr.Err.Error()
	}
	return myErr.Message
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
