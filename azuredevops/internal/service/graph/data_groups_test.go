//go:build (all || core || data_sources || data_groups) && (!exclude_data_sources || !exclude_data_groups)
// +build all core data_sources data_groups
// +build !exclude_data_sources !exclude_data_groups

package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/graph"
	"github.com/microsoft/terraform-provider-azuredevops/azdosdkmocks"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/converter"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/testhelper"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// verifies that the translation for project_id to project_descriptor has proper error handling
func TestGroupsDataSource_DoesNotSwallowProjectDescriptorLookupError_Generic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectID := uuid.New()
	resourceData := createGroupsDataSource(t, projectID.String())

	graphClient := azdosdkmocks.NewMockGraphClient(ctrl)
	clients := &client.AggregatedClient{GraphClient: graphClient, Ctx: context.Background()}

	expectedArgs := graph.GetDescriptorArgs{StorageKey: &projectID}
	graphClient.
		EXPECT().
		GetDescriptor(clients.Ctx, expectedArgs).
		Return(nil, errors.New("GetDescriptor() Failed"))

	err := dataSourceGroupsRead(resourceData, clients)
	require.Contains(t, err.Error(), "GetDescriptor() Failed")
}

// verifies that the translation for project_id to project_descriptor has proper error handling
func TestGroupsDataSource_DoesNotSwallowProjectDescriptorLookupError_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectID := uuid.New()
	resourceData := createGroupsDataSource(t, projectID.String())

	graphClient := azdosdkmocks.NewMockGraphClient(ctrl)
	clients := &client.AggregatedClient{GraphClient: graphClient, Ctx: context.Background()}

	expectedArgs := graph.GetDescriptorArgs{StorageKey: &projectID}
	graphClient.
		EXPECT().
		GetDescriptor(clients.Ctx, expectedArgs).
		Return(nil, azuredevops.WrappedError{
			StatusCode: converter.Int(404),
		})

	err := dataSourceGroupsRead(resourceData, clients)
	require.Contains(t, err.Error(), "was not found")
}

// verifies that the group lookup functionality has proper error handling
func TestGroupsDataSource_DoesNotSwallowListGroupError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectID := uuid.New()
	resourceData := createGroupsDataSource(t, projectID.String())

	graphClient := azdosdkmocks.NewMockGraphClient(ctrl)
	clients := &client.AggregatedClient{GraphClient: graphClient, Ctx: context.Background()}

	expectedProjectDescriptorLookupArgs := graph.GetDescriptorArgs{StorageKey: &projectID}
	projectDescriptor := converter.String("descriptor")
	projectDescriptorResponse := graph.GraphDescriptorResult{Value: projectDescriptor}
	graphClient.
		EXPECT().
		GetDescriptor(clients.Ctx, expectedProjectDescriptorLookupArgs).
		Return(&projectDescriptorResponse, nil)

	expectedListGroupArgs := graph.ListGroupsArgs{ScopeDescriptor: projectDescriptor}
	graphClient.
		EXPECT().
		ListGroups(clients.Ctx, expectedListGroupArgs).
		Return(nil, errors.New("ListGroups() Failed"))

	err := dataSourceGroupsRead(resourceData, clients)
	require.Contains(t, err.Error(), "ListGroups() Failed")
}

// verifies that the group lookup functionality will make multiple API calls using the continuation token
// returned from the `ListGroups` api, until the API no longer returns a token
func TestGroupsDataSource_HandlesContinuationToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectID := uuid.New()
	originID := uuid.New()
	resourceData := createGroupsDataSource(t, projectID.String())

	graphClient := azdosdkmocks.NewMockGraphClient(ctrl)
	clients := &client.AggregatedClient{GraphClient: graphClient, Ctx: context.Background()}

	expectedProjectDescriptorLookupArgs := graph.GetDescriptorArgs{StorageKey: &projectID}
	projectDescriptor := converter.String("descriptor")
	projectDescriptorResponse := graph.GraphDescriptorResult{Value: projectDescriptor}
	graphClient.
		EXPECT().
		GetDescriptor(clients.Ctx, expectedProjectDescriptorLookupArgs).
		Return(&projectDescriptorResponse, nil)

	firstListGroupCallArgs := graph.ListGroupsArgs{ScopeDescriptor: projectDescriptor}
	continuationToken := "continuation-token"
	firstListGroupCallResponse := createPaginatedResponse(continuationToken, groupMeta{name: "name1", descriptor: "descriptor1", origin: "vsts", originId: originID.String()})

	var calls []*gomock.Call

	firstCall := graphClient.
		EXPECT().
		ListGroups(clients.Ctx, firstListGroupCallArgs).
		Return(firstListGroupCallResponse, nil)

	secondListGroupCallArgs := graph.ListGroupsArgs{ScopeDescriptor: projectDescriptor, ContinuationToken: &continuationToken}
	secondListGroupCallResponse := createPaginatedResponse("", groupMeta{name: "name2", descriptor: "descriptor2", origin: "vsts", originId: uuid.New().String()})
	secondCall := graphClient.
		EXPECT().
		ListGroups(clients.Ctx, secondListGroupCallArgs).
		Return(secondListGroupCallResponse, nil)

	calls = append(calls, firstCall, secondCall)
	calls = append(calls, graphClient.
		EXPECT().
		GetStorageKey(clients.Ctx, gomock.Any()).
		Return(&graph.GraphStorageKeyResult{
			Links: "",
			Value: &id,
		}, nil).Times(2))

	gomock.InOrder(testhelper.UnpackArray(calls)...)

	err := dataSourceGroupsRead(resourceData, clients)
	require.Nil(t, err)
	groups, ok := resourceData.GetOk("groups")
	require.True(t, ok)
	require.NotNil(t, groups)
	groupsSet, ok := groups.(*schema.Set)
	require.True(t, ok)
	require.NotNil(t, groupsSet)
	require.Equal(t, 2, groupsSet.Len())
}

func createGroupsDataSource(t *testing.T, projectID string) *schema.ResourceData {
	resourceData := schema.TestResourceDataRaw(t, DataGroups().Schema, nil)
	if projectID != "" {
		resourceData.Set("project_id", projectID)
	}
	return resourceData
}
