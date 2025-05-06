package utils

import (
	"context"
	"project_structure/internal/core/constants"
)

func GetReqID(ctx context.Context) string {
	id := ctx.Value(constants.REQUEST_KEY)
	if id != nil {
		requestID, ok := id.(string)
		if ok {
			return requestID
		}
	}

	// log.Warn().Msg("Request ID not found in context")
	return ""
}
