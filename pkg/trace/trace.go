package trace

import "context"

type FinishFunc func()

type Trace interface {
	TraceFunc(c context.Context, kvs ...Entry) (context.Context, FinishFunc)
	TraceFuncNamed(c context.Context, name string, kvs ...Entry) (context.Context, FinishFunc)
}
