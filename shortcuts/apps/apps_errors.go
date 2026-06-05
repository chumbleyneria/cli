// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package apps

import (
	"errors"

	"github.com/larksuite/cli/errs"
	"github.com/larksuite/cli/extension/fileio"
	"github.com/larksuite/cli/internal/client"
)

func appsValidationError(format string, args ...any) *errs.ValidationError {
	return errs.NewValidationError(errs.SubtypeInvalidArgument, format, args...)
}

func appsValidationParamError(param, format string, args ...any) *errs.ValidationError {
	return appsValidationError(format, args...).WithParam(param)
}

func appsInvalidParam(name, reason string) errs.InvalidParam {
	return errs.InvalidParam{Name: name, Reason: reason}
}

func appsFailedPreconditionParamError(param, format string, args ...any) *errs.ValidationError {
	return errs.NewValidationError(errs.SubtypeFailedPrecondition, format, args...).WithParam(param)
}

func appsInputPathError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, fileio.ErrPathValidation) {
		return appsValidationParamError("--path", "unsafe --path: %s", err).WithCause(err)
	}
	return appsValidationParamError("--path", "cannot read --path: %s", err).WithCause(err)
}

func appsInputPathEntryError(path string, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, fileio.ErrPathValidation) {
		return appsValidationParamError("--path", "unsafe --path entry %s: %s", path, err).WithCause(err)
	}
	return appsValidationParamError("--path", "cannot read --path entry %s: %s", path, err).WithCause(err)
}

func appsFileIOError(err error, format string, args ...any) *errs.InternalError {
	return errs.NewInternalError(errs.SubtypeFileIO, format, args...).WithCause(err)
}

func appsAPIBoundaryError(err error) error {
	return client.WrapDoAPIError(err)
}

func enrichHTMLPublishAPIError(err error) error {
	if err == nil {
		return nil
	}
	p, ok := errs.ProblemOf(err)
	if !ok {
		return appsAPIBoundaryError(err)
	}
	if p.Message != "" {
		p.Message = "html-publish failed: " + p.Message
	}
	if hint := buildHTMLPublishFailureHint(p.Code); hint != "" {
		p.Hint = hint
	}
	return err
}
