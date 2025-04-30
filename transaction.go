package mongodbr

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RunTransactionOptions struct {
	ClientKey string
	// 单次事务最大超时时间
	Timeout time.Duration
	// 最大重试次数
	MaxRetry int

	withSessionOptions     []*options.SessionOptions
	withTransactionOptions []*options.TransactionOptions
}

type RunTransactionOption func(options *RunTransactionOptions)

// 使用指定的client key
func RunTransactionOptionWithClientKey(clientKey string) func(options *RunTransactionOptions) {
	return func(options *RunTransactionOptions) {
		options.ClientKey = clientKey
	}
}

// 修改SessionOptions参数
func RunTransactionOptionWithSessionOptions(sOptions ...*options.SessionOptions) func(options *RunTransactionOptions) {
	return func(runOptions *RunTransactionOptions) {
		runOptions.withSessionOptions = append(runOptions.withSessionOptions, sOptions...)
	}
}

// 修改TransactionOptions参数
func RunTransactionOptionWithTransactionOptions(tOptions ...*options.TransactionOptions) func(options *RunTransactionOptions) {
	return func(runOptions *RunTransactionOptions) {
		runOptions.withTransactionOptions = append(runOptions.withTransactionOptions, tOptions...)
	}
}

// 修改TransactionOptions参数
func RunTransactionOptionWithTimeout(timeout time.Duration) func(options *RunTransactionOptions) {
	return func(runOptions *RunTransactionOptions) {
		runOptions.Timeout = timeout
	}
}

// 修改TransactionOptions的最大重试次数，默认为3次
func RunTransactionOptionWithMaxRetry(maxRetry int) func(options *RunTransactionOptions) {
	return func(runOptions *RunTransactionOptions) {
		runOptions.MaxRetry = maxRetry
	}
}

func newDefaultRunTransactionOptions() *RunTransactionOptions {
	return &RunTransactionOptions{
		ClientKey: DefaultAlias,
		Timeout:   30 * time.Second,
		MaxRetry:  3,
	}
}

func RunTransaction(fn func(mongo.SessionContext) error, opts ...RunTransactionOption) error {
	return RunTransactionWithContext(context.Background(), fn, opts...)
}

// run fn with mongodb transaction
func RunTransactionWithContext(ctx context.Context,
	fn func(mongo.SessionContext) error,
	opts ...RunTransactionOption) error {

	transactionOptions := newDefaultRunTransactionOptions()
	for _, eachOpt := range opts {
		eachOpt(transactionOptions)
	}
	maxRetry := transactionOptions.MaxRetry
	if maxRetry <= 0 {
		maxRetry = 3
	}
	// client
	c := GetClient(transactionOptions.ClientKey)

	// start session
	session, err := c.StartSession(transactionOptions.withSessionOptions...)
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	for attempt := 0; attempt < transactionOptions.MaxRetry; attempt++ {
		// 每次事务执行新建一个超时ctx
		txnCtx, cancel := context.WithTimeout(ctx, transactionOptions.Timeout)
		defer cancel()

		// perform operation
		err = mongo.WithSession(txnCtx, session, func(sc mongo.SessionContext) error {
			// start transaction
			if err := session.StartTransaction(transactionOptions.withTransactionOptions...); err != nil {
				return err
			}

			innerErr := fn(sc)
			if innerErr != nil {
				// roolback
				_ = session.AbortTransaction(sc)
				return innerErr
			}
			if err := sc.CommitTransaction(sc); err != nil {
				return err
			}
			return nil
		})

		if err == nil {
			return nil
		}

		// 如果是瞬时事务错误，允许重试
		if isTransientTransactionError(err) {
			continue
		}

		return err
	}

	return fmt.Errorf("RunTransactionWithContext: transaction failed after %d attempts, last error: %w", maxRetry, err)
}

func RunTransactionWithResult[T any](fn func(mongo.SessionContext) (T, error), opts ...RunTransactionOption) (T, error) {
	return RunTransactionWithResultWithContext(context.Background(), fn, opts...)
}

// run fn with mongodb transaction
func RunTransactionWithResultWithContext[T any](ctx context.Context,
	fn func(mongo.SessionContext) (T, error),
	opts ...RunTransactionOption) (T, error) {

	var zero T
	transactionOptions := newDefaultRunTransactionOptions()
	for _, eachOpt := range opts {
		eachOpt(transactionOptions)
	}
	maxRetry := transactionOptions.MaxRetry
	if maxRetry <= 0 {
		maxRetry = 3
	}
	// client
	c := GetClient(transactionOptions.ClientKey)

	// start session
	session, err := c.StartSession(transactionOptions.withSessionOptions...)
	if err != nil {
		return zero, err
	}
	defer session.EndSession(context.Background())

	for attempt := 0; attempt < transactionOptions.MaxRetry; attempt++ {
		// 每次事务执行新建一个超时ctx
		txnCtx, cancel := context.WithTimeout(ctx, transactionOptions.Timeout)
		defer cancel()

		var result T

		// perform operation
		err = mongo.WithSession(txnCtx, session, func(sc mongo.SessionContext) error {
			// start transaction
			if err := session.StartTransaction(transactionOptions.withTransactionOptions...); err != nil {
				return err
			}

			var innerErr error
			result, innerErr = fn(sc)
			if innerErr != nil {
				// roolback
				_ = session.AbortTransaction(sc)
				return innerErr
			}
			if err := sc.CommitTransaction(sc); err != nil {
				return err
			}
			return nil
		})

		if err == nil {
			return result, nil
		}

		// 如果是瞬时事务错误，允许重试
		if isTransientTransactionError(err) {
			continue
		}

		return zero, err
	}

	return zero, fmt.Errorf("RunTransactionWithResultWithContext: transaction failed after %d attempts, last error: %w", maxRetry, err)
}

// isTransientTransactionError 判断是否是瞬时事务错误
func isTransientTransactionError(err error) bool {
	var cmdErr mongo.CommandError
	if errors.As(err, &cmdErr) {
		for _, label := range cmdErr.Labels {
			if label == "TransientTransactionError" {
				return true
			}
		}
	}

	var writeErr mongo.WriteException
	if errors.As(err, &writeErr) {
		for _, label := range writeErr.Labels {
			if label == "TransientTransactionError" {
				return true
			}
		}
	}

	return false
}
