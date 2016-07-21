package public

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// SourceMedia is the media type used to render source.
var SourceMedia = MediaType("application/vnd.source+json", func() {
	Description("SourceMedia is the media type used to render sources")
	TypeName("SourceMedia")
	Reference(SourcePayload)

	Attributes(func() {
		Attribute("id", Integer, "Source identifier")
		Attribute("href", String, "Source href")
		Attribute("name")
		Attribute("caches", CollectionOf(CacheMedia))
		Required("id", "href", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("caches")
	})
})

// SourceMediaLink is the media type used to render source.
var SourceMediaLink = MediaType("application/vnd.sourcelink+json", func() {
	Description("SourceMediaLink is the media type used to render sources links")
	TypeName("SourceMediaLink")
	Reference(SourcePayload)

	Attributes(func() {
		Attribute("id", Integer, "Source identifier")
		Attribute("href", String, "Source href")
		Attribute("name")
		Required("id", "href", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("href")
	})
})

// SourcePayload is the type used to create source.
var SourcePayload = Type("SourcePayload", func() {
	Description("SourcePayload is the type used to create sources")
	Reference(SourcePatchPayload)
	Attribute("name")
	Required("name")
})

// SourcePatchPayload is the type used to patch sources.
var SourcePatchPayload = Type("SourcePatchPayload", func() {
	Description("SourcePatchPayload is the type used to patch registrations")
	Attribute("name", String, "source name", func() {
		MinLength(2)
	})
})
