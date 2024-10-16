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

func newGenTableColumn(db *gorm.DB, opts ...gen.DOOption) genTableColumn {
	_genTableColumn := genTableColumn{}

	_genTableColumn.genTableColumnDo.UseDB(db, opts...)
	_genTableColumn.genTableColumnDo.UseModel(&model.GenTableColumn{})

	tableName := _genTableColumn.genTableColumnDo.TableName()
	_genTableColumn.ALL = field.NewAsterisk(tableName)
	_genTableColumn.ColumnID = field.NewInt64(tableName, "column_id")
	_genTableColumn.TableID = field.NewString(tableName, "table_id")
	_genTableColumn.ColumnName = field.NewString(tableName, "column_name")
	_genTableColumn.ColumnComment = field.NewString(tableName, "column_comment")
	_genTableColumn.ColumnType = field.NewString(tableName, "column_type")
	_genTableColumn.GoType = field.NewString(tableName, "go_type")
	_genTableColumn.GoField = field.NewString(tableName, "go_field")
	_genTableColumn.IsPk = field.NewString(tableName, "is_pk")
	_genTableColumn.IsRequired = field.NewString(tableName, "is_required")
	_genTableColumn.IsInsert = field.NewString(tableName, "is_insert")
	_genTableColumn.IsEdit = field.NewString(tableName, "is_edit")
	_genTableColumn.IsList = field.NewString(tableName, "is_list")
	_genTableColumn.IsQuery = field.NewString(tableName, "is_query")
	_genTableColumn.QueryType = field.NewString(tableName, "query_type")
	_genTableColumn.HTMLType = field.NewString(tableName, "html_type")
	_genTableColumn.DictType = field.NewString(tableName, "dict_type")
	_genTableColumn.Sort = field.NewInt32(tableName, "sort")
	_genTableColumn.CreateBy = field.NewInt64(tableName, "create_by")
	_genTableColumn.CreateTime = field.NewTime(tableName, "create_time")
	_genTableColumn.UpdateBy = field.NewInt64(tableName, "update_by")
	_genTableColumn.UpdateTime = field.NewTime(tableName, "update_time")
	_genTableColumn.HTMLField = field.NewString(tableName, "html_field")

	_genTableColumn.fillFieldMap()

	return _genTableColumn
}

// genTableColumn 代码生成业务表字段
type genTableColumn struct {
	genTableColumnDo genTableColumnDo

	ALL           field.Asterisk
	ColumnID      field.Int64  // 编号
	TableID       field.String // 归属表编号
	ColumnName    field.String // 列名称
	ColumnComment field.String // 列描述
	ColumnType    field.String // 列类型
	GoType        field.String // go类型
	GoField       field.String // go字段名
	IsPk          field.String // 是否主键（1是）
	IsRequired    field.String // 是否必填（1是）
	IsInsert      field.String // 是否为插入字段（1是）
	IsEdit        field.String // 是否编辑字段（1是）
	IsList        field.String // 是否列表字段（1是）
	IsQuery       field.String // 是否查询字段（1是）
	QueryType     field.String // 查询方式（等于、不等于、大于、小于、范围）
	HTMLType      field.String // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      field.String // 字典类型
	Sort          field.Int32  // 排序
	CreateBy      field.Int64  // 创建者
	CreateTime    field.Time   // 创建时间
	UpdateBy      field.Int64  // 更新者
	UpdateTime    field.Time   // 更新时间
	HTMLField     field.String

	fieldMap map[string]field.Expr
}

func (g genTableColumn) Table(newTableName string) *genTableColumn {
	g.genTableColumnDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g genTableColumn) As(alias string) *genTableColumn {
	g.genTableColumnDo.DO = *(g.genTableColumnDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *genTableColumn) updateTableName(table string) *genTableColumn {
	g.ALL = field.NewAsterisk(table)
	g.ColumnID = field.NewInt64(table, "column_id")
	g.TableID = field.NewString(table, "table_id")
	g.ColumnName = field.NewString(table, "column_name")
	g.ColumnComment = field.NewString(table, "column_comment")
	g.ColumnType = field.NewString(table, "column_type")
	g.GoType = field.NewString(table, "go_type")
	g.GoField = field.NewString(table, "go_field")
	g.IsPk = field.NewString(table, "is_pk")
	g.IsRequired = field.NewString(table, "is_required")
	g.IsInsert = field.NewString(table, "is_insert")
	g.IsEdit = field.NewString(table, "is_edit")
	g.IsList = field.NewString(table, "is_list")
	g.IsQuery = field.NewString(table, "is_query")
	g.QueryType = field.NewString(table, "query_type")
	g.HTMLType = field.NewString(table, "html_type")
	g.DictType = field.NewString(table, "dict_type")
	g.Sort = field.NewInt32(table, "sort")
	g.CreateBy = field.NewInt64(table, "create_by")
	g.CreateTime = field.NewTime(table, "create_time")
	g.UpdateBy = field.NewInt64(table, "update_by")
	g.UpdateTime = field.NewTime(table, "update_time")
	g.HTMLField = field.NewString(table, "html_field")

	g.fillFieldMap()

	return g
}

func (g *genTableColumn) WithContext(ctx context.Context) *genTableColumnDo {
	return g.genTableColumnDo.WithContext(ctx)
}

func (g genTableColumn) TableName() string { return g.genTableColumnDo.TableName() }

func (g genTableColumn) Alias() string { return g.genTableColumnDo.Alias() }

func (g genTableColumn) Columns(cols ...field.Expr) gen.Columns {
	return g.genTableColumnDo.Columns(cols...)
}

func (g *genTableColumn) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *genTableColumn) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 22)
	g.fieldMap["column_id"] = g.ColumnID
	g.fieldMap["table_id"] = g.TableID
	g.fieldMap["column_name"] = g.ColumnName
	g.fieldMap["column_comment"] = g.ColumnComment
	g.fieldMap["column_type"] = g.ColumnType
	g.fieldMap["go_type"] = g.GoType
	g.fieldMap["go_field"] = g.GoField
	g.fieldMap["is_pk"] = g.IsPk
	g.fieldMap["is_required"] = g.IsRequired
	g.fieldMap["is_insert"] = g.IsInsert
	g.fieldMap["is_edit"] = g.IsEdit
	g.fieldMap["is_list"] = g.IsList
	g.fieldMap["is_query"] = g.IsQuery
	g.fieldMap["query_type"] = g.QueryType
	g.fieldMap["html_type"] = g.HTMLType
	g.fieldMap["dict_type"] = g.DictType
	g.fieldMap["sort"] = g.Sort
	g.fieldMap["create_by"] = g.CreateBy
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["update_by"] = g.UpdateBy
	g.fieldMap["update_time"] = g.UpdateTime
	g.fieldMap["html_field"] = g.HTMLField
}

func (g genTableColumn) clone(db *gorm.DB) genTableColumn {
	g.genTableColumnDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g genTableColumn) replaceDB(db *gorm.DB) genTableColumn {
	g.genTableColumnDo.ReplaceDB(db)
	return g
}

type genTableColumnDo struct{ gen.DO }

func (g genTableColumnDo) Debug() *genTableColumnDo {
	return g.withDO(g.DO.Debug())
}

func (g genTableColumnDo) WithContext(ctx context.Context) *genTableColumnDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g genTableColumnDo) ReadDB() *genTableColumnDo {
	return g.Clauses(dbresolver.Read)
}

func (g genTableColumnDo) WriteDB() *genTableColumnDo {
	return g.Clauses(dbresolver.Write)
}

func (g genTableColumnDo) Session(config *gorm.Session) *genTableColumnDo {
	return g.withDO(g.DO.Session(config))
}

func (g genTableColumnDo) Clauses(conds ...clause.Expression) *genTableColumnDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g genTableColumnDo) Returning(value interface{}, columns ...string) *genTableColumnDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g genTableColumnDo) Not(conds ...gen.Condition) *genTableColumnDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g genTableColumnDo) Or(conds ...gen.Condition) *genTableColumnDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g genTableColumnDo) Select(conds ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g genTableColumnDo) Where(conds ...gen.Condition) *genTableColumnDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g genTableColumnDo) Order(conds ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g genTableColumnDo) Distinct(cols ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g genTableColumnDo) Omit(cols ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g genTableColumnDo) Join(table schema.Tabler, on ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g genTableColumnDo) LeftJoin(table schema.Tabler, on ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g genTableColumnDo) RightJoin(table schema.Tabler, on ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g genTableColumnDo) Group(cols ...field.Expr) *genTableColumnDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g genTableColumnDo) Having(conds ...gen.Condition) *genTableColumnDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g genTableColumnDo) Limit(limit int) *genTableColumnDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g genTableColumnDo) Offset(offset int) *genTableColumnDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g genTableColumnDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *genTableColumnDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g genTableColumnDo) Unscoped() *genTableColumnDo {
	return g.withDO(g.DO.Unscoped())
}

func (g genTableColumnDo) Create(values ...*model.GenTableColumn) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g genTableColumnDo) CreateInBatches(values []*model.GenTableColumn, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g genTableColumnDo) Save(values ...*model.GenTableColumn) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g genTableColumnDo) First() (*model.GenTableColumn, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenTableColumn), nil
	}
}

func (g genTableColumnDo) Take() (*model.GenTableColumn, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenTableColumn), nil
	}
}

func (g genTableColumnDo) Last() (*model.GenTableColumn, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenTableColumn), nil
	}
}

func (g genTableColumnDo) Find() ([]*model.GenTableColumn, error) {
	result, err := g.DO.Find()
	return result.([]*model.GenTableColumn), err
}

func (g genTableColumnDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GenTableColumn, err error) {
	buf := make([]*model.GenTableColumn, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g genTableColumnDo) FindInBatches(result *[]*model.GenTableColumn, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g genTableColumnDo) Attrs(attrs ...field.AssignExpr) *genTableColumnDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g genTableColumnDo) Assign(attrs ...field.AssignExpr) *genTableColumnDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g genTableColumnDo) Joins(fields ...field.RelationField) *genTableColumnDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g genTableColumnDo) Preload(fields ...field.RelationField) *genTableColumnDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g genTableColumnDo) FirstOrInit() (*model.GenTableColumn, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenTableColumn), nil
	}
}

func (g genTableColumnDo) FirstOrCreate() (*model.GenTableColumn, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenTableColumn), nil
	}
}

func (g genTableColumnDo) FindByPage(offset int, limit int) (result []*model.GenTableColumn, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g genTableColumnDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g genTableColumnDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g genTableColumnDo) Delete(models ...*model.GenTableColumn) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *genTableColumnDo) withDO(do gen.Dao) *genTableColumnDo {
	g.DO = *do.(*gen.DO)
	return g
}
