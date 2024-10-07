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

func newTimetableDay(db *gorm.DB, opts ...gen.DOOption) timetableDay {
	_timetableDay := timetableDay{}

	_timetableDay.timetableDayDo.UseDB(db, opts...)
	_timetableDay.timetableDayDo.UseModel(&model.TimetableDay{})

	tableName := _timetableDay.timetableDayDo.TableName()
	_timetableDay.ALL = field.NewAsterisk(tableName)
	_timetableDay.Day = field.NewString(tableName, "day")

	_timetableDay.fillFieldMap()

	return _timetableDay
}

type timetableDay struct {
	timetableDayDo timetableDayDo

	ALL field.Asterisk
	Day field.String

	fieldMap map[string]field.Expr
}

func (t timetableDay) Table(newTableName string) *timetableDay {
	t.timetableDayDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timetableDay) As(alias string) *timetableDay {
	t.timetableDayDo.DO = *(t.timetableDayDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timetableDay) updateTableName(table string) *timetableDay {
	t.ALL = field.NewAsterisk(table)
	t.Day = field.NewString(table, "day")

	t.fillFieldMap()

	return t
}

func (t *timetableDay) WithContext(ctx context.Context) *timetableDayDo {
	return t.timetableDayDo.WithContext(ctx)
}

func (t timetableDay) TableName() string { return t.timetableDayDo.TableName() }

func (t timetableDay) Alias() string { return t.timetableDayDo.Alias() }

func (t timetableDay) Columns(cols ...field.Expr) gen.Columns {
	return t.timetableDayDo.Columns(cols...)
}

func (t *timetableDay) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timetableDay) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 1)
	t.fieldMap["day"] = t.Day
}

func (t timetableDay) clone(db *gorm.DB) timetableDay {
	t.timetableDayDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timetableDay) replaceDB(db *gorm.DB) timetableDay {
	t.timetableDayDo.ReplaceDB(db)
	return t
}

type timetableDayDo struct{ gen.DO }

func (t timetableDayDo) Debug() *timetableDayDo {
	return t.withDO(t.DO.Debug())
}

func (t timetableDayDo) WithContext(ctx context.Context) *timetableDayDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timetableDayDo) ReadDB() *timetableDayDo {
	return t.Clauses(dbresolver.Read)
}

func (t timetableDayDo) WriteDB() *timetableDayDo {
	return t.Clauses(dbresolver.Write)
}

func (t timetableDayDo) Session(config *gorm.Session) *timetableDayDo {
	return t.withDO(t.DO.Session(config))
}

func (t timetableDayDo) Clauses(conds ...clause.Expression) *timetableDayDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timetableDayDo) Returning(value interface{}, columns ...string) *timetableDayDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timetableDayDo) Not(conds ...gen.Condition) *timetableDayDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timetableDayDo) Or(conds ...gen.Condition) *timetableDayDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timetableDayDo) Select(conds ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timetableDayDo) Where(conds ...gen.Condition) *timetableDayDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timetableDayDo) Order(conds ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timetableDayDo) Distinct(cols ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timetableDayDo) Omit(cols ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timetableDayDo) Join(table schema.Tabler, on ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timetableDayDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timetableDayDo) RightJoin(table schema.Tabler, on ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timetableDayDo) Group(cols ...field.Expr) *timetableDayDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timetableDayDo) Having(conds ...gen.Condition) *timetableDayDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timetableDayDo) Limit(limit int) *timetableDayDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timetableDayDo) Offset(offset int) *timetableDayDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timetableDayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timetableDayDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timetableDayDo) Unscoped() *timetableDayDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timetableDayDo) Create(values ...*model.TimetableDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timetableDayDo) CreateInBatches(values []*model.TimetableDay, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timetableDayDo) Save(values ...*model.TimetableDay) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timetableDayDo) First() (*model.TimetableDay, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimetableDay), nil
	}
}

func (t timetableDayDo) Take() (*model.TimetableDay, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimetableDay), nil
	}
}

func (t timetableDayDo) Last() (*model.TimetableDay, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimetableDay), nil
	}
}

func (t timetableDayDo) Find() ([]*model.TimetableDay, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimetableDay), err
}

func (t timetableDayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimetableDay, err error) {
	buf := make([]*model.TimetableDay, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timetableDayDo) FindInBatches(result *[]*model.TimetableDay, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timetableDayDo) Attrs(attrs ...field.AssignExpr) *timetableDayDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timetableDayDo) Assign(attrs ...field.AssignExpr) *timetableDayDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timetableDayDo) Joins(fields ...field.RelationField) *timetableDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timetableDayDo) Preload(fields ...field.RelationField) *timetableDayDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timetableDayDo) FirstOrInit() (*model.TimetableDay, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimetableDay), nil
	}
}

func (t timetableDayDo) FirstOrCreate() (*model.TimetableDay, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimetableDay), nil
	}
}

func (t timetableDayDo) FindByPage(offset int, limit int) (result []*model.TimetableDay, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t timetableDayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timetableDayDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timetableDayDo) Delete(models ...*model.TimetableDay) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timetableDayDo) withDO(do gen.Dao) *timetableDayDo {
	t.DO = *do.(*gen.DO)
	return t
}
