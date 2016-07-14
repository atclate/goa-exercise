package design

import (
	. "github.com/atclate/goa-exercise/design/public"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Define a few constants so we can easily switch between development and production values.
const (
	// Hostname is the service hostname and port in the form "hostname:port".
	// This is mainly used by the documentation, Swagger and client tool code.
	Hostname = "localhost:8080"

	// HTTPScheme is the scheme used by the API. As with Hostname this is mainly
	// used used by the documentation, Swagger and client tool code.
	HTTPScheme = "http"
)

// This block defines the global properties of the service API.
var _ = API("GoWorkshop", func() {

	// General metadata about the service
	Title("The Universal Workshop Service")
	Description("GoWorkshop is a simple example service that exposes a REST API")
	Version("1.0")
	Contact(func() {
		Name("The GoWorkshop developers")
		Email("ac1493@yp.com")
	})
	License(func() {
		Name("The MIT License (MIT)")
		URL("https://github.com/gophercon/buildingapis/blob/master/LICENSE")
	})

	// Endpoint definition
	Host(Hostname)
	Scheme(HTTPScheme)
	BasePath("/api")

	// Encoding and decoding
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("source", func() {

	Description("The source resource exposes the endpoints used to manage workshop sources")
	BasePath("/sources")

	// Create source
	Action("create", func() {
		Description("Create a new source")
		Routing(POST("/"))
		Payload(SourcePayload)
		Response(Created, func() {
			Headers(func() {
				Header("Location", String, "Newly created source href", func() {
					Pattern("/caches/[0-9]+")
				})
			})
			Media(SourceMedia)
		})
		Response(BadRequest, ErrorMedia)
	})

	// Retrieve source
	Action("show", func() {
		Description("Retrieve a source by ID")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "The source ID", func() {
				Minimum(1)
			})
		})
		Response(OK, SourceMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	// Delete a source
	Action("delete", func() {
		Description("Delete a source")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer, "The source ID", func() {
				Minimum(1)
			})
		})
		Response(NoContent)
		Response(NotFound)
	})

})

// This block defines the "cache" resource endpoints.
var _ = Resource("cache", func() {

	Description("The cache resource exposes the endpoints used to manage workshop caches")
	BasePath("/caches")

	// Create cache
	Action("create", func() {
		Description("Create a new cache")
		Routing(POST("/"))
		Payload(CachePayload)
		Response(Created, func() {
			Headers(func() {
				Header("Location", String, "Newly created cache href", func() {
					Pattern("/caches/[0-9]+")
				})
			})
			Media(CacheMedia)
		})
		Response(BadRequest, ErrorMedia)
	})

	// Retrieve cache
	Action("show", func() {
		Description("Retrieve a cache by ID")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "The cache ID", func() {
				Minimum(1)
			})
			Param("view", String, "The view used to render the cache", func() {
				Enum("default", "extended")
			})
		})
		Response(OK, CacheMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	// List caches
	Action("list", func() {
		Description("List all caches")
		Routing(GET("/"))
		Params(func() {
			Param("source_id", Integer, "Filter by source", func() {
				Minimum(1)
			})
		})
		Response(OK, CollectionOf(CacheMedia))
	})
})

// This block defines the "swagger" resource which serves static files (swagger definitions)
var _ = Resource("public", func() {

	Description("The public resource groups endpoints that serve static content")

	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})

	// Swagger JSON
	Files("/swagger.json", "swagger/swagger.json", func() {
		//Header("Access-Control-Allow-Origin", "http://swagger.io")
		//Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		Description("API Swagger spec in JSON format")
	})

	// Swagger YAML
	Files("/swagger.yaml", "swagger/swagger.yaml", func() {
		Description("API Swagger spec in YAML format")
	})

	// Swagger UI
	Files("/swagger/*file", "public/", func() {
		Description("Swagger UI")
	})
})

/*
var _ = StorageGroup("StorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the Postgres relational store")
		Model("SourcePayload", func() {
			BuildsFrom(func() {
				Payload("user", "create")
				Payload("user", "update")
			})
			RendersTo(User)
			Description("User Model Description")
			HasMany("Reviews", "Review")
			HasMany("Proposals", "Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Proposal", func() {
			BuildsFrom(func() {
				Payload("proposal", "create")
				Payload("proposal", "update")
			})
			RendersTo(Proposal)
			Description("Proposal Model")
			BelongsTo("User")
			HasMany("Reviews", "Review")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Payload Model PK field")
			})
			Field("title", func() {
				Alias("proposal_title")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Review", func() {
			BuildsFrom(func() {
				Payload("review", "create")
				Payload("review", "update")
			})
			RendersTo(Review)
			Description("Review Model")
			BelongsTo("User")
			BelongsTo("Proposal")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Review Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Test", func() {
			Description("TestModel")
			NoAutomaticIDFields()
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
*/
