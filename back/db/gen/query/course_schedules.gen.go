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

	"github.com/twin-te/twin-te/back/db/gen/model"
)

func newCourseSchedule(db *gorm.DB, opts ...gen.DOOption) courseSchedule {
	_courseSchedule := courseSchedule{}

	_courseSchedule.courseScheduleDo.UseDB(db, opts...)
	_courseSchedule.courseScheduleDo.UseModel(&model.CourseSchedule{})

	tableName := _courseSchedule.courseScheduleDo.TableName()
	_courseSchedule.ALL = field.NewAsterisk(tableName)
	_courseSchedule.ID = field.NewInt32(tableName, "id")
	_courseSchedule.Module = field.NewString(tableName, "module")
	_courseSchedule.Day = field.NewString(tableName, "day")
	_courseSchedule.Period = field.NewInt16(tableName, "period")
	_courseSchedule.Room = field.NewString(tableName, "room")
	_courseSchedule.CourseID = field.NewString(tableName, "course_id")

	_courseSchedule.fillFieldMap()

	return _courseSchedule
}

type courseSchedule struct {
	courseScheduleDo courseScheduleDo

	ALL      field.Asterisk
	ID       field.Int32
	Module   field.String
	Day      field.String
	Period   field.Int16
	Room     field.String
	CourseID field.String

	fieldMap map[string]field.Expr
}

func (c courseSchedule) Table(newTableName string) *courseSchedule {
	c.courseScheduleDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c courseSchedule) As(alias string) *courseSchedule {
	c.courseScheduleDo.DO = *(c.courseScheduleDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *courseSchedule) updateTableName(table string) *courseSchedule {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Module = field.NewString(table, "module")
	c.Day = field.NewString(table, "day")
	c.Period = field.NewInt16(table, "period")
	c.Room = field.NewString(table, "room")
	c.CourseID = field.NewString(table, "course_id")

	c.fillFieldMap()

	return c
}

func (c *courseSchedule) WithContext(ctx context.Context) *courseScheduleDo {
	return c.courseScheduleDo.WithContext(ctx)
}

func (c courseSchedule) TableName() string { return c.courseScheduleDo.TableName() }

func (c courseSchedule) Alias() string { return c.courseScheduleDo.Alias() }

func (c courseSchedule) Columns(cols ...field.Expr) gen.Columns {
	return c.courseScheduleDo.Columns(cols...)
}

func (c *courseSchedule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *courseSchedule) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["module"] = c.Module
	c.fieldMap["day"] = c.Day
	c.fieldMap["period"] = c.Period
	c.fieldMap["room"] = c.Room
	c.fieldMap["course_id"] = c.CourseID
}

func (c courseSchedule) clone(db *gorm.DB) courseSchedule {
	c.courseScheduleDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c courseSchedule) replaceDB(db *gorm.DB) courseSchedule {
	c.courseScheduleDo.ReplaceDB(db)
	return c
}

type courseScheduleDo struct{ gen.DO }

func (c courseScheduleDo) Debug() *courseScheduleDo {
	return c.withDO(c.DO.Debug())
}

func (c courseScheduleDo) WithContext(ctx context.Context) *courseScheduleDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c courseScheduleDo) ReadDB() *courseScheduleDo {
	return c.Clauses(dbresolver.Read)
}

func (c courseScheduleDo) WriteDB() *courseScheduleDo {
	return c.Clauses(dbresolver.Write)
}

func (c courseScheduleDo) Session(config *gorm.Session) *courseScheduleDo {
	return c.withDO(c.DO.Session(config))
}

func (c courseScheduleDo) Clauses(conds ...clause.Expression) *courseScheduleDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c courseScheduleDo) Returning(value interface{}, columns ...string) *courseScheduleDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c courseScheduleDo) Not(conds ...gen.Condition) *courseScheduleDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c courseScheduleDo) Or(conds ...gen.Condition) *courseScheduleDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c courseScheduleDo) Select(conds ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c courseScheduleDo) Where(conds ...gen.Condition) *courseScheduleDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c courseScheduleDo) Order(conds ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c courseScheduleDo) Distinct(cols ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c courseScheduleDo) Omit(cols ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c courseScheduleDo) Join(table schema.Tabler, on ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c courseScheduleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c courseScheduleDo) RightJoin(table schema.Tabler, on ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c courseScheduleDo) Group(cols ...field.Expr) *courseScheduleDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c courseScheduleDo) Having(conds ...gen.Condition) *courseScheduleDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c courseScheduleDo) Limit(limit int) *courseScheduleDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c courseScheduleDo) Offset(offset int) *courseScheduleDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c courseScheduleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *courseScheduleDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c courseScheduleDo) Unscoped() *courseScheduleDo {
	return c.withDO(c.DO.Unscoped())
}

func (c courseScheduleDo) Create(values ...*model.CourseSchedule) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c courseScheduleDo) CreateInBatches(values []*model.CourseSchedule, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c courseScheduleDo) Save(values ...*model.CourseSchedule) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c courseScheduleDo) First() (*model.CourseSchedule, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CourseSchedule), nil
	}
}

func (c courseScheduleDo) Take() (*model.CourseSchedule, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CourseSchedule), nil
	}
}

func (c courseScheduleDo) Last() (*model.CourseSchedule, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CourseSchedule), nil
	}
}

func (c courseScheduleDo) Find() ([]*model.CourseSchedule, error) {
	result, err := c.DO.Find()
	return result.([]*model.CourseSchedule), err
}

func (c courseScheduleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CourseSchedule, err error) {
	buf := make([]*model.CourseSchedule, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c courseScheduleDo) FindInBatches(result *[]*model.CourseSchedule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c courseScheduleDo) Attrs(attrs ...field.AssignExpr) *courseScheduleDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c courseScheduleDo) Assign(attrs ...field.AssignExpr) *courseScheduleDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c courseScheduleDo) Joins(fields ...field.RelationField) *courseScheduleDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c courseScheduleDo) Preload(fields ...field.RelationField) *courseScheduleDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c courseScheduleDo) FirstOrInit() (*model.CourseSchedule, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CourseSchedule), nil
	}
}

func (c courseScheduleDo) FirstOrCreate() (*model.CourseSchedule, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CourseSchedule), nil
	}
}

func (c courseScheduleDo) FindByPage(offset int, limit int) (result []*model.CourseSchedule, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c courseScheduleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c courseScheduleDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c courseScheduleDo) Delete(models ...*model.CourseSchedule) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *courseScheduleDo) withDO(do gen.Dao) *courseScheduleDo {
	c.DO = *do.(*gen.DO)
	return c
}
