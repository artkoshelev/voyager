package wiringplugin

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

// ShapeName is a globally unique identifier for the type of a shape.
type ShapeName string

// Shape represents an autowiring shape.
// Shapes are bits of information that an autowiring function exposes to provide information to other functions that
// depend on that resource.
// This is pretty much the same as Bazel providers. See https://docs.bazel.build/versions/master/skylark/rules.html#providers
//
// Shapes in JSON look like this:
// {
//    "name": "voyager.atl-paas.net/MyShape",
//    "data": {
//      "field1": 42,
//      "field2": { "a": 7, "b": "x" }
//    }
// }
type Shape interface {
	// Name returns the name of the shape.
	Name() ShapeName
	// DeepCopyShape returns a deep copy of the shape.
	DeepCopyShape() Shape
}

// ShapeMeta is a reusable container for bits of information common to all shapes.
type ShapeMeta struct {
	// ShapeName is the name of the shape.
	ShapeName ShapeName `json:"name"`
}

// BindableShapeStruct represents a bit of information that is needed to create a Service Catalog ServiceBinding
// object. To be embedded into other shapes' structs where a ServiceInstance needs to be bound to to get outputs
// for that shape.
// If an autowiring plugin exposes multiple shapes that have this struct embedded it may or may not be the case
// that they all refer to the same ServiceInstance. It is responsibility of the consuming side to track if more than
// one ServiceBinding needs to be created to consume values from those shapes.
// +k8s:deepcopy-gen=true
type BindableShapeStruct struct {
	ServiceInstanceName ProtoReference `json:"serviceInstanceName"`
}

// CopyShape copies source shape into the target shape.
// The main purpose of this function is to convert between typed and UnstructuredShape representations
// of the same shape.
func CopyShape(source, target Shape) error {
	st := reflect.TypeOf(source)
	tt := reflect.TypeOf(target)
	switch {
	case st == tt:
		// Same type already, just deep copy
		sv := reflect.ValueOf(source)
		tv := reflect.ValueOf(target)
		method := sv.MethodByName("DeepCopyInto")
		method.Call([]reflect.Value{tv})
	default:
		// Convert by round-tripping to and from JSON
		data, err := json.Marshal(source)
		if err != nil {
			return errors.WithStack(err)
		}
		err = json.Unmarshal(data, target)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
