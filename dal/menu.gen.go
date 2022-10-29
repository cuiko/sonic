// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"github.com/go-sonic/sonic/model/entity"
)

func newMenu(db *gorm.DB) menu {
	_menu := menu{}

	_menu.menuDo.UseDB(db)
	_menu.menuDo.UseModel(&entity.Menu{})

	tableName := _menu.menuDo.TableName()
	_menu.ALL = field.NewAsterisk(tableName)
	_menu.ID = field.NewInt32(tableName, "id")
	_menu.CreateTime = field.NewTime(tableName, "create_time")
	_menu.UpdateTime = field.NewTime(tableName, "update_time")
	_menu.Icon = field.NewString(tableName, "icon")
	_menu.Name = field.NewString(tableName, "name")
	_menu.ParentID = field.NewInt32(tableName, "parent_id")
	_menu.Priority = field.NewInt32(tableName, "priority")
	_menu.Target = field.NewString(tableName, "target")
	_menu.Team = field.NewString(tableName, "team")
	_menu.URL = field.NewString(tableName, "url")

	_menu.fillFieldMap()

	return _menu
}

type menu struct {
	menuDo menuDo

	ALL        field.Asterisk
	ID         field.Int32
	CreateTime field.Time
	UpdateTime field.Time
	Icon       field.String
	Name       field.String
	ParentID   field.Int32
	Priority   field.Int32
	Target     field.String
	Team       field.String
	URL        field.String

	fieldMap map[string]field.Expr
}

func (m menu) Table(newTableName string) *menu {
	m.menuDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m menu) As(alias string) *menu {
	m.menuDo.DO = *(m.menuDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *menu) updateTableName(table string) *menu {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")
	m.Icon = field.NewString(table, "icon")
	m.Name = field.NewString(table, "name")
	m.ParentID = field.NewInt32(table, "parent_id")
	m.Priority = field.NewInt32(table, "priority")
	m.Target = field.NewString(table, "target")
	m.Team = field.NewString(table, "team")
	m.URL = field.NewString(table, "url")

	m.fillFieldMap()

	return m
}

func (m *menu) WithContext(ctx context.Context) *menuDo { return m.menuDo.WithContext(ctx) }

func (m menu) TableName() string { return m.menuDo.TableName() }

func (m menu) Alias() string { return m.menuDo.Alias() }

func (m *menu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *menu) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["id"] = m.ID
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["name"] = m.Name
	m.fieldMap["parent_id"] = m.ParentID
	m.fieldMap["priority"] = m.Priority
	m.fieldMap["target"] = m.Target
	m.fieldMap["team"] = m.Team
	m.fieldMap["url"] = m.URL
}

func (m menu) clone(db *gorm.DB) menu {
	m.menuDo.ReplaceDB(db)
	return m
}

type menuDo struct{ gen.DO }

func (m menuDo) Debug() *menuDo {
	return m.withDO(m.DO.Debug())
}

func (m menuDo) WithContext(ctx context.Context) *menuDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m menuDo) ReadDB() *menuDo {
	return m.Clauses(dbresolver.Read)
}

func (m menuDo) WriteDB() *menuDo {
	return m.Clauses(dbresolver.Write)
}

func (m menuDo) Clauses(conds ...clause.Expression) *menuDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m menuDo) Returning(value interface{}, columns ...string) *menuDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m menuDo) Not(conds ...gen.Condition) *menuDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m menuDo) Or(conds ...gen.Condition) *menuDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m menuDo) Select(conds ...field.Expr) *menuDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m menuDo) Where(conds ...gen.Condition) *menuDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m menuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *menuDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m menuDo) Order(conds ...field.Expr) *menuDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m menuDo) Distinct(cols ...field.Expr) *menuDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m menuDo) Omit(cols ...field.Expr) *menuDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m menuDo) Join(table schema.Tabler, on ...field.Expr) *menuDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m menuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *menuDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m menuDo) RightJoin(table schema.Tabler, on ...field.Expr) *menuDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m menuDo) Group(cols ...field.Expr) *menuDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m menuDo) Having(conds ...gen.Condition) *menuDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m menuDo) Limit(limit int) *menuDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m menuDo) Offset(offset int) *menuDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m menuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *menuDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m menuDo) Unscoped() *menuDo {
	return m.withDO(m.DO.Unscoped())
}

func (m menuDo) Create(values ...*entity.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m menuDo) CreateInBatches(values []*entity.Menu, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m menuDo) Save(values ...*entity.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m menuDo) First() (*entity.Menu, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menu), nil
	}
}

func (m menuDo) Take() (*entity.Menu, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menu), nil
	}
}

func (m menuDo) Last() (*entity.Menu, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menu), nil
	}
}

func (m menuDo) Find() ([]*entity.Menu, error) {
	result, err := m.DO.Find()
	return result.([]*entity.Menu), err
}

func (m menuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Menu, err error) {
	buf := make([]*entity.Menu, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m menuDo) FindInBatches(result *[]*entity.Menu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m menuDo) Attrs(attrs ...field.AssignExpr) *menuDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m menuDo) Assign(attrs ...field.AssignExpr) *menuDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m menuDo) Joins(fields ...field.RelationField) *menuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m menuDo) Preload(fields ...field.RelationField) *menuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m menuDo) FirstOrInit() (*entity.Menu, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menu), nil
	}
}

func (m menuDo) FirstOrCreate() (*entity.Menu, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menu), nil
	}
}

func (m menuDo) FindByPage(offset int, limit int) (result []*entity.Menu, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m menuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m menuDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m menuDo) Delete(models ...*entity.Menu) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *menuDo) withDO(do gen.Dao) *menuDo {
	m.DO = *do.(*gen.DO)
	return m
}