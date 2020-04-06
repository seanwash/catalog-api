package graph

import (
	"context"

	"github.com/seanwash/catalog-api/api/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ensureUserIsAdmin(ctx context.Context) error {
	if middleware.UserFromContext(ctx) == nil {
		return gqlerror.Errorf("Unauthorized")
	} else {
		return nil
	}
}
