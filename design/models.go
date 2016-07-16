package design

import (
	"github.com/atclate/goa-exercise/design/public"
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("StorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the Mysql relational store")
		Model("Source", func() {
			BuildsFrom(func() {
				Payload("source", "create")
				Payload("source", "update")
			})
			RendersTo(public.SourceMedia)
			HasMany("Caches", "Cache")
			Description("Source Model Description")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Source Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Cache", func() {
			BuildsFrom(func() {
				Payload("cache", "create")
				Payload("cache", "update")
			})
			RendersTo(public.CacheMedia)
			Description("Cache Model")
			BelongsTo("Source")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Cache Payload Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("description", gorma.String, func() {})
			Field("text", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
