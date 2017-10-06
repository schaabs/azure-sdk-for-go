package cache

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
	"net/http"
)

// RedisLinkedServerClient is the REST API for Azure Redis Cache Service.
type RedisLinkedServerClient struct {
	ManagementClient
}

// NewRedisLinkedServerClient creates an instance of the RedisLinkedServerClient client.
func NewRedisLinkedServerClient(subscriptionID string) RedisLinkedServerClient {
	return NewRedisLinkedServerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRedisLinkedServerClientWithBaseURI creates an instance of the RedisLinkedServerClient client.
func NewRedisLinkedServerClientWithBaseURI(baseURI string, subscriptionID string) RedisLinkedServerClient {
	return RedisLinkedServerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create adds a linked server to the Redis cache (requires Premium SKU). This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the Redis cache. linkedServerName is the
// name of the linked server that is being added to the Redis cache. parameters is parameters supplied to the Create
// Linked server operation.
func (client RedisLinkedServerClient) Create(resourceGroupName string, name string, linkedServerName string, parameters RedisLinkedServerCreateParameters, cancel <-chan struct{}) (<-chan RedisLinkedServerWithProperties, <-chan error) {
	resultChan := make(chan RedisLinkedServerWithProperties, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RedisLinkedServerCreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.RedisLinkedServerCreateProperties.LinkedRedisCacheID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RedisLinkedServerCreateProperties.LinkedRedisCacheLocation", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cache.RedisLinkedServerClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result RedisLinkedServerWithProperties
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, name, linkedServerName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client RedisLinkedServerClient) CreatePreparer(resourceGroupName string, name string, linkedServerName string, parameters RedisLinkedServerCreateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RedisLinkedServerClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RedisLinkedServerClient) CreateResponder(resp *http.Response) (result RedisLinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the linked server from a redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache. linkedServerName is the
// name of the linked server that is being added to the Redis cache.
func (client RedisLinkedServerClient) Delete(resourceGroupName string, name string, linkedServerName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RedisLinkedServerClient) DeletePreparer(resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RedisLinkedServerClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RedisLinkedServerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the detailed information about a linked server of a redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache. linkedServerName is the
// name of the linked server.
func (client RedisLinkedServerClient) Get(resourceGroupName string, name string, linkedServerName string) (result RedisLinkedServerWithProperties, err error) {
	req, err := client.GetPreparer(resourceGroupName, name, linkedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RedisLinkedServerClient) GetPreparer(resourceGroupName string, name string, linkedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServerName":  autorest.Encode("path", linkedServerName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RedisLinkedServerClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RedisLinkedServerClient) GetResponder(resp *http.Response) (result RedisLinkedServerWithProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of linked servers associated with this redis cache (requires Premium SKU).
//
// resourceGroupName is the name of the resource group. name is the name of the redis cache.
func (client RedisLinkedServerClient) List(resourceGroupName string, name string) (result RedisLinkedServerWithPropertiesList, err error) {
	req, err := client.ListPreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cache.RedisLinkedServerClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RedisLinkedServerClient) ListPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RedisLinkedServerClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RedisLinkedServerClient) ListResponder(resp *http.Response) (result RedisLinkedServerWithPropertiesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
