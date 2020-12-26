package model

// QBittorrent qBittorrent服务器
type QBittorrent struct {
	User     int64
	URL      string
	Username string
	Password string
}

// CreateQb 添加qb服务
func CreateQb(qb QBittorrent) error {
	return db.Create(&qb).Error
}

// FineQb 查找qb服务
func FineQb(user int64) QBittorrent {
	var qb QBittorrent
	db.Where(&QBittorrent{
		User: user,
	}).First(&qb)
	return qb
}

// UpdateQb 更新qb服务
func UpdateQb(qb QBittorrent) error {
	return db.Model(&qb).Updates(qb).Error
}
