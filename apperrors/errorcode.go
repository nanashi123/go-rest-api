package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000" //想定外
	InsertDataFailed ErrCode = "S001" //記事の挿入
	GetDataFailed    ErrCode = "S002" //データ取得失敗
	NAData           ErrCode = "S003" //記事取得結果0件
	NiceTargetFailed ErrCode = "S004" //いいね対象記事なし
	UpdateDataFailed ErrCode = "S005" //いいね更新失敗

	ReqBodyDecodeFailed      ErrCode = "C001" //RequestBody Jsonデコード失敗
	BadParam                 ErrCode = "C002" //Request Parameterの値が不正
	ResponseBodyEncodeFailed ErrCode = "C003" //ResponseBody Jsonエンコード失敗

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidator         ErrCode = "A002"
	Unauthorizated              ErrCode = "A003"
	NotMatchUser                ErrCode = "A004"
)

type error interface {
	Error() string
}
