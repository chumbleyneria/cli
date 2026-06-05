// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package errclass

import "github.com/larksuite/cli/errs"

// appsCodeMeta holds Miaoda apps-service Lark code -> CodeMeta mappings.
// Only stable endpoint semantics are registered globally; endpoint-specific
// recovery hints stay in shortcuts/apps.
var appsCodeMeta = map[int]CodeMeta{
	90002: {Category: errs.CategoryAPI, Subtype: errs.SubtypeNotFound}, // app_id unknown or caller lacks access
}

func init() { mergeCodeMeta(appsCodeMeta, "apps") }
