package errcode

const (
	//通用
	SUCCESS = 0
	ERROR   = -1
)

var ErrMsg = map[int]string{
	SUCCESS: "success",
	ERROR:   "系统错误",
}
