// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package event

import "github.com/larksuite/cli/errs"

func eventValidationError(param string, format string, args ...any) *errs.ValidationError {
	err := errs.NewValidationError(errs.SubtypeInvalidArgument, format, args...)
	if param != "" {
		err = err.WithParam(param)
	}
	return err
}

func eventValidationErrorWithCause(param string, cause error, format string, args ...any) *errs.ValidationError {
	return eventValidationError(param, format, args...).WithCause(cause)
}

func eventFileIOError(cause error, format string, args ...any) *errs.InternalError {
	return errs.NewInternalError(errs.SubtypeFileIO, format, args...).WithCause(cause)
}

func eventNetworkError(cause error, format string, args ...any) *errs.NetworkError {
	return errs.NewNetworkError(errs.SubtypeNetworkTransport, format, args...).WithCause(cause)
}
