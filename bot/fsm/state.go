package fsm

// UserStatus UserStatus
type UserStatus int

const (
	// None None
	None UserStatus = iota
	// AddQbURL 添加qb服务器地址
	AddQbURL
	// AddQbUser 添加qb服务器用户名
	AddQbUser
	// AddQbPass 添加qb服务器密码
	AddQbPass
	// ChangeQbURL 修改qb服务器url
	ChangeQbURL
	// ChangeQbUser 修改qb服务器用户名
	ChangeQbUser
	// ChangeQbPass 修改qb服务器密码
	ChangeQbPass
)

// ChangeQbURLBtn 修改qb服务器url
const ChangeQbURLBtn = "ChangeQbURLBtn"

// ChangeQbUserBtn 修改qb服务器用户名
const ChangeQbUserBtn = "ChangeQbUserBtn"

// ChangeQbPassBtn 修改qb服务器密码
const ChangeQbPassBtn = "ChangeQbPassBtn"
