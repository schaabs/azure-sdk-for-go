package containerregistry

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BuildsClient is the client for the Builds methods of the Containerregistry service.
type BuildsClient struct {
	BaseClient
}

// NewBuildsClient creates an instance of the BuildsClient client.
func NewBuildsClient(subscriptionID string, subscriptionID1 string) BuildsClient {
	return NewBuildsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewBuildsClientWithBaseURI creates an instance of the BuildsClient client.
func NewBuildsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) BuildsClient {
	return BuildsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Cancel sends the cancel request.
//
// buildID is the build identifier. resourceGroupName is the name of the resource group to which the container registry
// belongs. registryName is the name of the container registry.
func (client BuildsClient) Cancel(ctx context.Context, buildID string, resourceGroupName string, registryName string) (result BuildsCancelFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "Cancel")
	}

	req, err := client.CancelPreparer(ctx, buildID, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client BuildsClient) CancelPreparer(ctx context.Context, buildID string, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildId":           autorest.Encode("path", buildID),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) CancelSender(req *http.Request) (future BuildsCancelFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound))
	return
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client BuildsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
//
// buildID is the build id. resourceGroupName is the name of the resource group to which the container registry
// belongs. registryName is the name of the container registry.
func (client BuildsClient) Get(ctx context.Context, buildID string, resourceGroupName string, registryName string) (result Build, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "Get")
	}

	req, err := client.GetPreparer(ctx, buildID, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BuildsClient) GetPreparer(ctx context.Context, buildID string, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildId":           autorest.Encode("path", buildID),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BuildsClient) GetResponder(resp *http.Response) (result Build, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLogLink sends the get log link request.
//
// buildID is the build id. logParameters is the build log parameters. resourceGroupName is the name of the resource
// group to which the container registry belongs. registryName is the name of the container registry.
func (client BuildsClient) GetLogLink(ctx context.Context, buildID string, logParameters BuildLogParameters, resourceGroupName string, registryName string) (result BuildLogResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: logParameters,
			Constraints: []validation.Constraint{{Target: "logParameters.LogType", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "GetLogLink")
	}

	req, err := client.GetLogLinkPreparer(ctx, buildID, logParameters, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLogLinkSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", resp, "Failure sending request")
		return
	}

	result, err = client.GetLogLinkResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "GetLogLink", resp, "Failure responding to request")
	}

	return
}

// GetLogLinkPreparer prepares the GetLogLink request.
func (client BuildsClient) GetLogLinkPreparer(ctx context.Context, buildID string, logParameters BuildLogParameters, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildId":           autorest.Encode("path", buildID),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}/getLogLink", pathParameters),
		autorest.WithJSON(logParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLogLinkSender sends the GetLogLink request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) GetLogLinkSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetLogLinkResponder handles the response to the GetLogLink request. The method always
// closes the http.Response Body.
func (client BuildsClient) GetLogLinkResponder(resp *http.Response) (result BuildLogResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
//
// resourceGroupName is the name of the resource group to which the container registry belongs. registryName is the
// name of the container registry. filter is oData Filter options
func (client BuildsClient) List(ctx context.Context, resourceGroupName string, registryName string, filter string) (result BuildListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "List")
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.blr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", resp, "Failure sending request")
		return
	}

	result.blr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BuildsClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BuildsClient) ListResponder(resp *http.Response) (result BuildListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BuildsClient) listNextResults(lastResults BuildListResult) (result BuildListResult, err error) {
	req, err := lastResults.buildListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BuildsClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string, filter string) (result BuildListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, registryName, filter)
	return
}

// Queue sends the queue request.
//
// buildRequest is the parameters of a build that needs to queued. resourceGroupName is the name of the resource group
// to which the container registry belongs. registryName is the name of the container registry.
func (client BuildsClient) Queue(ctx context.Context, buildRequest QueueBuildRequest, resourceGroupName string, registryName string) (result BuildsQueueFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: buildRequest,
			Constraints: []validation.Constraint{{Target: "buildRequest.ImageName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "buildRequest.SourceLocation", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "buildRequest.BuildParameters", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "Queue")
	}

	req, err := client.QueuePreparer(ctx, buildRequest, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Queue", nil, "Failure preparing request")
		return
	}

	result, err = client.QueueSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Queue", result.Response(), "Failure sending request")
		return
	}

	return
}

// QueuePreparer prepares the Queue request.
func (client BuildsClient) QueuePreparer(ctx context.Context, buildRequest QueueBuildRequest, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/queue", pathParameters),
		autorest.WithJSON(buildRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueueSender sends the Queue request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) QueueSender(req *http.Request) (future BuildsQueueFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// QueueResponder handles the response to the Queue request. The method always
// closes the http.Response Body.
func (client BuildsClient) QueueResponder(resp *http.Response) (result Build, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request.
//
// buildID is the build id. buildUpdateParameters is the build update properties. resourceGroupName is the name of the
// resource group to which the container registry belongs. registryName is the name of the container registry.
func (client BuildsClient) Update(ctx context.Context, buildID string, buildUpdateParameters BuildUpdateParameters, resourceGroupName string, registryName string) (result BuildsUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.BuildsClient", "Update")
	}

	req, err := client.UpdatePreparer(ctx, buildID, buildUpdateParameters, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.BuildsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BuildsClient) UpdatePreparer(ctx context.Context, buildID string, buildUpdateParameters BuildUpdateParameters, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"buildId":           autorest.Encode("path", buildID),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/builds/{buildId}", pathParameters),
		autorest.WithJSON(buildUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BuildsClient) UpdateSender(req *http.Request) (future BuildsUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BuildsClient) UpdateResponder(resp *http.Response) (result Build, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
