package public

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CacheMedia is the media type used to render caches.
var CacheMedia = MediaType("application/vnd.caches+json", func() {
	Description("CacheMedia is the media type used to render caches")
	TypeName("CacheMedia")
	Reference(CachePayload)

	Attributes(func() {
		Attribute("id", Integer, "Cache identifier")
		Attribute("href", String, "Cache href")
		Attribute("name")
		Attribute("description")
		Attribute("source", SourceMedia, "Source being cached")
		Attribute("text", String, "contents in source")
		Link("source")
		Required("id", "href", "name", "source", "text")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("description")
		Attribute("source")
		Attribute("text")
	})

	View("extended", func() {
		Attribute("id")
		Attribute("href")
		Attribute("source")
		Attribute("name")
	})
})

// CachePayload is the type used to create cache.
var CachePayload = Type("CachePayload", func() {
	Description("CachePayload is the type used to create caches")
	Reference(CachePatchPayload)
	Attribute("name", String, "Cache name", func() {
		MinLength(3)
	})
	Attribute("description")
	Attribute("text")
	Attribute("source_href", String, "The href to the source resource that describes the source being taught", func() {
		Pattern("/caches/[0-9]+")
	})
	Required("name", "description")
})

// CachePatchPayload is the type used to patch caches.
var CachePatchPayload = Type("CachePatchPayload", func() {
	Description("CachePatchPayload is the type used to patch caches")
	Attribute("description", String, "Cache description")
	Attribute("text", String, "Contents for cache", func() {
		MinLength(2)
	})
})