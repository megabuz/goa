package apidsl

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
)

// Metadata is a set of key/value pairs that can be assigned
// to an object. Each value consists of a slice of strings so
// that multiple invocation of the Metadata function on the
// same target using the same key builds up the slice.
//
// While keys can have any value the following names are
// handled explicitly by goagen:
//
// "struct:tag:xxx": sets the struct field tag xxx on generated structs.
//               Overrides tags that goagen would otherwise set.
//               If the metadata value is a slice then the
//               strings are joined with the space character as
//               separator.
//    Example:
//        Metadata("struct:tag:json", "myName,omitempty")
//        Metadata("struct:tag:xml", "myName,attr")
//
// "swagger:tag:xxx": sets the Swagger object field tag xxx.
//
//    Example:
//        Metadata("swagger:tag:Backend")
//        Metadata("swagger:tag:Backend:desc", "Quick description of what 'Backend' is")
//        Metadata("swagger:tag:Backend:url", "http://example.com")
//        Metadata("swagger:tag:Backend:url:desc", "See more docs here")
//
func Metadata(name string, value ...string) {
	switch def := dslengine.CurrentDefinition().(type) {
	case *design.AttributeDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.MediaTypeDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ActionDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ResourceDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.ResponseDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	case *design.APIDefinition:
		if def.Metadata == nil {
			def.Metadata = make(map[string][]string)
		}
		def.Metadata[name] = append(def.Metadata[name], value...)

	default:
		dslengine.IncompatibleDSL()
	}
}
