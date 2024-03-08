package model

import "time"

// Revision is a resource's status information
type Revision struct {
	Creator   string    `db:"creator" json:"creator" gorm:"column:creator"`
	Reviser   string    `db:"reviser" json:"reviser" gorm:"column:reviser"`
	CreatedAt time.Time `db:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" gorm:"column:updated_at"`
}
