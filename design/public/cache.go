package public

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CacheMedia is the media type used to render caches.
var CacheMedia = MediaType("application/vnd.cache+json", func() {
	Description("Cache is the media type used to render caches")
	TypeName("Cache")
	Reference(CachePayload)

	Attributes(func() {
		Attribute("id", Integer, "Cache identifier")
		Attribute("href", String, "Cache href")
		Attribute("name")
		Attribute("description")
		Attribute("source", SourceMediaLink, "Id of Source")
		Attribute("text", String, "contents in source")
		Required("id", "href", "name", "text")
	})

	View("link", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("text")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("href")
	})

	View("extended", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("source")
		Attribute("text")
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
	Attribute("source_id", Integer, "Id of Source")
	Required("name", "description")
})

// CachePatchPayload is the type used to patch caches.
var CachePatchPayload = Type("CachePatchPayload", func() {
	Description("CachePatchPayload is the type used to patch caches")
	Attribute("description", String, "Cache description")
	Attribute("text", String, "Contents for cache", func() {
		MinLength(2)
	})
	Attribute("source_id", Integer, "Id of Source")
})
