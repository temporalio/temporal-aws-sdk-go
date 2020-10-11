package sessionfactory

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
)

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

