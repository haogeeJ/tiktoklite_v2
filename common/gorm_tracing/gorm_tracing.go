package gorm_tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	traceLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const gormSpanKey = "gorm_span"

func before(db *gorm.DB) {
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")
	db.InstanceSet(gormSpanKey, span)
}
func after(db *gorm.DB) {
	_span, isExists := db.InstanceGet(gormSpanKey)
	if !isExists {
		return
	}
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	defer span.Finish()
	ext.DBType.Set(span, "MySQL")
	if db.Error != nil {
		span.LogFields(traceLog.Error(db.Error))
		return
	}
	span.LogFields(traceLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	return
}

const (
	opentracingBeforeName = "opentracing:before"
	opentracingAfterName  = "opentracing:after"
)

type OpentracingPlugin struct {
}

func (p *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}
func (p *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	db.Callback().Delete().Before("gorm:before_delete").Register(opentracingBeforeName, before)
	db.Callback().Create().Before("gorm:before_create").Register(opentracingBeforeName, before)
	db.Callback().Query().Before("gorm:before_query").Register(opentracingBeforeName, before)
	db.Callback().Update().Before("gorm:before_update").Register(opentracingBeforeName, before)
	db.Callback().Row().Before("gorm:before_row").Register(opentracingBeforeName, before)
	db.Callback().Raw().Before("gorm:before_raw").Register(opentracingBeforeName, before)

	db.Callback().Delete().After("gorm:after_delete").Register(opentracingAfterName, after)
	db.Callback().Create().After("gorm:after_create").Register(opentracingAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(opentracingAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(opentracingAfterName, after)
	db.Callback().Row().After("gorm:after_row").Register(opentracingAfterName, after)
	db.Callback().Raw().After("gorm:after_raw").Register(opentracingAfterName, after)
	return
}
