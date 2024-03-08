package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// TemplateSets 模版套餐
type TemplateSets struct {
	ID         uint32                 `json:"id" gorm:"primaryKey"`
	Spec       *TemplateSetSpec       `json:"spec" gorm:"embedded"`
	Attachment *TemplateSetAttachment `json:"attachment" gorm:"embedded"`
	Revision   *Revision              `json:"revision" gorm:"embedded"`
}

// TemplateSetSpec defines all the specifics for TemplateSet set by user.
type TemplateSetSpec struct {
	Name        string      `json:"name" gorm:"column:name"`
	Memo        string      `json:"memo" gorm:"column:memo"`
	TemplateIDs Uint32Slice `json:"template_ids" gorm:"column:template_ids;type:json;default:'[]'"`
	Public      bool        `json:"public" gorm:"column:public"`
	BoundApps   Uint32Slice `json:"bound_apps" gorm:"column:bound_apps;type:json;default:'[]'"`
}

// TemplateSetAttachment defines the TemplateSet attachments.
type TemplateSetAttachment struct {
	BizID           uint32 `json:"biz_id" gorm:"column:biz_id"`
	TemplateSpaceID uint32 `json:"template_space_id" gorm:"column:template_space_id"`
}

type Uint32Slice []uint32

// Value implements the driver.Valuer interface
// See gorm document about customizing data types: https://gorm.io/docs/data_types.html
func (u Uint32Slice) Value() (driver.Value, error) {
	// Convert the Uint32Slice to a JSON-encoded string
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

// Scan implements the sql.Scanner interface
// See gorm document about customizing data types: https://gorm.io/docs/data_types.html
func (u *Uint32Slice) Scan(value interface{}) error {
	// Check if the value is nil
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// The value is of type []byte (MySQL driver representation for JSON columns)
		// Unmarshal the JSON-encoded value to Uint32Slice
		err := json.Unmarshal(v, u)
		if err != nil {
			return err
		}
	case string:
		// The value is of type string (fallback for older versions of MySQL driver)
		// Unmarshal the JSON-encoded value to Uint32Slice
		err := json.Unmarshal([]byte(v), u)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported Scan type for Uint32Slice")
	}

	return nil
}
