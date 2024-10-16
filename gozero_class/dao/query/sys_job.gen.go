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

func newSysJob(db *gorm.DB, opts ...gen.DOOption) sysJob {
	_sysJob := sysJob{}

	_sysJob.sysJobDo.UseDB(db, opts...)
	_sysJob.sysJobDo.UseModel(&model.SysJob{})

	tableName := _sysJob.sysJobDo.TableName()
	_sysJob.ALL = field.NewAsterisk(tableName)
	_sysJob.JobID = field.NewInt64(tableName, "job_id")
	_sysJob.JobName = field.NewString(tableName, "job_name")
	_sysJob.JobParams = field.NewString(tableName, "job_params")
	_sysJob.InvokeTarget = field.NewString(tableName, "invoke_target")
	_sysJob.CronExpression = field.NewString(tableName, "cron_expression")
	_sysJob.Status = field.NewString(tableName, "status")
	_sysJob.CreateBy = field.NewString(tableName, "create_by")
	_sysJob.CreateTime = field.NewTime(tableName, "create_time")
	_sysJob.UpdateBy = field.NewString(tableName, "update_by")
	_sysJob.UpdateTime = field.NewTime(tableName, "update_time")

	_sysJob.fillFieldMap()

	return _sysJob
}

// sysJob 定时任务调度表
type sysJob struct {
	sysJobDo sysJobDo

	ALL            field.Asterisk
	JobID          field.Int64  // 任务ID
	JobName        field.String // 任务名称
	JobParams      field.String // 参数
	InvokeTarget   field.String // 调用目标字符串
	CronExpression field.String // cron执行表达式
	Status         field.String // 状态（0正常 1暂停）
	CreateBy       field.String // 创建者
	CreateTime     field.Time   // 创建时间
	UpdateBy       field.String // 更新者
	UpdateTime     field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (s sysJob) Table(newTableName string) *sysJob {
	s.sysJobDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysJob) As(alias string) *sysJob {
	s.sysJobDo.DO = *(s.sysJobDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysJob) updateTableName(table string) *sysJob {
	s.ALL = field.NewAsterisk(table)
	s.JobID = field.NewInt64(table, "job_id")
	s.JobName = field.NewString(table, "job_name")
	s.JobParams = field.NewString(table, "job_params")
	s.InvokeTarget = field.NewString(table, "invoke_target")
	s.CronExpression = field.NewString(table, "cron_expression")
	s.Status = field.NewString(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")

	s.fillFieldMap()

	return s
}

func (s *sysJob) WithContext(ctx context.Context) *sysJobDo { return s.sysJobDo.WithContext(ctx) }

func (s sysJob) TableName() string { return s.sysJobDo.TableName() }

func (s sysJob) Alias() string { return s.sysJobDo.Alias() }

func (s sysJob) Columns(cols ...field.Expr) gen.Columns { return s.sysJobDo.Columns(cols...) }

func (s *sysJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysJob) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["job_id"] = s.JobID
	s.fieldMap["job_name"] = s.JobName
	s.fieldMap["job_params"] = s.JobParams
	s.fieldMap["invoke_target"] = s.InvokeTarget
	s.fieldMap["cron_expression"] = s.CronExpression
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
}

func (s sysJob) clone(db *gorm.DB) sysJob {
	s.sysJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysJob) replaceDB(db *gorm.DB) sysJob {
	s.sysJobDo.ReplaceDB(db)
	return s
}

type sysJobDo struct{ gen.DO }

func (s sysJobDo) Debug() *sysJobDo {
	return s.withDO(s.DO.Debug())
}

func (s sysJobDo) WithContext(ctx context.Context) *sysJobDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysJobDo) ReadDB() *sysJobDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysJobDo) WriteDB() *sysJobDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysJobDo) Session(config *gorm.Session) *sysJobDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysJobDo) Clauses(conds ...clause.Expression) *sysJobDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysJobDo) Returning(value interface{}, columns ...string) *sysJobDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysJobDo) Not(conds ...gen.Condition) *sysJobDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysJobDo) Or(conds ...gen.Condition) *sysJobDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysJobDo) Select(conds ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysJobDo) Where(conds ...gen.Condition) *sysJobDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysJobDo) Order(conds ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysJobDo) Distinct(cols ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysJobDo) Omit(cols ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysJobDo) Join(table schema.Tabler, on ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysJobDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysJobDo) Group(cols ...field.Expr) *sysJobDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysJobDo) Having(conds ...gen.Condition) *sysJobDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysJobDo) Limit(limit int) *sysJobDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysJobDo) Offset(offset int) *sysJobDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysJobDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysJobDo) Unscoped() *sysJobDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysJobDo) Create(values ...*model.SysJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysJobDo) CreateInBatches(values []*model.SysJob, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysJobDo) Save(values ...*model.SysJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysJobDo) First() (*model.SysJob, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Take() (*model.SysJob, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Last() (*model.SysJob, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) Find() ([]*model.SysJob, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysJob), err
}

func (s sysJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysJob, err error) {
	buf := make([]*model.SysJob, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysJobDo) FindInBatches(result *[]*model.SysJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysJobDo) Attrs(attrs ...field.AssignExpr) *sysJobDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysJobDo) Assign(attrs ...field.AssignExpr) *sysJobDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysJobDo) Joins(fields ...field.RelationField) *sysJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysJobDo) Preload(fields ...field.RelationField) *sysJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysJobDo) FirstOrInit() (*model.SysJob, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) FirstOrCreate() (*model.SysJob, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJob), nil
	}
}

func (s sysJobDo) FindByPage(offset int, limit int) (result []*model.SysJob, count int64, err error) {
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

func (s sysJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysJobDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysJobDo) Delete(models ...*model.SysJob) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysJobDo) withDO(do gen.Dao) *sysJobDo {
	s.DO = *do.(*gen.DO)
	return s
}
