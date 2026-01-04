package types

import (
	"errors"
	"fmt"
	"sort"

	"github.com/JacobDoucet/forge/utils"
)

type Object struct {
	Name            string               `yaml:"name"`
	Collection      []Collection         `yaml:"collection"`
	HTTP            ObjectHTTPDef        `yaml:"http"`
	Indexes         []Index              `yaml:"indexes"`
	Fields          []Field              `yaml:"fields"`
	Permissions     ObjectPermissions    `yaml:"permissions"`
	Abac            []AbacImpl           `yaml:"abac"`
	IsActor         bool                 `yaml:"-"`
	Actor           Actor                `yaml:"actor"`
	Aggregation     ObjectAggregationDef `yaml:"aggregation"`
	refObjects      map[string]Object
	refFields       []Field
	parentRefFields []Field
	childRefFields  []Field
	toManyRefFields []Field
	toOneRefFields  []Field
	fieldMap        map[string]Field
	httpMethodsMap  map[HttpMethod]bool
}

func (o *Object) HasCollection() bool {
	return o.Collection != nil
}

func (o *Object) HasCollectionType(t CollectionType) bool {
	for _, c := range o.Collection {
		if CollectionType(c.Type) == t {
			return true
		}
	}
	return false
}

func getCreatedUpdatedFields(registry Registry) []Field {
	var actorSpecificFields []Field
	for _, actorType := range registry.ListActorTypes() {
		actorSpecificFields = append(actorSpecificFields, Field{
			Name: "updatedBy" + utils.UCC(actorType),
			Type: "ActorTrace",
		})
	}
	return append([]Field{
		{Name: "created", Type: "ActorTrace"},
		{Name: "updated", Type: "ActorTrace"},
	}, actorSpecificFields...)
}

func getActorRoleField() Field {
	return Field{
		Name: "actorRoles",
		Type: "List<ActorRole>",
	}
}

func getDefaultIndexes() []Index {
	return []Index{
		{Name: "createdAt", Unique: false, Fields: []IndexField{{Name: "created.at", Order: -1}}},
		{Name: "updatedAt", Unique: false, Fields: []IndexField{{Name: "updated.at", Order: -1}}},
	}
}

func (o *Object) Validate(registry Registry) error {
	var err error
	if o.Name == "" {
		err = errors.Join(err, fmt.Errorf("object name is required"))
	}
	for _, c := range o.Collection {
		if collectionErr := c.Validate(registry); collectionErr != nil {
			err = errors.Join(err, collectionErr)
		}
	}

	if len(o.Collection) == 0 && len(o.HTTP.Methods) > 0 {
		err = errors.Join(err, fmt.Errorf("object %s cannot http methods without a db collection", o.Name))
	}

	if fieldErr := o.validateFields(registry); fieldErr != nil {
		err = errors.Join(err, fieldErr)
	}

	// If there are explicit permissions, grant the super role
	if len(o.Permissions.Read) > 0 && !utils.SliceContainsFunc(o.Permissions.Read, func(p ObjectPermissionsDef) bool {
		return p.Rbac == "Super"
	}) {
		o.Permissions.Read = append([]ObjectPermissionsDef{{
			Rbac: "Super",
		}}, o.Permissions.Read...)
	}
	// If there are explicit permissions, grant the super role
	if len(o.Permissions.Write) > 0 && !utils.SliceContainsFunc(o.Permissions.Write, func(p ObjectPermissionsDef) bool {
		return p.Rbac == "Super"
	}) {
		o.Permissions.Write = append([]ObjectPermissionsDef{{
			Rbac: "Super",
		}}, o.Permissions.Write...)
	}
	if permErr := o.validatePermissions(registry); permErr != nil {
		err = errors.Join(err, permErr)
	}

	if indexErr := o.validateIndexes(registry); indexErr != nil {
		err = errors.Join(err, indexErr)
	}

	if httpErr := o.validateHTTP(registry); httpErr != nil {
		err = errors.Join(err, httpErr)
	}

	if aggErr := o.validateAggregation(registry); aggErr != nil {
		err = errors.Join(err, aggErr)
	}

	if err != nil {
		return errors.Join(fmt.Errorf("%s has invalid spec", o.Name), err)
	}

	sort.Slice(o.Fields, func(i, j int) bool {
		return o.Fields[i].Name < o.Fields[j].Name
	})
	sort.Slice(o.Indexes, func(i, j int) bool {
		return o.Indexes[i].Name < o.Indexes[j].Name
	})
	sort.Slice(o.Abac, func(i, j int) bool {
		return o.Abac[i].Name < o.Abac[j].Name
	})
	sort.Slice(o.Permissions.Read, func(i, j int) bool {
		return o.Permissions.Read[i].Rbac < o.Permissions.Read[j].Rbac
	})
	sort.Slice(o.Permissions.Write, func(i, j int) bool {
		return o.Permissions.Write[i].Rbac < o.Permissions.Write[j].Rbac
	})
	sort.Slice(o.refFields, func(i, j int) bool {
		return o.refFields[i].Name < o.refFields[j].Name
	})
	sort.Slice(o.parentRefFields, func(i, j int) bool {
		return o.parentRefFields[i].Name < o.parentRefFields[j].Name
	})
	sort.Slice(o.childRefFields, func(i, j int) bool {
		return o.childRefFields[i].Name < o.childRefFields[j].Name
	})
	sort.Slice(o.toManyRefFields, func(i, j int) bool {
		return o.toManyRefFields[i].Name < o.toManyRefFields[j].Name
	})
	sort.Slice(o.toOneRefFields, func(i, j int) bool {
		return o.toOneRefFields[i].Name < o.toOneRefFields[j].Name
	})

	return nil
}

func (o *Object) buildFieldMap(registry Registry) error {
	if o.HasCollection() {
		o.Fields = append(o.Fields, getCreatedUpdatedFields(registry)...)
	}

	o.fieldMap = make(map[string]Field)
	var err error
	for _, f := range o.Fields {
		if _, ok := o.fieldMap[f.Name]; ok {
			err = errors.Join(err, fmt.Errorf("field %s is duplicated", f.Name))
		}
		o.fieldMap[f.Name] = f
	}
	return err
}

func (o *Object) appendActorFields() error {
	if !o.IsActor {
		return nil
	}
	actorRoleField := getActorRoleField()
	o.Fields = append(o.Fields, actorRoleField)
	o.fieldMap[actorRoleField.Name] = actorRoleField
	return nil
}

func (o *Object) validateFields(registry Registry) error {
	var err error
	for _, f := range o.Fields {
		if fieldErr := f.Validate(registry); fieldErr != nil {
			err = errors.Join(err, fmt.Errorf("field: %s - %v", f.Name, fieldErr))
		}
		if permErr := registry.ValidateObjectPermission(utils.UCC(o.Name)+utils.UCC(f.Name), f.Permissions); permErr != nil {
			err = errors.Join(err, fmt.Errorf("object %s has invalid permissions for field %s: %v", o.Name, f.Name, permErr))
		}
	}

	if permErr := registry.ValidateObjectPermission(utils.UCC(o.Name), o.Permissions); permErr != nil {
		err = errors.Join(permErr, fmt.Errorf("object %s has invalid permissions: %v", o.Name, permErr))
	}
	return err
}

func (o *Object) validateIndexes(registry Registry) error {
	if o.HasCollection() {
		o.Indexes = append(getDefaultIndexes(), o.Indexes...)
	}
	var err error
	for _, i := range o.Indexes {
		if indexErr := i.Validate(registry, o); indexErr != nil {
			err = errors.Join(err, indexErr)
		}
	}
	return err
}

func (o *Object) validatePermissions(registry Registry) error {
	var err error
	for _, p := range o.Abac {
		if _, ok := registry.GetAbacPermission(p.Name); !ok {
			err = errors.Join(err, fmt.Errorf("object %s has invalid abac permission implementation name %s", o.Name, p.Name))
		}
		if _, ok := o.GetField(p.Field); !ok {
			err = errors.Join(err, fmt.Errorf("object %s has invalid abac permission implementation field %s", o.Name, p.Field))
		}
	}
	if permErr := registry.ValidateObjectPermission(utils.UCC(o.Name), o.Permissions); err != nil {
		err = errors.Join(permErr, fmt.Errorf("object %s has invalid permissions %v", o.Name, err))
	}
	return err
}

func (o *Object) validateHTTP(_ Registry) error {
	o.httpMethodsMap = make(map[HttpMethod]bool)

	if o.HTTP.Endpoint == "" {
		o.HTTP.Endpoint = utils.LCC(o.Name)
	}

	for _, s := range o.HTTP.Methods {
		method, err := sanitizeHTTPMethod(s)
		if err != nil {
			return fmt.Errorf("object %s has invalid http method %s", o.Name, s)
		}
		o.httpMethodsMap[method] = true
	}
	return nil
}

func (o *Object) validateAggregation(registry Registry) error {
	// Auto-enable aggregations for objects with collections that have numeric fields
	if !o.Aggregation.Enabled && o.HasCollection() {
		o.autoPopulateAggregation(registry)
	}

	if !o.Aggregation.Enabled {
		return nil
	}

	if !o.HasCollection() {
		return fmt.Errorf("object %s cannot have aggregations without a db collection", o.Name)
	}

	if err := o.Aggregation.Validate(o, registry); err != nil {
		return fmt.Errorf("object %s has invalid aggregation config: %v", o.Name, err)
	}

	// Sort aggregation fields for deterministic output
	SortAggregationDef(&o.Aggregation)

	return nil
}

// autoPopulateAggregation automatically enables aggregation for objects with numeric fields
func (o *Object) autoPopulateAggregation(registry Registry) {
	var numericFields []AggregateFieldDef
	var groupByFields []string

	for _, field := range o.Fields {
		// Check if field is numeric
		if isNumericFieldType(field.Type) {
			numericFields = append(numericFields, AggregateFieldDef{
				Field:   field.Name,
				Methods: []string{"sum", "avg", "min", "max", "count"},
			})
		}

		// Add primitive fields and refs as potential groupBy fields
		if field.IsPrimitive() || field.IsRef() {
			groupByFields = append(groupByFields, field.Name)
		}
	}

	// Only enable if there are numeric fields to aggregate
	if len(numericFields) > 0 {
		o.Aggregation = ObjectAggregationDef{
			Enabled: true,
			Fields:  numericFields,
			GroupBy: groupByFields,
		}
	}
}

func (o *Object) HasAggregation() bool {
	return o.Aggregation.Enabled
}

func (o *Object) ListAggregateFields() []AggregateFieldDef {
	return o.Aggregation.ListAggregateFields()
}

func (o *Object) ListGroupByFields() []string {
	return o.Aggregation.ListGroupByFields()
}

func (o *Object) parseActor() error {
	if o.Actor.Id == "" {
		return nil
	}

	_, ok := o.GetField(o.Actor.Id)
	if !ok {
		return fmt.Errorf("object %s has actor id field %s, but the field is not defined", o.Name, o.Actor.Id)
	}

	o.IsActor = true
	o.Abac = append(o.Abac, AbacImpl{
		Name:  "ActorId",
		Field: o.Actor.Id,
	})

	if o.Actor.Username == "" {
		o.Actor.Username = "{id}"
	}
	if o.Actor.Name == "" {
		o.Actor.Name = o.Actor.Username
	}
	if o.Actor.AdminName == "" {
		o.Actor.AdminName = o.Actor.Username
	}

	o.Actor.LanguageParts = ParseActorNamePattern(o, o.Actor.Language)
	o.Actor.UsernameParts = ParseActorNamePattern(o, o.Actor.Username)
	o.Actor.NameParts = ParseActorNamePattern(o, o.Actor.Name)
	o.Actor.AdminNameParts = ParseActorNamePattern(o, o.Actor.AdminName)

	return nil
}

func (o *Object) ListFields() []Field {
	fields := o.Fields
	if o.HasCollection() {
		fields = append([]Field{
			{
				Name: "id",
				Type: fmt.Sprintf("Ref<%s>", o.Name),
			},
		}, fields...)
	}
	return fields
}

func (o *Object) ListImmutableFields(registry Registry) []Field {
	fields := make([]Field, 0)
	for _, f := range o.Fields {
		if f.Immutable {
			fields = append(fields, f)
		}
	}
	return fields
}

func (o *Object) ListEnumFields(registry Registry) []Field {
	fields := make([]Field, 0)
	for _, f := range o.Fields {
		rootType, _, _ := f.ResolveRootType(registry)
		if _, isEnum := registry.GetEnum(rootType); isEnum {
			fields = append(fields, f)
		}
	}
	return fields
}

func (o *Object) ListEnumFieldEnums(registry Registry) []Enum {
	enums := make([]Enum, 0)
	for _, f := range o.Fields {
		rootType, _, _ := f.ResolveRootType(registry)
		if enum, isEnum := registry.GetEnum(rootType); isEnum {
			if isEnum {
				enums = append(enums, enum)
			}
		}
	}
	return enums
}

func (o *Object) ListEnumListFields(registry Registry) []Field {
	fields := make([]Field, 0)
	for _, f := range o.Fields {
		if _, isList := f.ParseList(); isList {
			if _, isEnum := f.ParseEnum(); isEnum {
				fields = append(fields, f)
			}
		}
	}
	return fields
}

func (o *Object) ListPrimitiveFields(registry Registry) []Field {
	fields := make([]Field, 0)
	for _, f := range o.Fields {
		if _, rootFieldClass, _ := f.ResolveRootType(registry); rootFieldClass == RootFieldTypePrimitive {
			fields = append(fields, f)
		}
	}
	return fields
}

func (o *Object) ListObjectFields(registry Registry, omitRefs bool) []Object {
	var refs []Object
	refFound := make(map[string]struct{})
	refFound[o.Name] = struct{}{}
	fieldsToCheck := append(o.Fields, o.refFields...)
	for _, f := range fieldsToCheck {
		_, isRef := f.ParseRef()
		if omitRefs && isRef {
			continue
		}
		if rootFieldType, rootFieldClass, _ := f.ResolveRootType(registry); rootFieldClass == RootFieldTypeObject {
			if _, ok := refFound[rootFieldType]; ok {
				continue
			}
			ref, ok := registry.Get(rootFieldType)
			if ok {
				refs = append(refs, ref)
				refFound[rootFieldType] = struct{}{}
			}
		}
	}
	return refs
}

func (o *Object) ListObjectListFields(registry Registry, omitRefs bool) []Object {
	var refs []Object
	refFound := make(map[string]struct{})
	refFound[o.Name] = struct{}{}
	fieldsToCheck := append(o.Fields, o.refFields...)
	for _, f := range fieldsToCheck {
		_, isList := f.ParseList()
		if !isList {
			continue
		}
		_, isRef := f.ParseRef()
		if omitRefs && isRef {
			continue
		}
		if rootFieldType, rootFieldClass, _ := f.ResolveRootType(registry); rootFieldClass == RootFieldTypeObject {
			if _, ok := refFound[rootFieldType]; ok {
				continue
			}
			ref, ok := registry.Get(rootFieldType)
			if ok {
				refs = append(refs, ref)
				refFound[rootFieldType] = struct{}{}
			}
		}
	}
	return refs
}

func (o *Object) ListRefObjects() []Object {
	var refs []Object
	for _, refObj := range o.refObjects {
		refs = append(refs, refObj)
	}
	return refs
}

func (o *Object) ListRefFields() []Field {
	return o.refFields
}

func (o *Object) ListToManyRefFields() []Field {
	return o.toManyRefFields
}

func (o *Object) ListToOneRefFields() []Field {
	return o.toOneRefFields
}

func (o *Object) ListParentRefFields() []Field {
	return o.parentRefFields
}

func (o *Object) ListChildRefFields() []Field {
	return o.childRefFields
}

func (o *Object) HasHTTPMethods() bool {
	return len(o.httpMethodsMap) > 0
}

func (o *Object) HasHTTPMethod(method HttpMethod) bool {
	return o.httpMethodsMap[method]
}

func (o *Object) HasAtLeastOneHTTPMethod(methods ...HttpMethod) bool {
	for _, m := range methods {
		if o.HasHTTPMethod(m) {
			return true
		}
	}
	return false
}

func (o *Object) BuildRefs(registry Registry) error {
	if o.refObjects == nil {
		o.refObjects = make(map[string]Object)
	}
	for i, field := range o.Fields {
		if elemName, isArray := field.ParseList(); isArray {
			field.Type = elemName
		}
		if valName, isKeyVal := field.ParseKeyVal(); isKeyVal {
			field.Type = valName
		}
		if refObjName, isRef := field.ParseRef(); isRef {
			refObj, ok := registry.Get(refObjName)
			if refObj.Name == o.Name {
				continue
			}
			if !ok {
				return fmt.Errorf("object %s has ref %s, but %s was not registered", o.Name, field.Name, refObjName)
			}
			fmt.Println("Adding ref", refObj.Name, "to", o.Name)
			o.refObjects[refObj.Name] = refObj
			o.refFields = append(o.refFields, o.Fields[i])
			o.toOneRefFields = append(o.toOneRefFields, o.Fields[i])
			if err := refObj.acceptRef(o.Fields[i], *o, registry); err != nil {
				return errors.Join(err, fmt.Errorf("object %s has invalid ref %s", o.Name, refObj.Name))
			}
			_, isParentRef := field.ParseParentRef()
			if isParentRef {
				o.parentRefFields = append(o.parentRefFields, o.Fields[i])
			}
		}
	}
	return nil
}

func (o *Object) acceptRef(refField Field, refObj Object, registry Registry) error {
	if o.refObjects == nil {
		o.refObjects = make(map[string]Object)
	}
	if _, ok := o.refObjects[refObj.Name]; ok {
		return fmt.Errorf("object %s already has ref %s", o.Name, refObj.Name)
	}
	o.refObjects[refObj.Name] = refObj
	field := Field{Name: refObj.Name + "s", Type: fmt.Sprintf("List<Ref<%s>>", refObj.Name)}
	o.refFields = append(o.refFields, field)
	o.toManyRefFields = append(o.toManyRefFields, field)

	_, isChildRef := refField.ParseParentRef()
	if isChildRef {
		o.childRefFields = append(o.childRefFields, field)
	}

	if err := registry.OverwriteObject(*o); err != nil {
		return errors.Join(err, fmt.Errorf("error overiting object %s with ref to %s", refObj.Name, o.Name))
	}

	return nil
}

func (o *Object) GetField(name string) (Field, bool) {
	if name == "id" {
		return Field{
			Name: "id",
			Type: fmt.Sprintf("Ref<%s>", o.Name),
		}, true
	}
	field, ok := o.fieldMap[name]
	return field, ok
}

func (o *Object) ApplyExtendedPermissions(registry Registry) error {
	if len(o.Permissions.Read) > 0 {
		if err := o.applyExtendedPermissionsToPermissionDef(registry, "read", nil, &o.Permissions.Read); err != nil {
			return err
		}
	}
	if len(o.Permissions.Write) > 0 {
		if err := o.applyExtendedPermissionsToPermissionDef(registry, "write", nil,
			&o.Permissions.Write); err != nil {
			return err
		}
	}
	for fieldIdx := range o.Fields {
		field := &o.Fields[fieldIdx]
		if len(field.Permissions.Read) > 0 {
			if err := o.applyExtendedPermissionsToPermissionDef(registry, "read", field, &field.Permissions.Read); err != nil {
				return err
			}
		}
		if len(field.Permissions.Write) > 0 {
			if err := o.applyExtendedPermissionsToPermissionDef(registry, "write", field, &field.Permissions.Write); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *Object) applyExtendedPermissionsToPermissionDef(
	registry Registry,
	permType string,
	field *Field,
	fieldPermissions *[]ObjectPermissionsDef,
) error {
	for _, readPerm := range *fieldPermissions {
		if readPerm.Rbac == "" {
			continue
		}
		extendedRoles := registry.GetExtendedRoles(readPerm.Rbac)
		for _, extendedRole := range extendedRoles {
			label := o.Name
			if field != nil {
				label = o.Name + "." + field.Name
			}
			fmt.Println(label, "- extending", readPerm.Rbac, permType, "permissions to role", extendedRole)
			isRoleAlreadyExists := false
			for _, permCheck := range *fieldPermissions {
				if permCheck.Rbac == extendedRole {
					isRoleAlreadyExists = true
					break
				}
			}
			if isRoleAlreadyExists {
				fmt.Println(" -> skipping", extendedRole, "already has explicit", permType, "permission on", label)
				continue
			}
			readPermCopy := readPerm
			readPermCopy.Rbac = extendedRole
			(*fieldPermissions) = append(*fieldPermissions, readPermCopy)
		}
	}
	return nil
}
