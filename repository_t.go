package mongodbr

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// find one T by _id
func FindTByObjectId[T any](repository IRepository, id primitive.ObjectID, opts ...MongodbrFindOneOption) (*T, error) {
	result := new(T)
	err := repository.FindOneByObjectId(id, result, opts...)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// find one T by filter
// filter: filter
func FindOneTByFilter[T any](repository IRepository, filter interface{}, opts ...MongodbrFindOneOption) (*T, error) {
	result := new(T)
	err := repository.FindOne(filter, result, opts...)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// find all t
func FindAllT[T any](repository IRepository, opts ...MongodbrFindOption) ([]*T, error) {
	list := make([]*T, 0)
	err := repository.FindAll(&list, opts...)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// find list T by filter
func FindTByFilter[T any](repository IRepository, filter interface{}, opts ...MongodbrFindOption) ([]*T, error) {
	list := make([]*T, 0)
	err := repository.FindListByFilter(filter, &list, opts...)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// find list T by _id list
func FindTListByObjectIdList[T any](repository IRepository, idList []primitive.ObjectID, opts ...MongodbrFindOption) ([]*T, error) {
	list := make([]*T, 0)
	err := repository.FindListByObjectIdList(idList, &list, opts...)
	if err != nil {
		return nil, err
	}
	return list, nil
}
