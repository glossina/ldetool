/*
 This file was autogenerated via
 -------------------------------------------------------------------------------
 gen-builtin --declarable --type Hex --name hex --handler AddUint --go-name uint
 -------------------------------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Hex("")

// Hex represents field of type hex
type Hex string

// Name returns field name
func (i Hex) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Hex) TypeName() string {
	return "hex"
}

// Register registers a field
func (i Hex) Register(registrator FieldRegistrator) {
	registrator.AddUint(i.Name())
}

// GoName returns Go's representation of this field's type
func (i Hex) GoName() string {
	return "uint"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["hex"] = func(fieldName string) Field {
		return Hex(fieldName)
	}

	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["hex"] = struct{}{}

	if backedBy == nil {
		backedBy = map[string]string{}
	}
	backedBy["hex"] = "uint"
}