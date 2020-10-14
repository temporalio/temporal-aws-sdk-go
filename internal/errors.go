package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"go.temporal.io/sdk/temporal"
)

// errorType identifies a custom Cadence error as an AWS SDK error.
const errorType = "AWS_SDK_ERROR"

// encodedError is a serialized version of the AWS SDK error.
type encodedError struct {
	Code    string
	Message string
	OrigErr string
}

// Error implements the error interface.
func (e encodedError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// EncodeError encodes an error returned by the AWS SDK as a Cadence error.
func EncodeError(err error) error {
	if err == nil {
		return nil
	}

	if aerr, ok := err.(awserr.Error); ok {
		newErr := encodedError{
			Code:    aerr.Code(),
			Message: aerr.Message(),
		}

		// TODO: try to encode the original error as well?
		if aerr.OrigErr() != nil {
			newErr.OrigErr = aerr.OrigErr().Error()
		}

		return temporal.NewApplicationErrorWithCause(newErr.Error(), "AWS_SDK_" + aerr.Code(), aerr.OrigErr(), newErr)
	}

	return err
}

// decodedError is a serialized version of the AWS SDK error.
type decodedError struct {
	Code    string
	Message string
	OrigErr error
}

// decodedError implements the error interface.
func (e decodedError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// DecodeError decodes an error returned by a Cadence activity.
// If the error is an AWS SDK error, it gets converted to an AWS SDK compatible error interface.
// Otherwise the original error is returned.
func DecodeError(err error) error {
	if err == nil {
		return nil
	}

	var applicationError *temporal.ApplicationError

	if errors.As(err, &applicationError) && strings.HasPrefix(applicationError.Type(), "AWS_SDK_") {
		var aerr decodedError

		_ = applicationError.Details(&aerr)

		return awserr.New(aerr.Code, aerr.Message, aerr.OrigErr)
	}

	return err
}
