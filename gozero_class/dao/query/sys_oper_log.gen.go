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

	"github.com/smallq_class/dao/model"
)

func newSysOperLog(db *gorm.DB, opts ...gen.DOOption) sysOperLog {
	_sysOperLog := sysOperLog{}

	_sysOperLog.sysOperLogDo.UseDB(db, opts...)
	_sysOperLog.sysOperLogDo.UseModel(&model.SysOperLog{})

	tableName := _sysOperLog.sysOperLogDo.TableName()
	_sysOperLog.ALL = field.NewAsterisk(tableName)
	_sysOperLog.OperID = field.NewInt64(tableName, "oper_id")
	_sysOperLog.Title = field.NewString(tableName, "title")
	_sysOperLog.BusinessType = field.NewInt32(tableName, "business_type")
	_sysOperLog.Method = field.NewString(tableName, "method")
	_sysOperLog.RequestMethod = field.NewString(tableName, "request_method")
	_sysOperLog.OperName = field.NewString(tableName, "oper_name")
	_sysOperLog.OperURL = field.NewString(tableName, "oper_url")
	_sysOperLog.OperIP = field.NewString(tableName, "oper_ip")
	_sysOperLog.OperParam = field.NewString(tableName, "oper_param")
	_sysOperLog.JSONResult = field.NewString(tableName, "json_result")
	_sysOperLog.Status = field.NewInt32(tableName, "status")
	_sysOperLog.UserID = field.NewInt64(tableName, "user_id")
	_sysOperLog.OperTime = field.NewTime(tableName, "oper_time")
	_sysOperLog.CostTime = field.NewInt64(tableName, "cost_time")

	_sysOperLog.fillFieldMap()

	return _sysOperLog
}

// sysOperLog 操作日志记录
type sysOperLog struct {
	sysOperLogDo sysOperLogDo

	ALL           field.Asterisk
	OperID        field.Int64  // 日志主键
	Title         field.String // 模块标题
	BusinessType  field.Int32  // 业务类型（0其它 1新增 2修改 3删除）
	Method        field.String // 方法名称
	RequestMethod field.String // 请求方式
	OperName      field.String // 操作人员
	OperURL       field.String // 请求URL
	OperIP        field.String // 主机地址
	OperParam     field.String // 请求参数
	JSONResult    field.String // 返回参数
	Status        field.Int32  // 操作状态（0正常 1异常）
	UserID        field.Int64  // 用户ID
	OperTime      field.Time   // 操作时间
	CostTime      field.Int64  // 消耗时间

	fieldMap map[string]field.Expr
}

func (s sysOperLog) Table(newTableName string) *sysOperLog {
	s.sysOperLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOperLog) As(alias string) *sysOperLog {
	s.sysOperLogDo.DO = *(s.sysOperLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOperLog) updateTableName(table string) *sysOperLog {
	s.ALL = field.NewAsterisk(table)
	s.OperID = field.NewInt64(table, "oper_id")
	s.Title = field.NewString(table, "title")
	s.BusinessType = field.NewInt32(table, "business_type")
	s.Method = field.NewString(table, "method")
	s.RequestMethod = field.NewString(table, "request_method")
	s.OperName = field.NewString(table, "oper_name")
	s.OperURL = field.NewString(table, "oper_url")
	s.OperIP = field.NewString(table, "oper_ip")
	s.OperParam = field.NewString(table, "oper_param")
	s.JSONResult = field.NewString(table, "json_result")
	s.Status = field.NewInt32(table, "status")
	s.UserID = field.NewInt64(table, "user_id")
	s.OperTime = field.NewTime(table, "oper_time")
	s.CostTime = field.NewInt64(table, "cost_time")

	s.fillFieldMap()

	return s
}

func (s *sysOperLog) WithContext(ctx context.Context) *sysOperLogDo {
	return s.sysOperLogDo.WithContext(ctx)
}

func (s sysOperLog) TableName() string { return s.sysOperLogDo.TableName() }

func (s sysOperLog) Alias() string { return s.sysOperLogDo.Alias() }

func (s sysOperLog) Columns(cols ...field.Expr) gen.Columns { return s.sysOperLogDo.Columns(cols...) }

func (s *sysOperLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOperLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["oper_id"] = s.OperID
	s.fieldMap["title"] = s.Title
	s.fieldMap["business_type"] = s.BusinessType
	s.fieldMap["method"] = s.Method
	s.fieldMap["request_method"] = s.RequestMethod
	s.fieldMap["oper_name"] = s.OperName
	s.fieldMap["oper_url"] = s.OperURL
	s.fieldMap["oper_ip"] = s.OperIP
	s.fieldMap["oper_param"] = s.OperParam
	s.fieldMap["json_result"] = s.JSONResult
	s.fieldMap["status"] = s.Status
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["oper_time"] = s.OperTime
	s.fieldMap["cost_time"] = s.CostTime
}

func (s sysOperLog) clone(db *gorm.DB) sysOperLog {
	s.sysOperLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysOperLog) replaceDB(db *gorm.DB) sysOperLog {
	s.sysOperLogDo.ReplaceDB(db)
	return s
}

type sysOperLogDo struct{ gen.DO }

func (s sysOperLogDo) Debug() *sysOperLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOperLogDo) WithContext(ctx context.Context) *sysOperLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOperLogDo) ReadDB() *sysOperLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOperLogDo) WriteDB() *sysOperLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOperLogDo) Session(config *gorm.Session) *sysOperLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysOperLogDo) Clauses(conds ...clause.Expression) *sysOperLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOperLogDo) Returning(value interface{}, columns ...string) *sysOperLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOperLogDo) Not(conds ...gen.Condition) *sysOperLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOperLogDo) Or(conds ...gen.Condition) *sysOperLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOperLogDo) Select(conds ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOperLogDo) Where(conds ...gen.Condition) *sysOperLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOperLogDo) Order(conds ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOperLogDo) Distinct(cols ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOperLogDo) Omit(cols ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOperLogDo) Join(table schema.Tabler, on ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOperLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOperLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOperLogDo) Group(cols ...field.Expr) *sysOperLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOperLogDo) Having(conds ...gen.Condition) *sysOperLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOperLogDo) Limit(limit int) *sysOperLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOperLogDo) Offset(offset int) *sysOperLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOperLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysOperLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOperLogDo) Unscoped() *sysOperLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOperLogDo) Create(values ...*model.SysOperLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOperLogDo) CreateInBatches(values []*model.SysOperLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOperLogDo) Save(values ...*model.SysOperLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOperLogDo) First() (*model.SysOperLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperLog), nil
	}
}

func (s sysOperLogDo) Take() (*model.SysOperLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperLog), nil
	}
}

func (s sysOperLogDo) Last() (*model.SysOperLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperLog), nil
	}
}

func (s sysOperLogDo) Find() ([]*model.SysOperLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOperLog), err
}

func (s sysOperLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperLog, err error) {
	buf := make([]*model.SysOperLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOperLogDo) FindInBatches(result *[]*model.SysOperLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOperLogDo) Attrs(attrs ...field.AssignExpr) *sysOperLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOperLogDo) Assign(attrs ...field.AssignExpr) *sysOperLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOperLogDo) Joins(fields ...field.RelationField) *sysOperLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOperLogDo) Preload(fields ...field.RelationField) *sysOperLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOperLogDo) FirstOrInit() (*model.SysOperLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperLog), nil
	}
}

func (s sysOperLogDo) FirstOrCreate() (*model.SysOperLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperLog), nil
	}
}

func (s sysOperLogDo) FindByPage(offset int, limit int) (result []*model.SysOperLog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysOperLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOperLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOperLogDo) Delete(models ...*model.SysOperLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOperLogDo) withDO(do gen.Dao) *sysOperLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
