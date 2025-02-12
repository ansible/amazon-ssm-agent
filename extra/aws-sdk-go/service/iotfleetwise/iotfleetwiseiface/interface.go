// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotfleetwiseiface provides an interface to enable mocking the AWS IoT FleetWise service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotfleetwiseiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotfleetwise"
)

// IoTFleetWiseAPI provides an interface to enable mocking the
// iotfleetwise.IoTFleetWise service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT FleetWise.
//    func myFunc(svc iotfleetwiseiface.IoTFleetWiseAPI) bool {
//        // Make svc.AssociateVehicleFleet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iotfleetwise.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTFleetWiseClient struct {
//        iotfleetwiseiface.IoTFleetWiseAPI
//    }
//    func (m *mockIoTFleetWiseClient) AssociateVehicleFleet(input *iotfleetwise.AssociateVehicleFleetInput) (*iotfleetwise.AssociateVehicleFleetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTFleetWiseClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTFleetWiseAPI interface {
	AssociateVehicleFleet(*iotfleetwise.AssociateVehicleFleetInput) (*iotfleetwise.AssociateVehicleFleetOutput, error)
	AssociateVehicleFleetWithContext(aws.Context, *iotfleetwise.AssociateVehicleFleetInput, ...request.Option) (*iotfleetwise.AssociateVehicleFleetOutput, error)
	AssociateVehicleFleetRequest(*iotfleetwise.AssociateVehicleFleetInput) (*request.Request, *iotfleetwise.AssociateVehicleFleetOutput)

	BatchCreateVehicle(*iotfleetwise.BatchCreateVehicleInput) (*iotfleetwise.BatchCreateVehicleOutput, error)
	BatchCreateVehicleWithContext(aws.Context, *iotfleetwise.BatchCreateVehicleInput, ...request.Option) (*iotfleetwise.BatchCreateVehicleOutput, error)
	BatchCreateVehicleRequest(*iotfleetwise.BatchCreateVehicleInput) (*request.Request, *iotfleetwise.BatchCreateVehicleOutput)

	BatchUpdateVehicle(*iotfleetwise.BatchUpdateVehicleInput) (*iotfleetwise.BatchUpdateVehicleOutput, error)
	BatchUpdateVehicleWithContext(aws.Context, *iotfleetwise.BatchUpdateVehicleInput, ...request.Option) (*iotfleetwise.BatchUpdateVehicleOutput, error)
	BatchUpdateVehicleRequest(*iotfleetwise.BatchUpdateVehicleInput) (*request.Request, *iotfleetwise.BatchUpdateVehicleOutput)

	CreateCampaign(*iotfleetwise.CreateCampaignInput) (*iotfleetwise.CreateCampaignOutput, error)
	CreateCampaignWithContext(aws.Context, *iotfleetwise.CreateCampaignInput, ...request.Option) (*iotfleetwise.CreateCampaignOutput, error)
	CreateCampaignRequest(*iotfleetwise.CreateCampaignInput) (*request.Request, *iotfleetwise.CreateCampaignOutput)

	CreateDecoderManifest(*iotfleetwise.CreateDecoderManifestInput) (*iotfleetwise.CreateDecoderManifestOutput, error)
	CreateDecoderManifestWithContext(aws.Context, *iotfleetwise.CreateDecoderManifestInput, ...request.Option) (*iotfleetwise.CreateDecoderManifestOutput, error)
	CreateDecoderManifestRequest(*iotfleetwise.CreateDecoderManifestInput) (*request.Request, *iotfleetwise.CreateDecoderManifestOutput)

	CreateFleet(*iotfleetwise.CreateFleetInput) (*iotfleetwise.CreateFleetOutput, error)
	CreateFleetWithContext(aws.Context, *iotfleetwise.CreateFleetInput, ...request.Option) (*iotfleetwise.CreateFleetOutput, error)
	CreateFleetRequest(*iotfleetwise.CreateFleetInput) (*request.Request, *iotfleetwise.CreateFleetOutput)

	CreateModelManifest(*iotfleetwise.CreateModelManifestInput) (*iotfleetwise.CreateModelManifestOutput, error)
	CreateModelManifestWithContext(aws.Context, *iotfleetwise.CreateModelManifestInput, ...request.Option) (*iotfleetwise.CreateModelManifestOutput, error)
	CreateModelManifestRequest(*iotfleetwise.CreateModelManifestInput) (*request.Request, *iotfleetwise.CreateModelManifestOutput)

	CreateSignalCatalog(*iotfleetwise.CreateSignalCatalogInput) (*iotfleetwise.CreateSignalCatalogOutput, error)
	CreateSignalCatalogWithContext(aws.Context, *iotfleetwise.CreateSignalCatalogInput, ...request.Option) (*iotfleetwise.CreateSignalCatalogOutput, error)
	CreateSignalCatalogRequest(*iotfleetwise.CreateSignalCatalogInput) (*request.Request, *iotfleetwise.CreateSignalCatalogOutput)

	CreateVehicle(*iotfleetwise.CreateVehicleInput) (*iotfleetwise.CreateVehicleOutput, error)
	CreateVehicleWithContext(aws.Context, *iotfleetwise.CreateVehicleInput, ...request.Option) (*iotfleetwise.CreateVehicleOutput, error)
	CreateVehicleRequest(*iotfleetwise.CreateVehicleInput) (*request.Request, *iotfleetwise.CreateVehicleOutput)

	DeleteCampaign(*iotfleetwise.DeleteCampaignInput) (*iotfleetwise.DeleteCampaignOutput, error)
	DeleteCampaignWithContext(aws.Context, *iotfleetwise.DeleteCampaignInput, ...request.Option) (*iotfleetwise.DeleteCampaignOutput, error)
	DeleteCampaignRequest(*iotfleetwise.DeleteCampaignInput) (*request.Request, *iotfleetwise.DeleteCampaignOutput)

	DeleteDecoderManifest(*iotfleetwise.DeleteDecoderManifestInput) (*iotfleetwise.DeleteDecoderManifestOutput, error)
	DeleteDecoderManifestWithContext(aws.Context, *iotfleetwise.DeleteDecoderManifestInput, ...request.Option) (*iotfleetwise.DeleteDecoderManifestOutput, error)
	DeleteDecoderManifestRequest(*iotfleetwise.DeleteDecoderManifestInput) (*request.Request, *iotfleetwise.DeleteDecoderManifestOutput)

	DeleteFleet(*iotfleetwise.DeleteFleetInput) (*iotfleetwise.DeleteFleetOutput, error)
	DeleteFleetWithContext(aws.Context, *iotfleetwise.DeleteFleetInput, ...request.Option) (*iotfleetwise.DeleteFleetOutput, error)
	DeleteFleetRequest(*iotfleetwise.DeleteFleetInput) (*request.Request, *iotfleetwise.DeleteFleetOutput)

	DeleteModelManifest(*iotfleetwise.DeleteModelManifestInput) (*iotfleetwise.DeleteModelManifestOutput, error)
	DeleteModelManifestWithContext(aws.Context, *iotfleetwise.DeleteModelManifestInput, ...request.Option) (*iotfleetwise.DeleteModelManifestOutput, error)
	DeleteModelManifestRequest(*iotfleetwise.DeleteModelManifestInput) (*request.Request, *iotfleetwise.DeleteModelManifestOutput)

	DeleteSignalCatalog(*iotfleetwise.DeleteSignalCatalogInput) (*iotfleetwise.DeleteSignalCatalogOutput, error)
	DeleteSignalCatalogWithContext(aws.Context, *iotfleetwise.DeleteSignalCatalogInput, ...request.Option) (*iotfleetwise.DeleteSignalCatalogOutput, error)
	DeleteSignalCatalogRequest(*iotfleetwise.DeleteSignalCatalogInput) (*request.Request, *iotfleetwise.DeleteSignalCatalogOutput)

	DeleteVehicle(*iotfleetwise.DeleteVehicleInput) (*iotfleetwise.DeleteVehicleOutput, error)
	DeleteVehicleWithContext(aws.Context, *iotfleetwise.DeleteVehicleInput, ...request.Option) (*iotfleetwise.DeleteVehicleOutput, error)
	DeleteVehicleRequest(*iotfleetwise.DeleteVehicleInput) (*request.Request, *iotfleetwise.DeleteVehicleOutput)

	DisassociateVehicleFleet(*iotfleetwise.DisassociateVehicleFleetInput) (*iotfleetwise.DisassociateVehicleFleetOutput, error)
	DisassociateVehicleFleetWithContext(aws.Context, *iotfleetwise.DisassociateVehicleFleetInput, ...request.Option) (*iotfleetwise.DisassociateVehicleFleetOutput, error)
	DisassociateVehicleFleetRequest(*iotfleetwise.DisassociateVehicleFleetInput) (*request.Request, *iotfleetwise.DisassociateVehicleFleetOutput)

	GetCampaign(*iotfleetwise.GetCampaignInput) (*iotfleetwise.GetCampaignOutput, error)
	GetCampaignWithContext(aws.Context, *iotfleetwise.GetCampaignInput, ...request.Option) (*iotfleetwise.GetCampaignOutput, error)
	GetCampaignRequest(*iotfleetwise.GetCampaignInput) (*request.Request, *iotfleetwise.GetCampaignOutput)

	GetDecoderManifest(*iotfleetwise.GetDecoderManifestInput) (*iotfleetwise.GetDecoderManifestOutput, error)
	GetDecoderManifestWithContext(aws.Context, *iotfleetwise.GetDecoderManifestInput, ...request.Option) (*iotfleetwise.GetDecoderManifestOutput, error)
	GetDecoderManifestRequest(*iotfleetwise.GetDecoderManifestInput) (*request.Request, *iotfleetwise.GetDecoderManifestOutput)

	GetFleet(*iotfleetwise.GetFleetInput) (*iotfleetwise.GetFleetOutput, error)
	GetFleetWithContext(aws.Context, *iotfleetwise.GetFleetInput, ...request.Option) (*iotfleetwise.GetFleetOutput, error)
	GetFleetRequest(*iotfleetwise.GetFleetInput) (*request.Request, *iotfleetwise.GetFleetOutput)

	GetLoggingOptions(*iotfleetwise.GetLoggingOptionsInput) (*iotfleetwise.GetLoggingOptionsOutput, error)
	GetLoggingOptionsWithContext(aws.Context, *iotfleetwise.GetLoggingOptionsInput, ...request.Option) (*iotfleetwise.GetLoggingOptionsOutput, error)
	GetLoggingOptionsRequest(*iotfleetwise.GetLoggingOptionsInput) (*request.Request, *iotfleetwise.GetLoggingOptionsOutput)

	GetModelManifest(*iotfleetwise.GetModelManifestInput) (*iotfleetwise.GetModelManifestOutput, error)
	GetModelManifestWithContext(aws.Context, *iotfleetwise.GetModelManifestInput, ...request.Option) (*iotfleetwise.GetModelManifestOutput, error)
	GetModelManifestRequest(*iotfleetwise.GetModelManifestInput) (*request.Request, *iotfleetwise.GetModelManifestOutput)

	GetRegisterAccountStatus(*iotfleetwise.GetRegisterAccountStatusInput) (*iotfleetwise.GetRegisterAccountStatusOutput, error)
	GetRegisterAccountStatusWithContext(aws.Context, *iotfleetwise.GetRegisterAccountStatusInput, ...request.Option) (*iotfleetwise.GetRegisterAccountStatusOutput, error)
	GetRegisterAccountStatusRequest(*iotfleetwise.GetRegisterAccountStatusInput) (*request.Request, *iotfleetwise.GetRegisterAccountStatusOutput)

	GetSignalCatalog(*iotfleetwise.GetSignalCatalogInput) (*iotfleetwise.GetSignalCatalogOutput, error)
	GetSignalCatalogWithContext(aws.Context, *iotfleetwise.GetSignalCatalogInput, ...request.Option) (*iotfleetwise.GetSignalCatalogOutput, error)
	GetSignalCatalogRequest(*iotfleetwise.GetSignalCatalogInput) (*request.Request, *iotfleetwise.GetSignalCatalogOutput)

	GetVehicle(*iotfleetwise.GetVehicleInput) (*iotfleetwise.GetVehicleOutput, error)
	GetVehicleWithContext(aws.Context, *iotfleetwise.GetVehicleInput, ...request.Option) (*iotfleetwise.GetVehicleOutput, error)
	GetVehicleRequest(*iotfleetwise.GetVehicleInput) (*request.Request, *iotfleetwise.GetVehicleOutput)

	GetVehicleStatus(*iotfleetwise.GetVehicleStatusInput) (*iotfleetwise.GetVehicleStatusOutput, error)
	GetVehicleStatusWithContext(aws.Context, *iotfleetwise.GetVehicleStatusInput, ...request.Option) (*iotfleetwise.GetVehicleStatusOutput, error)
	GetVehicleStatusRequest(*iotfleetwise.GetVehicleStatusInput) (*request.Request, *iotfleetwise.GetVehicleStatusOutput)

	GetVehicleStatusPages(*iotfleetwise.GetVehicleStatusInput, func(*iotfleetwise.GetVehicleStatusOutput, bool) bool) error
	GetVehicleStatusPagesWithContext(aws.Context, *iotfleetwise.GetVehicleStatusInput, func(*iotfleetwise.GetVehicleStatusOutput, bool) bool, ...request.Option) error

	ImportDecoderManifest(*iotfleetwise.ImportDecoderManifestInput) (*iotfleetwise.ImportDecoderManifestOutput, error)
	ImportDecoderManifestWithContext(aws.Context, *iotfleetwise.ImportDecoderManifestInput, ...request.Option) (*iotfleetwise.ImportDecoderManifestOutput, error)
	ImportDecoderManifestRequest(*iotfleetwise.ImportDecoderManifestInput) (*request.Request, *iotfleetwise.ImportDecoderManifestOutput)

	ImportSignalCatalog(*iotfleetwise.ImportSignalCatalogInput) (*iotfleetwise.ImportSignalCatalogOutput, error)
	ImportSignalCatalogWithContext(aws.Context, *iotfleetwise.ImportSignalCatalogInput, ...request.Option) (*iotfleetwise.ImportSignalCatalogOutput, error)
	ImportSignalCatalogRequest(*iotfleetwise.ImportSignalCatalogInput) (*request.Request, *iotfleetwise.ImportSignalCatalogOutput)

	ListCampaigns(*iotfleetwise.ListCampaignsInput) (*iotfleetwise.ListCampaignsOutput, error)
	ListCampaignsWithContext(aws.Context, *iotfleetwise.ListCampaignsInput, ...request.Option) (*iotfleetwise.ListCampaignsOutput, error)
	ListCampaignsRequest(*iotfleetwise.ListCampaignsInput) (*request.Request, *iotfleetwise.ListCampaignsOutput)

	ListCampaignsPages(*iotfleetwise.ListCampaignsInput, func(*iotfleetwise.ListCampaignsOutput, bool) bool) error
	ListCampaignsPagesWithContext(aws.Context, *iotfleetwise.ListCampaignsInput, func(*iotfleetwise.ListCampaignsOutput, bool) bool, ...request.Option) error

	ListDecoderManifestNetworkInterfaces(*iotfleetwise.ListDecoderManifestNetworkInterfacesInput) (*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, error)
	ListDecoderManifestNetworkInterfacesWithContext(aws.Context, *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, ...request.Option) (*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, error)
	ListDecoderManifestNetworkInterfacesRequest(*iotfleetwise.ListDecoderManifestNetworkInterfacesInput) (*request.Request, *iotfleetwise.ListDecoderManifestNetworkInterfacesOutput)

	ListDecoderManifestNetworkInterfacesPages(*iotfleetwise.ListDecoderManifestNetworkInterfacesInput, func(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, bool) bool) error
	ListDecoderManifestNetworkInterfacesPagesWithContext(aws.Context, *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, func(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, bool) bool, ...request.Option) error

	ListDecoderManifestSignals(*iotfleetwise.ListDecoderManifestSignalsInput) (*iotfleetwise.ListDecoderManifestSignalsOutput, error)
	ListDecoderManifestSignalsWithContext(aws.Context, *iotfleetwise.ListDecoderManifestSignalsInput, ...request.Option) (*iotfleetwise.ListDecoderManifestSignalsOutput, error)
	ListDecoderManifestSignalsRequest(*iotfleetwise.ListDecoderManifestSignalsInput) (*request.Request, *iotfleetwise.ListDecoderManifestSignalsOutput)

	ListDecoderManifestSignalsPages(*iotfleetwise.ListDecoderManifestSignalsInput, func(*iotfleetwise.ListDecoderManifestSignalsOutput, bool) bool) error
	ListDecoderManifestSignalsPagesWithContext(aws.Context, *iotfleetwise.ListDecoderManifestSignalsInput, func(*iotfleetwise.ListDecoderManifestSignalsOutput, bool) bool, ...request.Option) error

	ListDecoderManifests(*iotfleetwise.ListDecoderManifestsInput) (*iotfleetwise.ListDecoderManifestsOutput, error)
	ListDecoderManifestsWithContext(aws.Context, *iotfleetwise.ListDecoderManifestsInput, ...request.Option) (*iotfleetwise.ListDecoderManifestsOutput, error)
	ListDecoderManifestsRequest(*iotfleetwise.ListDecoderManifestsInput) (*request.Request, *iotfleetwise.ListDecoderManifestsOutput)

	ListDecoderManifestsPages(*iotfleetwise.ListDecoderManifestsInput, func(*iotfleetwise.ListDecoderManifestsOutput, bool) bool) error
	ListDecoderManifestsPagesWithContext(aws.Context, *iotfleetwise.ListDecoderManifestsInput, func(*iotfleetwise.ListDecoderManifestsOutput, bool) bool, ...request.Option) error

	ListFleets(*iotfleetwise.ListFleetsInput) (*iotfleetwise.ListFleetsOutput, error)
	ListFleetsWithContext(aws.Context, *iotfleetwise.ListFleetsInput, ...request.Option) (*iotfleetwise.ListFleetsOutput, error)
	ListFleetsRequest(*iotfleetwise.ListFleetsInput) (*request.Request, *iotfleetwise.ListFleetsOutput)

	ListFleetsPages(*iotfleetwise.ListFleetsInput, func(*iotfleetwise.ListFleetsOutput, bool) bool) error
	ListFleetsPagesWithContext(aws.Context, *iotfleetwise.ListFleetsInput, func(*iotfleetwise.ListFleetsOutput, bool) bool, ...request.Option) error

	ListFleetsForVehicle(*iotfleetwise.ListFleetsForVehicleInput) (*iotfleetwise.ListFleetsForVehicleOutput, error)
	ListFleetsForVehicleWithContext(aws.Context, *iotfleetwise.ListFleetsForVehicleInput, ...request.Option) (*iotfleetwise.ListFleetsForVehicleOutput, error)
	ListFleetsForVehicleRequest(*iotfleetwise.ListFleetsForVehicleInput) (*request.Request, *iotfleetwise.ListFleetsForVehicleOutput)

	ListFleetsForVehiclePages(*iotfleetwise.ListFleetsForVehicleInput, func(*iotfleetwise.ListFleetsForVehicleOutput, bool) bool) error
	ListFleetsForVehiclePagesWithContext(aws.Context, *iotfleetwise.ListFleetsForVehicleInput, func(*iotfleetwise.ListFleetsForVehicleOutput, bool) bool, ...request.Option) error

	ListModelManifestNodes(*iotfleetwise.ListModelManifestNodesInput) (*iotfleetwise.ListModelManifestNodesOutput, error)
	ListModelManifestNodesWithContext(aws.Context, *iotfleetwise.ListModelManifestNodesInput, ...request.Option) (*iotfleetwise.ListModelManifestNodesOutput, error)
	ListModelManifestNodesRequest(*iotfleetwise.ListModelManifestNodesInput) (*request.Request, *iotfleetwise.ListModelManifestNodesOutput)

	ListModelManifestNodesPages(*iotfleetwise.ListModelManifestNodesInput, func(*iotfleetwise.ListModelManifestNodesOutput, bool) bool) error
	ListModelManifestNodesPagesWithContext(aws.Context, *iotfleetwise.ListModelManifestNodesInput, func(*iotfleetwise.ListModelManifestNodesOutput, bool) bool, ...request.Option) error

	ListModelManifests(*iotfleetwise.ListModelManifestsInput) (*iotfleetwise.ListModelManifestsOutput, error)
	ListModelManifestsWithContext(aws.Context, *iotfleetwise.ListModelManifestsInput, ...request.Option) (*iotfleetwise.ListModelManifestsOutput, error)
	ListModelManifestsRequest(*iotfleetwise.ListModelManifestsInput) (*request.Request, *iotfleetwise.ListModelManifestsOutput)

	ListModelManifestsPages(*iotfleetwise.ListModelManifestsInput, func(*iotfleetwise.ListModelManifestsOutput, bool) bool) error
	ListModelManifestsPagesWithContext(aws.Context, *iotfleetwise.ListModelManifestsInput, func(*iotfleetwise.ListModelManifestsOutput, bool) bool, ...request.Option) error

	ListSignalCatalogNodes(*iotfleetwise.ListSignalCatalogNodesInput) (*iotfleetwise.ListSignalCatalogNodesOutput, error)
	ListSignalCatalogNodesWithContext(aws.Context, *iotfleetwise.ListSignalCatalogNodesInput, ...request.Option) (*iotfleetwise.ListSignalCatalogNodesOutput, error)
	ListSignalCatalogNodesRequest(*iotfleetwise.ListSignalCatalogNodesInput) (*request.Request, *iotfleetwise.ListSignalCatalogNodesOutput)

	ListSignalCatalogNodesPages(*iotfleetwise.ListSignalCatalogNodesInput, func(*iotfleetwise.ListSignalCatalogNodesOutput, bool) bool) error
	ListSignalCatalogNodesPagesWithContext(aws.Context, *iotfleetwise.ListSignalCatalogNodesInput, func(*iotfleetwise.ListSignalCatalogNodesOutput, bool) bool, ...request.Option) error

	ListSignalCatalogs(*iotfleetwise.ListSignalCatalogsInput) (*iotfleetwise.ListSignalCatalogsOutput, error)
	ListSignalCatalogsWithContext(aws.Context, *iotfleetwise.ListSignalCatalogsInput, ...request.Option) (*iotfleetwise.ListSignalCatalogsOutput, error)
	ListSignalCatalogsRequest(*iotfleetwise.ListSignalCatalogsInput) (*request.Request, *iotfleetwise.ListSignalCatalogsOutput)

	ListSignalCatalogsPages(*iotfleetwise.ListSignalCatalogsInput, func(*iotfleetwise.ListSignalCatalogsOutput, bool) bool) error
	ListSignalCatalogsPagesWithContext(aws.Context, *iotfleetwise.ListSignalCatalogsInput, func(*iotfleetwise.ListSignalCatalogsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*iotfleetwise.ListTagsForResourceInput) (*iotfleetwise.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotfleetwise.ListTagsForResourceInput, ...request.Option) (*iotfleetwise.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotfleetwise.ListTagsForResourceInput) (*request.Request, *iotfleetwise.ListTagsForResourceOutput)

	ListVehicles(*iotfleetwise.ListVehiclesInput) (*iotfleetwise.ListVehiclesOutput, error)
	ListVehiclesWithContext(aws.Context, *iotfleetwise.ListVehiclesInput, ...request.Option) (*iotfleetwise.ListVehiclesOutput, error)
	ListVehiclesRequest(*iotfleetwise.ListVehiclesInput) (*request.Request, *iotfleetwise.ListVehiclesOutput)

	ListVehiclesPages(*iotfleetwise.ListVehiclesInput, func(*iotfleetwise.ListVehiclesOutput, bool) bool) error
	ListVehiclesPagesWithContext(aws.Context, *iotfleetwise.ListVehiclesInput, func(*iotfleetwise.ListVehiclesOutput, bool) bool, ...request.Option) error

	ListVehiclesInFleet(*iotfleetwise.ListVehiclesInFleetInput) (*iotfleetwise.ListVehiclesInFleetOutput, error)
	ListVehiclesInFleetWithContext(aws.Context, *iotfleetwise.ListVehiclesInFleetInput, ...request.Option) (*iotfleetwise.ListVehiclesInFleetOutput, error)
	ListVehiclesInFleetRequest(*iotfleetwise.ListVehiclesInFleetInput) (*request.Request, *iotfleetwise.ListVehiclesInFleetOutput)

	ListVehiclesInFleetPages(*iotfleetwise.ListVehiclesInFleetInput, func(*iotfleetwise.ListVehiclesInFleetOutput, bool) bool) error
	ListVehiclesInFleetPagesWithContext(aws.Context, *iotfleetwise.ListVehiclesInFleetInput, func(*iotfleetwise.ListVehiclesInFleetOutput, bool) bool, ...request.Option) error

	PutLoggingOptions(*iotfleetwise.PutLoggingOptionsInput) (*iotfleetwise.PutLoggingOptionsOutput, error)
	PutLoggingOptionsWithContext(aws.Context, *iotfleetwise.PutLoggingOptionsInput, ...request.Option) (*iotfleetwise.PutLoggingOptionsOutput, error)
	PutLoggingOptionsRequest(*iotfleetwise.PutLoggingOptionsInput) (*request.Request, *iotfleetwise.PutLoggingOptionsOutput)

	RegisterAccount(*iotfleetwise.RegisterAccountInput) (*iotfleetwise.RegisterAccountOutput, error)
	RegisterAccountWithContext(aws.Context, *iotfleetwise.RegisterAccountInput, ...request.Option) (*iotfleetwise.RegisterAccountOutput, error)
	RegisterAccountRequest(*iotfleetwise.RegisterAccountInput) (*request.Request, *iotfleetwise.RegisterAccountOutput)

	TagResource(*iotfleetwise.TagResourceInput) (*iotfleetwise.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotfleetwise.TagResourceInput, ...request.Option) (*iotfleetwise.TagResourceOutput, error)
	TagResourceRequest(*iotfleetwise.TagResourceInput) (*request.Request, *iotfleetwise.TagResourceOutput)

	UntagResource(*iotfleetwise.UntagResourceInput) (*iotfleetwise.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotfleetwise.UntagResourceInput, ...request.Option) (*iotfleetwise.UntagResourceOutput, error)
	UntagResourceRequest(*iotfleetwise.UntagResourceInput) (*request.Request, *iotfleetwise.UntagResourceOutput)

	UpdateCampaign(*iotfleetwise.UpdateCampaignInput) (*iotfleetwise.UpdateCampaignOutput, error)
	UpdateCampaignWithContext(aws.Context, *iotfleetwise.UpdateCampaignInput, ...request.Option) (*iotfleetwise.UpdateCampaignOutput, error)
	UpdateCampaignRequest(*iotfleetwise.UpdateCampaignInput) (*request.Request, *iotfleetwise.UpdateCampaignOutput)

	UpdateDecoderManifest(*iotfleetwise.UpdateDecoderManifestInput) (*iotfleetwise.UpdateDecoderManifestOutput, error)
	UpdateDecoderManifestWithContext(aws.Context, *iotfleetwise.UpdateDecoderManifestInput, ...request.Option) (*iotfleetwise.UpdateDecoderManifestOutput, error)
	UpdateDecoderManifestRequest(*iotfleetwise.UpdateDecoderManifestInput) (*request.Request, *iotfleetwise.UpdateDecoderManifestOutput)

	UpdateFleet(*iotfleetwise.UpdateFleetInput) (*iotfleetwise.UpdateFleetOutput, error)
	UpdateFleetWithContext(aws.Context, *iotfleetwise.UpdateFleetInput, ...request.Option) (*iotfleetwise.UpdateFleetOutput, error)
	UpdateFleetRequest(*iotfleetwise.UpdateFleetInput) (*request.Request, *iotfleetwise.UpdateFleetOutput)

	UpdateModelManifest(*iotfleetwise.UpdateModelManifestInput) (*iotfleetwise.UpdateModelManifestOutput, error)
	UpdateModelManifestWithContext(aws.Context, *iotfleetwise.UpdateModelManifestInput, ...request.Option) (*iotfleetwise.UpdateModelManifestOutput, error)
	UpdateModelManifestRequest(*iotfleetwise.UpdateModelManifestInput) (*request.Request, *iotfleetwise.UpdateModelManifestOutput)

	UpdateSignalCatalog(*iotfleetwise.UpdateSignalCatalogInput) (*iotfleetwise.UpdateSignalCatalogOutput, error)
	UpdateSignalCatalogWithContext(aws.Context, *iotfleetwise.UpdateSignalCatalogInput, ...request.Option) (*iotfleetwise.UpdateSignalCatalogOutput, error)
	UpdateSignalCatalogRequest(*iotfleetwise.UpdateSignalCatalogInput) (*request.Request, *iotfleetwise.UpdateSignalCatalogOutput)

	UpdateVehicle(*iotfleetwise.UpdateVehicleInput) (*iotfleetwise.UpdateVehicleOutput, error)
	UpdateVehicleWithContext(aws.Context, *iotfleetwise.UpdateVehicleInput, ...request.Option) (*iotfleetwise.UpdateVehicleOutput, error)
	UpdateVehicleRequest(*iotfleetwise.UpdateVehicleInput) (*request.Request, *iotfleetwise.UpdateVehicleOutput)
}

var _ IoTFleetWiseAPI = (*iotfleetwise.IoTFleetWise)(nil)
