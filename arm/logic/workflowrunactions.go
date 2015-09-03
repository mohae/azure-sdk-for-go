package logic

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// WorkflowRunActions Client
type WorkflowRunActionsClient struct {
	LogicManagementClient
}

func NewWorkflowRunActionsClient(subscriptionId string) WorkflowRunActionsClient {
	return NewWorkflowRunActionsClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewWorkflowRunActionsClientWithBaseUri(baseUri string, subscriptionId string) WorkflowRunActionsClient {
	return WorkflowRunActionsClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// List gets a list of workflow run actions.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name. top is the number of items to be
// included in the result. filter is the filter to apply on the operation.
func (client WorkflowRunActionsClient) List(resourceGroupName string, workflowName string, runName string, top int, filter string) (result WorkflowRunActionListResult, ae autorest.Error) {
	req, err := client.NewListRequest(resourceGroupName, workflowName, runName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client WorkflowRunActionsClient) NewListRequest(resourceGroupName string, workflowName string, runName string, top int, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"$top":        top,
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client WorkflowRunActionsClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions"))
}

// Get gets a workflow run action.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name. actionName is the workflow action
// name.
func (client WorkflowRunActionsClient) Get(resourceGroupName string, workflowName string, runName string, actionName string) (result WorkflowRunAction, ae autorest.Error) {
	req, err := client.NewGetRequest(resourceGroupName, workflowName, runName, actionName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "Get", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "Get", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(
		client,
		req,
		autorest.DoErrorUnlessStatusCode(http.StatusOK))

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "Get", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "logic.WorkflowRunActionsClient", "Get", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the Get request.
func (client WorkflowRunActionsClient) NewGetRequest(resourceGroupName string, workflowName string, runName string, actionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        url.QueryEscape(actionName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionId),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.GetRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Get request.
func (client WorkflowRunActionsClient) GetRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}"))
}