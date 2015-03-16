// Copyright 2013 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// +build appengine

package aedmztest

import (
	"io"
	"net/http"

	"github.com/maruel/aedmz"

	"appengine"
	"appengine/aetest"
)

// AppContextImplMock implements fake specific properties. It can be passed to
// NewAppMock to create an aedmz.AppContext.
type AppContextImplMock struct{}

// NewContext is called through aedmz.AppContext on each new HTTP request.
//
// It creates all the fake necessary to simulate an HTTP context for logging,
// DB queries, user context, etc.
func (a AppContextImplMock) NewContext(r *http.Request) appengine.Context {
	// https://developers.google.com/appengine/docs/go/tools/localunittesting
	//
	// The call to aetest.NewContext will start dev_appserver.py in a subprocess,
	// which will be used to service API calls during the test. This subprocess
	// will be shutdown with the call to Close.
	//
	// This means it's costly.
	//
	// aetest.NewContext() won't use by default the 'application' specified in
	// app.yaml. NewAppMock() prefills appContext.appID so appengine.AppID() is
	// never called in RequestContext.AppID(). If this becomes an issue, pass
	// aetest.Options{AppID:"Yo"} instead of nil.
	c, err := aetest.NewContext(nil)
	if err != nil {
		panic(err)
	}
	return c
}

// NewAppMock returns an AppContext to be used in unit tests.
//
// It has AppID "Yo" and version "v1".
func NewAppMock(a aedmz.AppContextImpl) aedmz.AppContext {
	if a == nil {
		a = &AppContextImplMock{}
	}
	return aedmz.NewAppInternal("Yo", "v1", a)
}

// CloseRequest closes a testing aedmz.RequestContext.
//
// It is only necessary to call it when creating a RequestContext with
// AppContext.NewContext() directly without using a route where InjectContext()
// was not called on. InjectContext() cleans up the context on the user's
// behalf.
func CloseRequest(r aedmz.RequestContext) {
	c, _ := r.(aedmz.RequestContextAppengine).AppengineContext().(io.Closer)
	c.Close()
}
