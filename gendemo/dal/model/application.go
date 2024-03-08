package model

import (
	"errors"
	"fmt"
	"time"
)

// App defines an application's detail information
type App struct {
	// ID is an auto-increased value, which is an application's
	// unique identity.
	ID uint32 `json:"id" gorm:"primaryKey"`
	// BizID is the business is which this app belongs to
	BizID uint32 `json:"biz_id" gorm:"column:biz_id"`
	// Spec is a collection of app's specifics defined with user
	Spec *AppSpec `json:"spec" gorm:"embedded"`
	// Revision record this app's revision information
	Revision *Revision `json:"revision" gorm:"embedded"`
}

// TableName is the app's database table name.
func (a *App) TableName() string {
	return "applications"
}

// AppID AuditRes interface
func (a *App) AppID() uint32 {
	return a.ID
}

// ResID AuditRes interface
func (a *App) ResID() uint32 {
	return a.ID
}

// ResType AuditRes interface
func (a *App) ResType() string {
	return "app"
}

// AppSpec is a collection of app's specifics defined with user
type AppSpec struct {
	// Name is application's name
	Name string `json:"name" gorm:"column:name"`
	// ConfigType defines which type is this configuration, different type has the
	// different ways to be consumed.
	ConfigType ConfigType `json:"config_type" gorm:"column:config_type"`
	// Mode defines what mode of this app works at.
	// Mode can not be updated once it is created.
	Mode     AppMode  `json:"mode" gorm:"column:mode"`
	Memo     string   `json:"memo" gorm:"column:memo"`
	Reload   *Reload  `json:"reload" gorm:"embedded"`
	Alias    string   `json:"alias" gorm:"alias"`
	DataType DataType `json:"data_type" gorm:"data_type"`
}

const (
	// Normal means this is a normal app, and configuration
	// items can be consumed directly.
	Normal AppMode = "normal"

	// Namespace means that this app runs in the namespace
	// mode, which means user must consume app's configuration
	// item with namespace information.
	Namespace AppMode = "namespace"
)

// AppMode is the mode of an app works at, different mode has the
// different way or restricts to consume this strategy's configurations.
type AppMode string

// Validate strategy set type.
func (s AppMode) Validate() error {
	switch s {
	case Normal:
	case Namespace:
	default:
		return fmt.Errorf("unsupported app working mode: %s", s)
	}

	return nil
}

// Reload is a collection of app reload specifics defined with user. only is used when this app is file config type.
// Reload is used to control how bscp sidecar notifies applications to go to reload config files.
type Reload struct {
	ReloadType     AppReloadType   `json:"reload_type" gorm:"column:reload_type"`
	FileReloadSpec *FileReloadSpec `json:"file_reload_spec" gorm:"embedded"`
}

// IsEmpty reload.
func (r *Reload) IsEmpty() error {
	if r == nil {
		return nil
	}

	if len(r.ReloadType) != 0 {
		return errors.New("reload type is not nil")
	}

	if r.FileReloadSpec != nil {
		if err := r.FileReloadSpec.IsEmpty(); err != nil {
			return err
		}
	}

	return nil
}

// FileReloadSpec is a collection of file reload spec's specifics defined with user.
type FileReloadSpec struct {
	ReloadFilePath string `json:"reload_file_path" gorm:"column:reload_file_path"`
}

// IsEmpty file reload spec.
func (f *FileReloadSpec) IsEmpty() error {
	if f == nil {
		return nil
	}

	if len(f.ReloadFilePath) != 0 {
		return errors.New("reload file path is not nil")
	}

	return nil
}

const (
	// KV is kv configuration type
	KV ConfigType = "kv"
	// File is file configuration type
	File ConfigType = "file"
	// Table is table configuration type
	Table ConfigType = "table"
)

// ConfigType is the app's config item's type
type ConfigType string

// Validate the config type is supported or not.
func (c ConfigType) Validate() error {
	switch c {
	case KV:
	case File:
	case Table:
		return errors.New("not support table config type for now")
	default:
		return fmt.Errorf("unsupported config type: %s", c)
	}

	return nil
}

const (
	// ReloadWithFile the app's sidecar instance will write the downloaded configuration release information to the
	// reload file, then the application instance uses this reload file to determine whether has a new configuration
	// need to load.
	ReloadWithFile AppReloadType = "file"
)

// AppReloadType is the app's sidecar instance to notify application reload config files way.
type AppReloadType string

// Validate app reload type
func (rt AppReloadType) Validate() error {
	switch rt {
	case ReloadWithFile:
	default:
		return fmt.Errorf("unsupported app reload type: %s", rt)
	}

	return nil
}

// ArchivedApp is used to record applications basic information
// which is used to purge resources related with this application
// asynchronously.
type ArchivedApp struct {
	ID        uint32    `json:"id" gorm:"primaryKey"`
	BizID     uint32    `json:"biz_id" gorm:"column:biz_id"`
	AppID     uint32    `json:"app_id" gorm:"column:app_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

// TableName is the archived app's database table name.
func (a *ArchivedApp) TableName() string {
	return "archived_apps"
}

// DataType is the app's kv type
type DataType string

const (
	// KvAny 任意类型
	KvAny DataType = "any"
	// KvStr is the type for string kv
	KvStr DataType = "string"
	// KvNumber is the type for number kv
	KvNumber DataType = "number"
	// KvText is the type for text kv
	KvText DataType = "text"
	// KvJson is the type for json kv
	KvJson DataType = "json"
	// KvYAML is the type for yaml kv
	KvYAML DataType = "yaml"
	// KvXml is the type for xml kv
	KvXml DataType = "xml"
)

// ValidateApp the kvType and value match
func (k DataType) ValidateApp() error {
	switch k {
	case KvAny:
	case KvStr:
	case KvNumber:
	case KvText:
	case KvJson:
	case KvYAML:
	case KvXml:
	default:
		return errors.New("invalid data-type")
	}
	return nil
}
