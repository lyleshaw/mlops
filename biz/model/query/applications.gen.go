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

	model "github.com/lyleshaw/mlops/biz/model/orm_gen"
)

func newApplication(db *gorm.DB, opts ...gen.DOOption) application {
	_application := application{}

	_application.applicationDo.UseDB(db, opts...)
	_application.applicationDo.UseModel(&model.Application{})

	tableName := _application.applicationDo.TableName()
	_application.ALL = field.NewAsterisk(tableName)
	_application.AppID = field.NewInt32(tableName, "app_id")
	_application.AppName = field.NewString(tableName, "app_name")
	_application.AppStatus = field.NewString(tableName, "app_status")
	_application.AppPort = field.NewInt32(tableName, "app_port")
	_application.AppImage = field.NewString(tableName, "app_image")
	_application.DockerID = field.NewString(tableName, "docker_id")
	_application.UserID = field.NewInt32(tableName, "user_id")
	_application.IsDeleted = field.NewBool(tableName, "is_deleted")
	_application.CreateAt = field.NewTime(tableName, "create_at")
	_application.UpdateAt = field.NewTime(tableName, "update_at")

	_application.fillFieldMap()

	return _application
}

type application struct {
	applicationDo

	ALL       field.Asterisk
	AppID     field.Int32
	AppName   field.String
	AppStatus field.String
	AppPort   field.Int32
	AppImage  field.String
	DockerID  field.String
	UserID    field.Int32
	IsDeleted field.Bool
	CreateAt  field.Time
	UpdateAt  field.Time

	fieldMap map[string]field.Expr
}

func (a application) Table(newTableName string) *application {
	a.applicationDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a application) As(alias string) *application {
	a.applicationDo.DO = *(a.applicationDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *application) updateTableName(table string) *application {
	a.ALL = field.NewAsterisk(table)
	a.AppID = field.NewInt32(table, "app_id")
	a.AppName = field.NewString(table, "app_name")
	a.AppStatus = field.NewString(table, "app_status")
	a.AppPort = field.NewInt32(table, "app_port")
	a.AppImage = field.NewString(table, "app_image")
	a.DockerID = field.NewString(table, "docker_id")
	a.UserID = field.NewInt32(table, "user_id")
	a.IsDeleted = field.NewBool(table, "is_deleted")
	a.CreateAt = field.NewTime(table, "create_at")
	a.UpdateAt = field.NewTime(table, "update_at")

	a.fillFieldMap()

	return a
}

func (a *application) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *application) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["app_name"] = a.AppName
	a.fieldMap["app_status"] = a.AppStatus
	a.fieldMap["app_port"] = a.AppPort
	a.fieldMap["app_image"] = a.AppImage
	a.fieldMap["docker_id"] = a.DockerID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["is_deleted"] = a.IsDeleted
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
}

func (a application) clone(db *gorm.DB) application {
	a.applicationDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a application) replaceDB(db *gorm.DB) application {
	a.applicationDo.ReplaceDB(db)
	return a
}

type applicationDo struct{ gen.DO }

type IApplicationDo interface {
	gen.SubQuery
	Debug() IApplicationDo
	WithContext(ctx context.Context) IApplicationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IApplicationDo
	WriteDB() IApplicationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IApplicationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IApplicationDo
	Not(conds ...gen.Condition) IApplicationDo
	Or(conds ...gen.Condition) IApplicationDo
	Select(conds ...field.Expr) IApplicationDo
	Where(conds ...gen.Condition) IApplicationDo
	Order(conds ...field.Expr) IApplicationDo
	Distinct(cols ...field.Expr) IApplicationDo
	Omit(cols ...field.Expr) IApplicationDo
	Join(table schema.Tabler, on ...field.Expr) IApplicationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IApplicationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IApplicationDo
	Group(cols ...field.Expr) IApplicationDo
	Having(conds ...gen.Condition) IApplicationDo
	Limit(limit int) IApplicationDo
	Offset(offset int) IApplicationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IApplicationDo
	Unscoped() IApplicationDo
	Create(values ...*model.Application) error
	CreateInBatches(values []*model.Application, batchSize int) error
	Save(values ...*model.Application) error
	First() (*model.Application, error)
	Take() (*model.Application, error)
	Last() (*model.Application, error)
	Find() ([]*model.Application, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Application, err error)
	FindInBatches(result *[]*model.Application, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Application) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IApplicationDo
	Assign(attrs ...field.AssignExpr) IApplicationDo
	Joins(fields ...field.RelationField) IApplicationDo
	Preload(fields ...field.RelationField) IApplicationDo
	FirstOrInit() (*model.Application, error)
	FirstOrCreate() (*model.Application, error)
	FindByPage(offset int, limit int) (result []*model.Application, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IApplicationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a applicationDo) Debug() IApplicationDo {
	return a.withDO(a.DO.Debug())
}

func (a applicationDo) WithContext(ctx context.Context) IApplicationDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a applicationDo) ReadDB() IApplicationDo {
	return a.Clauses(dbresolver.Read)
}

func (a applicationDo) WriteDB() IApplicationDo {
	return a.Clauses(dbresolver.Write)
}

func (a applicationDo) Session(config *gorm.Session) IApplicationDo {
	return a.withDO(a.DO.Session(config))
}

func (a applicationDo) Clauses(conds ...clause.Expression) IApplicationDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a applicationDo) Returning(value interface{}, columns ...string) IApplicationDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a applicationDo) Not(conds ...gen.Condition) IApplicationDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a applicationDo) Or(conds ...gen.Condition) IApplicationDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a applicationDo) Select(conds ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a applicationDo) Where(conds ...gen.Condition) IApplicationDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a applicationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IApplicationDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a applicationDo) Order(conds ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a applicationDo) Distinct(cols ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a applicationDo) Omit(cols ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a applicationDo) Join(table schema.Tabler, on ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a applicationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a applicationDo) RightJoin(table schema.Tabler, on ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a applicationDo) Group(cols ...field.Expr) IApplicationDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a applicationDo) Having(conds ...gen.Condition) IApplicationDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a applicationDo) Limit(limit int) IApplicationDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a applicationDo) Offset(offset int) IApplicationDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a applicationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IApplicationDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a applicationDo) Unscoped() IApplicationDo {
	return a.withDO(a.DO.Unscoped())
}

func (a applicationDo) Create(values ...*model.Application) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a applicationDo) CreateInBatches(values []*model.Application, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a applicationDo) Save(values ...*model.Application) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a applicationDo) First() (*model.Application, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Application), nil
	}
}

func (a applicationDo) Take() (*model.Application, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Application), nil
	}
}

func (a applicationDo) Last() (*model.Application, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Application), nil
	}
}

func (a applicationDo) Find() ([]*model.Application, error) {
	result, err := a.DO.Find()
	return result.([]*model.Application), err
}

func (a applicationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Application, err error) {
	buf := make([]*model.Application, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a applicationDo) FindInBatches(result *[]*model.Application, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a applicationDo) Attrs(attrs ...field.AssignExpr) IApplicationDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a applicationDo) Assign(attrs ...field.AssignExpr) IApplicationDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a applicationDo) Joins(fields ...field.RelationField) IApplicationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a applicationDo) Preload(fields ...field.RelationField) IApplicationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a applicationDo) FirstOrInit() (*model.Application, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Application), nil
	}
}

func (a applicationDo) FirstOrCreate() (*model.Application, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Application), nil
	}
}

func (a applicationDo) FindByPage(offset int, limit int) (result []*model.Application, count int64, err error) {
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

func (a applicationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a applicationDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a applicationDo) Delete(models ...*model.Application) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *applicationDo) withDO(do gen.Dao) *applicationDo {
	a.DO = *do.(*gen.DO)
	return a
}
