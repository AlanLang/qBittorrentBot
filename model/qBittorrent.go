package model

// QBittorrent qBittorrent服务器
type QBittorrent struct {
	Name     string `gorm:"primary_key;"`
	Path     string
	Username string
	Password string
}

// GetAllQb 获取所有qb服务
func GetAllQb() ([]QBittorrent, error) {
	var qbs []QBittorrent

	db.Find(&qbs)

	return qbs, nil
}

// CreateQb 添加qb服务
func CreateQb(qb QBittorrent) error {
	return db.Create(&qb).Error
}

// FineQb 查找qb服务
func FineQb(name string) QBittorrent {
	var qb QBittorrent
	db.Where(&QBittorrent{Name: name}).First(&qb)
	return qb
}
