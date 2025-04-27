package mongodbr

import "go.mongodb.org/mongo-driver/mongo/options"

// aggregate
func (r *RepositoryBase) Aggregate(pipeline interface{}, dataList interface{}, opts ...MongodbrAggregateOption) (err error) {
	aOptions := &MongodbrAggregateOptions{
		AggregateOptions: options.Aggregate(),
	}
	for _, o := range opts {
		o(aOptions)
	}
	ctx, cancel := CreateContextWith(r.configuration, aOptions.WithCtx)
	defer cancel()

	cur, err := r.collection.Aggregate(ctx, pipeline, aOptions.AggregateOptions)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	return cur.All(ctx, dataList)
}
