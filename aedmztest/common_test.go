// Copyright 2013 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

package aedmztest

import (
	"github.com/maruel/ut"
	"net/http"
	"testing"
)

func TestFunctional(t *testing.T) {
	app := NewAppMock(nil)
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	ut.AssertEqual(t, nil, err)
	c := app.NewContext(req)
	defer CloseRequest(c)

	ut.AssertEqual(t, "Yo", c.AppID())
	ut.AssertEqual(t, "v1", c.AppVersion())
}
