package middleware

import (
	"context"
)

// Stack provides protocol and transport agnostic set of middleware split into
// distinct steps. Steps have specific transitions between them, that is
// managed by the individual step.
//
// Steps are composed as middleware around the underlying handler in the
// following order:
//
//   Initialize -> Serialize -> Build -> Finalize -> Deserialize -> Handler
//
// Any middleware within the chain may chose to stop and return an error or
// response. Since the middleware decorate the handler like a call stack, each
// middleware will receive the result of the next middleware in the chain.
// Middleware that does not need to react to an input, or result must forward
// along the input down the chain, or return the result back up the chain.
//
//   Initialize <- Serialize -> Build -> Finalize <- Deserialize <- Handler
type Stack struct {
	// Initialize Prepares the input, and sets any default parameters as
	// needed, (e.g. idempotency token, and presigned URLs).
	//
	// Takes Input Parameters, and returns result or error.
	//
	// Receives result or error from Serialize step.
	Initialize *InitializeStep

	// Serializes the prepared input into a data structure that can be consumed
	// by the target transport's message, (e.g. REST-JSON serialization)
	//
	// Converts Input Parameters into a Request, and returns the result or error.
	//
	// Receives result or error from Build step.
	Serialize *SerializeStep

	// Adds additional metadata to the serialized transport message,
	// (e.g. HTTP's Content-Length header, or body checksum). Decorations and
	// modifications to the message should be copied to all message attempts.
	//
	// Takes Request, and returns result or error.
	//
	// Receives result or error from Finalize step.
	Build *BuildStep

	// Preforms final preparations needed before sending the message. The
	// message should already be complete by this stage, and is only alternated
	// to meet the expectations of the recipient, (e.g. Retry and AWS SigV4
	// request signing)
	//
	// Takes Request, and returns result or error.
	//
	// Receives result or error from Deserialize step.
	Finalize *FinalizeStep

	// Reacts to the handler's response returned by the recipient of the request
	// message. Deserializes the response into a structured type or error above
	// stacks can react to.
	//
	// Should only forward Request to underlying handler.
	//
	// Takes Request, and returns result or error.
	//
	// Receives raw response, or error from underlying handler.
	Deserialize *DeserializeStep

	id string
}

// NewStack returns an initialize empty stack.
func NewStack(id string, newRequestFn func() interface{}) *Stack {
	return &Stack{
		id:          id,
		Initialize:  NewInitializeStep(),
		Serialize:   NewSerializeStep(newRequestFn),
		Build:       NewBuildStep(),
		Finalize:    NewFinalizeStep(),
		Deserialize: NewDeserializeStep(),
	}
}

// ID returns the unique ID for the stack as a middleware.
func (s *Stack) ID() string { return s.id }

// HandleMiddleware invokes the middleware stack decorating the next handler.
// Each step of stack will be invoked in order before calling the next step.
// With the next handler call last.
//
// The input value must be the input parameters of the operation being
// performed.
//
// Will return the result of the operation, or error.
func (s *Stack) HandleMiddleware(ctx context.Context, input interface{}, next Handler) (
	output interface{}, metadata Metadata, err error,
) {
	h := DecorateHandler(next,
		s.Initialize,
		s.Serialize,
		s.Build,
		s.Finalize,
		s.Deserialize,
	)

	return h.Handle(ctx, input)
}
