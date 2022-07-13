package global

import (
	"time"

	"gorm.io/gorm"
)

// G_MODEL  所有结构体的通用的设置，该三个字段都是gorm中约定好的字段
type G_MODEL struct {
	ID        uint           `gorm:"primarykey; auto_increment;"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
