package mongodbr

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// find one T by _id
func FindTByObjectId[T any](repository IRepository, id primitive.ObjectID, opts ...FindOneOption) (*T, error) {
	res := repository.FindByObjectId(id, opts...)
	result := new(T)
	if err := res.One(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// find one T by filter
// filter: filter
func FindOneTByFilter[T any](repository IRepository, filter interface{}, opts ...FindOneOption) (*T, error) {
	res := repository.FindOne(filter, opts...)
	result := new(T)
	if err := res.One(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// find all t
func FindAllT[T any](repository IRepository, opts ...FindOption) ([]*T, error) {
	res := repository.FindAll()
	list := make([]*T, 0)
	if err := res.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

// find list T by filter
func FindTByFilter[T any](repository IRepository, filter interface{}, opts ...FindOption) ([]*T, error) {
	res := repository.FindByFilter(filter, opts...)
	list := make([]*T, 0)
	if err := res.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}

// find list T by _id list
func FindTListByObjectIdList[T any](repository IRepository, idList []primitive.ObjectID) ([]*T, error) {
	res := repository.FindListByObjectIdList(idList)
	list := make([]*T, 0)
	if err := res.All(&list); err != nil {
		return nil, err
	}
	return list, nil
}
