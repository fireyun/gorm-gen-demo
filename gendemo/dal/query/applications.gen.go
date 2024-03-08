// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fireyun/gorm-gen-demo/gendemo/dal/model"
)

func newApp(db *gorm.DB, opts ...gen.DOOption) app {
	_app := app{}

	_app.appDo.UseDB(db, opts...)
	_app.appDo.UseModel(&model.App{})

	tableName := _app.appDo.TableName()
	_app.ALL = field.NewAsterisk(tableName)
	_app.ID = field.NewUint32(tableName, "id")
	_app.BizID = field.NewUint32(tableName, "biz_id")
	_app.Name = field.NewString(tableName, "name")
	_app.ConfigType = field.NewString(tableName, "config_type")
	_app.Mode = field.NewString(tableName, "mode")
	_app.Memo = field.NewString(tableName, "memo")
	_app.ReloadType = field.NewString(tableName, "reload_type")
	_app.ReloadFilePath = field.NewString(tableName, "reload_file_path")
	_app.Alias_ = field.NewString(tableName, "alias")
	_app.DataType = field.NewString(tableName, "data_type")
	_app.Creator = field.NewString(tableName, "creator")
	_app.Reviser = field.NewString(tableName, "reviser")
	_app.CreatedAt = field.NewTime(tableName, "created_at")
	_app.UpdatedAt = field.NewTime(tableName, "updated_at")

	_app.fillFieldMap()

	return _app
}

type app struct {
	appDo appDo

	ALL            field.Asterisk
	ID             field.Uint32
	BizID          field.Uint32
	Name           field.String
	ConfigType     field.String
	Mode           field.String
	Memo           field.String
	ReloadType     field.String
	ReloadFilePath field.String
	Alias_         field.String
	DataType       field.String
	Creator        field.String
	Reviser        field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time

	fieldMap map[string]field.Expr
}

func (a app) Table(newTableName string) *app {
	a.appDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a app) As(alias string) *app {
	a.appDo.DO = *(a.appDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *app) updateTableName(table string) *app {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.BizID = field.NewUint32(table, "biz_id")
	a.Name = field.NewString(table, "name")
	a.ConfigType = field.NewString(table, "config_type")
	a.Mode = field.NewString(table, "mode")
	a.Memo = field.NewString(table, "memo")
	a.ReloadType = field.NewString(table, "reload_type")
	a.ReloadFilePath = field.NewString(table, "reload_file_path")
	a.Alias_ = field.NewString(table, "alias")
	a.DataType = field.NewString(table, "data_type")
	a.Creator = field.NewString(table, "creator")
	a.Reviser = field.NewString(table, "reviser")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *app) WithContext(ctx context.Context) *appDo { return a.appDo.WithContext(ctx) }

func (a app) TableName() string { return a.appDo.TableName() }

func (a app) Alias() string { return a.appDo.Alias() }

func (a *app) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *app) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.ID
	a.fieldMap["biz_id"] = a.BizID
	a.fieldMap["name"] = a.Name
	a.fieldMap["config_type"] = a.ConfigType
	a.fieldMap["mode"] = a.Mode
	a.fieldMap["memo"] = a.Memo
	a.fieldMap["reload_type"] = a.ReloadType
	a.fieldMap["reload_file_path"] = a.ReloadFilePath
	a.fieldMap["alias"] = a.Alias_
	a.fieldMap["data_type"] = a.DataType
	a.fieldMap["creator"] = a.Creator
	a.fieldMap["reviser"] = a.Reviser
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a app) clone(db *gorm.DB) app {
	a.appDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a app) replaceDB(db *gorm.DB) app {
	a.appDo.ReplaceDB(db)
	return a
}

type appDo struct{ gen.DO }

func (a appDo) Debug() *appDo {
	return a.withDO(a.DO.Debug())
}

func (a appDo) WithContext(ctx context.Context) *appDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appDo) ReadDB() *appDo {
	return a.Clauses(dbresolver.Read)
}

func (a appDo) WriteDB() *appDo {
	return a.Clauses(dbresolver.Write)
}

func (a appDo) Session(config *gorm.Session) *appDo {
	return a.withDO(a.DO.Session(config))
}

func (a appDo) Clauses(conds ...clause.Expression) *appDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appDo) Returning(value interface{}, columns ...string) *appDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appDo) Not(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appDo) Or(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appDo) Select(conds ...field.Expr) *appDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appDo) Where(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *appDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a appDo) Order(conds ...field.Expr) *appDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appDo) Distinct(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appDo) Omit(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appDo) Join(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appDo) LeftJoin(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appDo) RightJoin(table schema.Tabler, on ...field.Expr) *appDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appDo) Group(cols ...field.Expr) *appDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appDo) Having(conds ...gen.Condition) *appDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appDo) Limit(limit int) *appDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appDo) Offset(offset int) *appDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *appDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appDo) Unscoped() *appDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appDo) Create(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appDo) CreateInBatches(values []*model.App, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appDo) Save(values ...*model.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appDo) First() (*model.App, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Take() (*model.App, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Last() (*model.App, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) Find() ([]*model.App, error) {
	result, err := a.DO.Find()
	return result.([]*model.App), err
}

func (a appDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.App, err error) {
	buf := make([]*model.App, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appDo) FindInBatches(result *[]*model.App, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appDo) Attrs(attrs ...field.AssignExpr) *appDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appDo) Assign(attrs ...field.AssignExpr) *appDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appDo) Joins(fields ...field.RelationField) *appDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appDo) Preload(fields ...field.RelationField) *appDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appDo) FirstOrInit() (*model.App, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FirstOrCreate() (*model.App, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.App), nil
	}
}

func (a appDo) FindByPage(offset int, limit int) (result []*model.App, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appDo) Delete(models ...*model.App) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appDo) withDO(do gen.Dao) *appDo {
	a.DO = *do.(*gen.DO)
	return a
}