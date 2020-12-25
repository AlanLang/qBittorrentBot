package fsm

// UserStatus UserStatus
type UserStatus int

const (
	// None None
	None UserStatus = iota
	// AddQbName 添加qb服务名
	AddQbName
	// AddQbPath 添加qb服务器地址
	AddQbPath
)
