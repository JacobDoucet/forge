package types

import "fmt"

type CollectionRegistry interface {
	Get(name string) (Collection, bool)
	Register(coll Collection) error
	List() []Collection
}

func newCollectionRegistry() CollectionRegistry {
	return &collectionRegistry{
		collections: make(map[string]Collection),
	}
}

type collectionRegistry struct {
	collections map[string]Collection
}

func (r *collectionRegistry) Get(name string) (Collection, bool) {
	coll, ok := r.collections[name]
	return coll, ok
}

func (r *collectionRegistry) Register(coll Collection) error {
	if _, ok := r.collections[coll.Name]; ok {
		return fmt.Errorf("collection %s already registered", coll.Name)
	}
	r.collections[coll.Name] = coll
	return nil
}

func (r *collectionRegistry) List() []Collection {
	var collections []Collection
	for _, coll := range r.collections {
		collections = append(collections, coll)
	}
	return collections
}
