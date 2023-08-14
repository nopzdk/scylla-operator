// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStorageServiceSstablesByKeyspacePostParams creates a new StorageServiceSstablesByKeyspacePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceSstablesByKeyspacePostParams() *StorageServiceSstablesByKeyspacePostParams {
	return &StorageServiceSstablesByKeyspacePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceSstablesByKeyspacePostParamsWithTimeout creates a new StorageServiceSstablesByKeyspacePostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceSstablesByKeyspacePostParamsWithTimeout(timeout time.Duration) *StorageServiceSstablesByKeyspacePostParams {
	return &StorageServiceSstablesByKeyspacePostParams{
		timeout: timeout,
	}
}

// NewStorageServiceSstablesByKeyspacePostParamsWithContext creates a new StorageServiceSstablesByKeyspacePostParams object
// with the ability to set a context for a request.
func NewStorageServiceSstablesByKeyspacePostParamsWithContext(ctx context.Context) *StorageServiceSstablesByKeyspacePostParams {
	return &StorageServiceSstablesByKeyspacePostParams{
		Context: ctx,
	}
}

// NewStorageServiceSstablesByKeyspacePostParamsWithHTTPClient creates a new StorageServiceSstablesByKeyspacePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceSstablesByKeyspacePostParamsWithHTTPClient(client *http.Client) *StorageServiceSstablesByKeyspacePostParams {
	return &StorageServiceSstablesByKeyspacePostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceSstablesByKeyspacePostParams contains all the parameters to send to the API endpoint

	for the storage service sstables by keyspace post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceSstablesByKeyspacePostParams struct {

	/* Cf.

	   Column family name
	*/
	Cf string

	/* Keyspace.

	   The keyspace
	*/
	Keyspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service sstables by keyspace post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceSstablesByKeyspacePostParams) WithDefaults() *StorageServiceSstablesByKeyspacePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service sstables by keyspace post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceSstablesByKeyspacePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) WithTimeout(timeout time.Duration) *StorageServiceSstablesByKeyspacePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) WithContext(ctx context.Context) *StorageServiceSstablesByKeyspacePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) WithHTTPClient(client *http.Client) *StorageServiceSstablesByKeyspacePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) WithCf(cf string) *StorageServiceSstablesByKeyspacePostParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) SetCf(cf string) {
	o.Cf = cf
}

// WithKeyspace adds the keyspace to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) WithKeyspace(keyspace string) *StorageServiceSstablesByKeyspacePostParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service sstables by keyspace post params
func (o *StorageServiceSstablesByKeyspacePostParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceSstablesByKeyspacePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cf
	qrCf := o.Cf
	qCf := qrCf
	if qCf != "" {

		if err := r.SetQueryParam("cf", qCf); err != nil {
			return err
		}
	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}