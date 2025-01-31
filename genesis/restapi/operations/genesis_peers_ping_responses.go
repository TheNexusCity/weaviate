//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GenesisPeersPingOKCode is the HTTP code returned for type GenesisPeersPingOK
const GenesisPeersPingOKCode int = 200

/*GenesisPeersPingOK Ping received

swagger:response genesisPeersPingOK
*/
type GenesisPeersPingOK struct {
}

// NewGenesisPeersPingOK creates GenesisPeersPingOK with default headers values
func NewGenesisPeersPingOK() *GenesisPeersPingOK {
	return &GenesisPeersPingOK{}
}

// WriteResponse to the client
func (o *GenesisPeersPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GenesisPeersPingUnauthorizedCode is the HTTP code returned for type GenesisPeersPingUnauthorized
const GenesisPeersPingUnauthorizedCode int = 401

/*GenesisPeersPingUnauthorized Unauthorized or invalid credentials.

swagger:response genesisPeersPingUnauthorized
*/
type GenesisPeersPingUnauthorized struct {
}

// NewGenesisPeersPingUnauthorized creates GenesisPeersPingUnauthorized with default headers values
func NewGenesisPeersPingUnauthorized() *GenesisPeersPingUnauthorized {
	return &GenesisPeersPingUnauthorized{}
}

// WriteResponse to the client
func (o *GenesisPeersPingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GenesisPeersPingForbiddenCode is the HTTP code returned for type GenesisPeersPingForbidden
const GenesisPeersPingForbiddenCode int = 403

/*GenesisPeersPingForbidden The used API-key has insufficient permissions.

swagger:response genesisPeersPingForbidden
*/
type GenesisPeersPingForbidden struct {
}

// NewGenesisPeersPingForbidden creates GenesisPeersPingForbidden with default headers values
func NewGenesisPeersPingForbidden() *GenesisPeersPingForbidden {
	return &GenesisPeersPingForbidden{}
}

// WriteResponse to the client
func (o *GenesisPeersPingForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GenesisPeersPingNotFoundCode is the HTTP code returned for type GenesisPeersPingNotFound
const GenesisPeersPingNotFoundCode int = 404

/*GenesisPeersPingNotFound Successful query result but no such peer was found.

swagger:response genesisPeersPingNotFound
*/
type GenesisPeersPingNotFound struct {
}

// NewGenesisPeersPingNotFound creates GenesisPeersPingNotFound with default headers values
func NewGenesisPeersPingNotFound() *GenesisPeersPingNotFound {
	return &GenesisPeersPingNotFound{}
}

// WriteResponse to the client
func (o *GenesisPeersPingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
