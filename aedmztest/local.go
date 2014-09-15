// Copyright 2013 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// +build !appengine

package aedmztest

import (
	"bytes"
	"code.google.com/p/leveldb-go/leveldb/memdb"
	"github.com/maruel/aedmz"
	"github.com/maruel/ofh"
	"net/http"
	"time"
)

// AppContextImplMock implements fake specific properties. It can be passed to
// NewAppMock to create an aedmz.AppContext.
type AppContextImplMock struct {
	Timestamp time.Time
}

// Now simulates an hardcoded time constant. If AppContextImplMock.Timestamp is
// not set, the current time in UTC is returned.
func (a AppContextImplMock) Now() time.Time {
	if a.Timestamp.IsZero() {
		return time.Now().UTC()
	}
	return a.Timestamp
}

// NewAppMock returns an aedmz.AppContext to be used in unit tests.
//
// It has AppID "Yo" and version "v1".
func NewAppMock(a aedmz.AppContextImpl) aedmz.AppContext {
	if a == nil {
		a = &AppContextImplMock{}
	}
	return aedmz.NewAppInternal("Yo", "v1", &bytes.Buffer{}, ofh.MakeStubProvider(http.DefaultClient), memdb.New(nil), a)
}

// CloseRequest closes a testing aedmz.RequestContext.
//
// It is only necessary to call it when creating a RequestContext with
// AppContext.NewContext() directly without using a route where InjectContext()
// was not called on. InjectContext() cleans up the context on the user's
// behalf.
func CloseRequest(r aedmz.RequestContext) {
	// Closing is not necessary when running standalone.
}
