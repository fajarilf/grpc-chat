package utils

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapGRPCError(code codes.Code, message string, source error) error {
	if source == nil {
		return nil
	}
	return status.Errorf(code, "%s: %v", message, source)
}

func WrapError(message string, source error) error {
	if source == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, source)
}
