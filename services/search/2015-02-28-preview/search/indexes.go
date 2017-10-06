package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// IndexesClient is the search Client
type IndexesClient struct {
	ManagementClient
}

// NewIndexesClient creates an instance of the IndexesClient client.
func NewIndexesClient() IndexesClient {
	return NewIndexesClientWithBaseURI(DefaultBaseURI)
}

// NewIndexesClientWithBaseURI creates an instance of the IndexesClient client.
func NewIndexesClientWithBaseURI(baseURI string) IndexesClient {
	return IndexesClient{NewWithBaseURI(baseURI)}
}

// Analyze shows how an analyzer breaks text into tokens.
//
// indexName is the name of the index for which to test an analyzer. request is the text and analyzer or analysis
// components to test. clientRequestID is the tracking ID sent with the request to help with debugging.
func (client IndexesClient) Analyze(indexName string, request AnalyzeRequest, clientRequestID *uuid.UUID) (result AnalyzeResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.Text", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "search.IndexesClient", "Analyze")
	}

	req, err := client.AnalyzePreparer(indexName, request, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Analyze", nil, "Failure preparing request")
		return
	}

	resp, err := client.AnalyzeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Analyze", resp, "Failure sending request")
		return
	}

	result, err = client.AnalyzeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Analyze", resp, "Failure responding to request")
	}

	return
}

// AnalyzePreparer prepares the Analyze request.
func (client IndexesClient) AnalyzePreparer(indexName string, request AnalyzeRequest, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"indexName": autorest.Encode("path", indexName),
	}

	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/indexes('{indexName}')/search.analyze", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// AnalyzeSender sends the Analyze request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) AnalyzeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// AnalyzeResponder handles the response to the Analyze request. The method always
// closes the http.Response Body.
func (client IndexesClient) AnalyzeResponder(resp *http.Response) (result AnalyzeResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates a new Azure Search index.
//
// indexParameter is the definition of the index to create. clientRequestID is the tracking ID sent with the request to
// help with debugging.
func (client IndexesClient) Create(indexParameter Index, clientRequestID *uuid.UUID) (result Index, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: indexParameter,
			Constraints: []validation.Constraint{{Target: "indexParameter.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexParameter.Fields", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexParameter.CorsOptions", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "indexParameter.CorsOptions.AllowedOrigins", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "search.IndexesClient", "Create")
	}

	req, err := client.CreatePreparer(indexParameter, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client IndexesClient) CreatePreparer(indexParameter Index, clientRequestID *uuid.UUID) (*http.Request, error) {
	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/indexes"),
		autorest.WithJSON(indexParameter),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client IndexesClient) CreateResponder(resp *http.Response) (result Index, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates a new Azure Search index or updates an index if it already exists.
//
// indexName is the definition of the index to create or update. indexParameter is the definition of the index to
// create or update. allowIndexDowntime is allows new analyzers, tokenizers, token filters, or char filters to be added
// to an index by taking the index offline for at least a few seconds. This temporarily causes indexing and query
// requests to fail. Performance and write availability of the index can be impaired for several minutes after the
// index is updated, or longer for very large indexes. clientRequestID is the tracking ID sent with the request to help
// with debugging. ifMatch is defines the If-Match condition. The operation will be performed only if the ETag on the
// server matches this value. ifNoneMatch is defines the If-None-Match condition. The operation will be performed only
// if the ETag on the server does not match this value.
func (client IndexesClient) CreateOrUpdate(indexName string, indexParameter Index, allowIndexDowntime *bool, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result Index, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: indexParameter,
			Constraints: []validation.Constraint{{Target: "indexParameter.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexParameter.Fields", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "indexParameter.CorsOptions", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "indexParameter.CorsOptions.AllowedOrigins", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "search.IndexesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(indexName, indexParameter, allowIndexDowntime, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IndexesClient) CreateOrUpdatePreparer(indexName string, indexParameter Index, allowIndexDowntime *bool, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"indexName": autorest.Encode("path", indexName),
	}

	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if allowIndexDowntime != nil {
		queryParameters["allowIndexDowntime"] = autorest.Encode("query", *allowIndexDowntime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/indexes('{indexName}')", pathParameters),
		autorest.WithJSON(indexParameter),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IndexesClient) CreateOrUpdateResponder(resp *http.Response) (result Index, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an Azure Search index and all the documents it contains.
//
// indexName is the name of the index to delete. clientRequestID is the tracking ID sent with the request to help with
// debugging. ifMatch is defines the If-Match condition. The operation will be performed only if the ETag on the server
// matches this value. ifNoneMatch is defines the If-None-Match condition. The operation will be performed only if the
// ETag on the server does not match this value.
func (client IndexesClient) Delete(indexName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(indexName, clientRequestID, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IndexesClient) DeletePreparer(indexName string, clientRequestID *uuid.UUID, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"indexName": autorest.Encode("path", indexName),
	}

	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/indexes('{indexName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IndexesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves an index definition from Azure Search.
//
// indexName is the name of the index to retrieve. clientRequestID is the tracking ID sent with the request to help
// with debugging.
func (client IndexesClient) Get(indexName string, clientRequestID *uuid.UUID) (result Index, err error) {
	req, err := client.GetPreparer(indexName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IndexesClient) GetPreparer(indexName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"indexName": autorest.Encode("path", indexName),
	}

	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/indexes('{indexName}')", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IndexesClient) GetResponder(resp *http.Response) (result Index, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatistics returns statistics for the given index, including a document count and storage usage.
//
// indexName is the name of the index for which to retrieve statistics. clientRequestID is the tracking ID sent with
// the request to help with debugging.
func (client IndexesClient) GetStatistics(indexName string, clientRequestID *uuid.UUID) (result IndexGetStatisticsResult, err error) {
	req, err := client.GetStatisticsPreparer(indexName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "GetStatistics", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatisticsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "GetStatistics", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatisticsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "GetStatistics", resp, "Failure responding to request")
	}

	return
}

// GetStatisticsPreparer prepares the GetStatistics request.
func (client IndexesClient) GetStatisticsPreparer(indexName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"indexName": autorest.Encode("path", indexName),
	}

	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/indexes('{indexName}')/search.stats", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// GetStatisticsSender sends the GetStatistics request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) GetStatisticsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetStatisticsResponder handles the response to the GetStatistics request. The method always
// closes the http.Response Body.
func (client IndexesClient) GetStatisticsResponder(resp *http.Response) (result IndexGetStatisticsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all indexes available for an Azure Search service.
//
// selectParameter is selects which properties of the index definitions to retrieve. Specified as a comma-separated
// list of JSON property names, or '*' for all properties. The default is all properties. clientRequestID is the
// tracking ID sent with the request to help with debugging.
func (client IndexesClient) List(selectParameter string, clientRequestID *uuid.UUID) (result IndexListResult, err error) {
	req, err := client.ListPreparer(selectParameter, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.IndexesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IndexesClient) ListPreparer(selectParameter string, clientRequestID *uuid.UUID) (*http.Request, error) {
	const APIVersion = "2015-02-28-Preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/indexes"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IndexesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IndexesClient) ListResponder(resp *http.Response) (result IndexListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
