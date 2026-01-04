package types

import (
	"errors"
	"fmt"
	"sort"

	"d3tech.com/platform/utils"
)

// AggregateMethod represents the type of aggregation operation
type AggregateMethod string

const (
	AggregateSum   AggregateMethod = "sum"
	AggregateAvg   AggregateMethod = "avg"
	AggregateMin   AggregateMethod = "min"
	AggregateMax   AggregateMethod = "max"
	AggregateCount AggregateMethod = "count"
	AggregateFirst AggregateMethod = "first"
	AggregateLast  AggregateMethod = "last"
)

// ValidAggregateMethods lists all supported aggregation methods
var ValidAggregateMethods = []AggregateMethod{
	AggregateSum,
	AggregateAvg,
	AggregateMin,
	AggregateMax,
	AggregateCount,
	AggregateFirst,
	AggregateLast,
}

// NumericAggregateMethods lists methods that only work with numeric fields
var NumericAggregateMethods = []AggregateMethod{
	AggregateSum,
	AggregateAvg,
	AggregateMin,
	AggregateMax,
}

func (m AggregateMethod) IsValid() bool {
	for _, valid := range ValidAggregateMethods {
		if m == valid {
			return true
		}
	}
	return false
}

func (m AggregateMethod) RequiresNumericField() bool {
	for _, numeric := range NumericAggregateMethods {
		if m == numeric {
			return true
		}
	}
	return false
}

// ToMongoOperator returns the MongoDB aggregation operator for this method
func (m AggregateMethod) ToMongoOperator() string {
	switch m {
	case AggregateSum:
		return "$sum"
	case AggregateAvg:
		return "$avg"
	case AggregateMin:
		return "$min"
	case AggregateMax:
		return "$max"
	case AggregateCount:
		return "$sum"
	case AggregateFirst:
		return "$first"
	case AggregateLast:
		return "$last"
	default:
		return ""
	}
}

// AggregateFieldDef defines an aggregation on a specific field
type AggregateFieldDef struct {
	// Name is the field to aggregate
	Field string `yaml:"field"`
	// Methods are the aggregation methods to apply (e.g., sum, avg)
	Methods []string `yaml:"methods"`
}

// ObjectAggregationDef defines aggregation configuration for an object
type ObjectAggregationDef struct {
	// Enabled indicates if aggregations are enabled for this object
	Enabled bool `yaml:"enabled"`
	// Fields are the fields that can be aggregated
	Fields []AggregateFieldDef `yaml:"fields"`
	// GroupBy are the fields that can be used for grouping
	GroupBy []string `yaml:"groupBy"`
}

// Validate validates the aggregation definition against the object's fields
func (a *ObjectAggregationDef) Validate(obj *Object, registry Registry) error {
	if !a.Enabled {
		return nil
	}

	var err error

	// Validate aggregation fields
	for _, aggField := range a.Fields {
		field, ok := obj.GetField(aggField.Field)
		if !ok {
			err = errors.Join(err, fmt.Errorf("aggregation field %s not found in object %s", aggField.Field, obj.Name))
			continue
		}

		// Validate each method
		for _, method := range aggField.Methods {
			aggMethod := AggregateMethod(method)
			if !aggMethod.IsValid() {
				err = errors.Join(err, fmt.Errorf("invalid aggregation method %s for field %s", method, aggField.Field))
				continue
			}

			// Validate numeric requirement
			if aggMethod.RequiresNumericField() && !isNumericFieldType(field.Type) {
				err = errors.Join(err, fmt.Errorf("aggregation method %s requires numeric field, but %s is of type %s", method, aggField.Field, field.Type))
			}
		}
	}

	// Validate groupBy fields
	for _, groupByField := range a.GroupBy {
		if _, ok := obj.GetField(groupByField); !ok {
			err = errors.Join(err, fmt.Errorf("groupBy field %s not found in object %s", groupByField, obj.Name))
		}
	}

	return err
}

// ListAggregateFields returns the list of aggregatable fields for this object
func (a *ObjectAggregationDef) ListAggregateFields() []AggregateFieldDef {
	return a.Fields
}

// ListGroupByFields returns the list of fields that can be grouped by
func (a *ObjectAggregationDef) ListGroupByFields() []string {
	return a.GroupBy
}

// GetAggregateFieldMethods returns the methods available for a specific field
func (a *ObjectAggregationDef) GetAggregateFieldMethods(fieldName string) []AggregateMethod {
	for _, field := range a.Fields {
		if field.Field == fieldName {
			methods := make([]AggregateMethod, len(field.Methods))
			for i, m := range field.Methods {
				methods[i] = AggregateMethod(m)
			}
			return methods
		}
	}
	return nil
}

// isNumericFieldType checks if a field type is numeric
func isNumericFieldType(fieldType string) bool {
	switch FieldType(fieldType) {
	case FieldTypeInt, FieldTypeInt32, FieldTypeInt64:
		return true
	default:
		return false
	}
}

// AggregatedField represents a field with an aggregation result
type AggregatedField struct {
	Field  string          `json:"field"`
	Method AggregateMethod `json:"method"`
	Alias  string          `json:"alias"`
}

// GetAlias returns the alias for this aggregated field (e.g., "totalAmount", "avgPrice")
func (af *AggregatedField) GetAlias() string {
	if af.Alias != "" {
		return af.Alias
	}
	return string(af.Method) + utils.UCC(af.Field)
}

// SortAggregationDef sorts the aggregation definition fields and groupBy for deterministic output
func SortAggregationDef(a *ObjectAggregationDef) {
	if a == nil {
		return
	}
	sort.Slice(a.Fields, func(i, j int) bool {
		return a.Fields[i].Field < a.Fields[j].Field
	})
	sort.Strings(a.GroupBy)
	for i := range a.Fields {
		sort.Strings(a.Fields[i].Methods)
	}
}
