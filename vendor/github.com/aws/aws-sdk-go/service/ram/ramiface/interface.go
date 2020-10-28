// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ramiface provides an interface to enable mocking the AWS Resource Access Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ramiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ram"
)

// RAMAPI provides an interface to enable mocking the
// ram.RAM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Resource Access Manager.
//    func myFunc(svc ramiface.RAMAPI) bool {
//        // Make svc.AcceptResourceShareInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ram.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRAMClient struct {
//        ramiface.RAMAPI
//    }
//    func (m *mockRAMClient) AcceptResourceShareInvitation(input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRAMClient{}
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
type RAMAPI interface {
	AcceptResourceShareInvitation(*ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationWithContext(aws.Context, *ram.AcceptResourceShareInvitationInput, ...request.Option) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationRequest(*ram.AcceptResourceShareInvitationInput) (*request.Request, *ram.AcceptResourceShareInvitationOutput)

	AssociateResourceShare(*ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareWithContext(aws.Context, *ram.AssociateResourceShareInput, ...request.Option) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareRequest(*ram.AssociateResourceShareInput) (*request.Request, *ram.AssociateResourceShareOutput)

	CreateResourceShare(*ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareWithContext(aws.Context, *ram.CreateResourceShareInput, ...request.Option) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareRequest(*ram.CreateResourceShareInput) (*request.Request, *ram.CreateResourceShareOutput)

	DeleteResourceShare(*ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareWithContext(aws.Context, *ram.DeleteResourceShareInput, ...request.Option) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareRequest(*ram.DeleteResourceShareInput) (*request.Request, *ram.DeleteResourceShareOutput)

	DisassociateResourceShare(*ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareWithContext(aws.Context, *ram.DisassociateResourceShareInput, ...request.Option) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareRequest(*ram.DisassociateResourceShareInput) (*request.Request, *ram.DisassociateResourceShareOutput)

	EnableSharingWithAwsOrganization(*ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationWithContext(aws.Context, *ram.EnableSharingWithAwsOrganizationInput, ...request.Option) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationRequest(*ram.EnableSharingWithAwsOrganizationInput) (*request.Request, *ram.EnableSharingWithAwsOrganizationOutput)

	GetResourcePolicies(*ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesWithContext(aws.Context, *ram.GetResourcePoliciesInput, ...request.Option) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesRequest(*ram.GetResourcePoliciesInput) (*request.Request, *ram.GetResourcePoliciesOutput)

	GetResourcePoliciesPages(*ram.GetResourcePoliciesInput, func(*ram.GetResourcePoliciesOutput, bool) bool) error
	GetResourcePoliciesPagesWithContext(aws.Context, *ram.GetResourcePoliciesInput, func(*ram.GetResourcePoliciesOutput, bool) bool, ...request.Option) error

	GetResourceShareAssociations(*ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsWithContext(aws.Context, *ram.GetResourceShareAssociationsInput, ...request.Option) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsRequest(*ram.GetResourceShareAssociationsInput) (*request.Request, *ram.GetResourceShareAssociationsOutput)

	GetResourceShareAssociationsPages(*ram.GetResourceShareAssociationsInput, func(*ram.GetResourceShareAssociationsOutput, bool) bool) error
	GetResourceShareAssociationsPagesWithContext(aws.Context, *ram.GetResourceShareAssociationsInput, func(*ram.GetResourceShareAssociationsOutput, bool) bool, ...request.Option) error

	GetResourceShareInvitations(*ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsWithContext(aws.Context, *ram.GetResourceShareInvitationsInput, ...request.Option) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsRequest(*ram.GetResourceShareInvitationsInput) (*request.Request, *ram.GetResourceShareInvitationsOutput)

	GetResourceShareInvitationsPages(*ram.GetResourceShareInvitationsInput, func(*ram.GetResourceShareInvitationsOutput, bool) bool) error
	GetResourceShareInvitationsPagesWithContext(aws.Context, *ram.GetResourceShareInvitationsInput, func(*ram.GetResourceShareInvitationsOutput, bool) bool, ...request.Option) error

	GetResourceShares(*ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesWithContext(aws.Context, *ram.GetResourceSharesInput, ...request.Option) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesRequest(*ram.GetResourceSharesInput) (*request.Request, *ram.GetResourceSharesOutput)

	GetResourceSharesPages(*ram.GetResourceSharesInput, func(*ram.GetResourceSharesOutput, bool) bool) error
	GetResourceSharesPagesWithContext(aws.Context, *ram.GetResourceSharesInput, func(*ram.GetResourceSharesOutput, bool) bool, ...request.Option) error

	ListPendingInvitationResources(*ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesWithContext(aws.Context, *ram.ListPendingInvitationResourcesInput, ...request.Option) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesRequest(*ram.ListPendingInvitationResourcesInput) (*request.Request, *ram.ListPendingInvitationResourcesOutput)

	ListPendingInvitationResourcesPages(*ram.ListPendingInvitationResourcesInput, func(*ram.ListPendingInvitationResourcesOutput, bool) bool) error
	ListPendingInvitationResourcesPagesWithContext(aws.Context, *ram.ListPendingInvitationResourcesInput, func(*ram.ListPendingInvitationResourcesOutput, bool) bool, ...request.Option) error

	ListPrincipals(*ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsWithContext(aws.Context, *ram.ListPrincipalsInput, ...request.Option) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsRequest(*ram.ListPrincipalsInput) (*request.Request, *ram.ListPrincipalsOutput)

	ListPrincipalsPages(*ram.ListPrincipalsInput, func(*ram.ListPrincipalsOutput, bool) bool) error
	ListPrincipalsPagesWithContext(aws.Context, *ram.ListPrincipalsInput, func(*ram.ListPrincipalsOutput, bool) bool, ...request.Option) error

	ListResources(*ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *ram.ListResourcesInput, ...request.Option) (*ram.ListResourcesOutput, error)
	ListResourcesRequest(*ram.ListResourcesInput) (*request.Request, *ram.ListResourcesOutput)

	ListResourcesPages(*ram.ListResourcesInput, func(*ram.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *ram.ListResourcesInput, func(*ram.ListResourcesOutput, bool) bool, ...request.Option) error

	RejectResourceShareInvitation(*ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationWithContext(aws.Context, *ram.RejectResourceShareInvitationInput, ...request.Option) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationRequest(*ram.RejectResourceShareInvitationInput) (*request.Request, *ram.RejectResourceShareInvitationOutput)

	TagResource(*ram.TagResourceInput) (*ram.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ram.TagResourceInput, ...request.Option) (*ram.TagResourceOutput, error)
	TagResourceRequest(*ram.TagResourceInput) (*request.Request, *ram.TagResourceOutput)

	UntagResource(*ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ram.UntagResourceInput, ...request.Option) (*ram.UntagResourceOutput, error)
	UntagResourceRequest(*ram.UntagResourceInput) (*request.Request, *ram.UntagResourceOutput)

	UpdateResourceShare(*ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareWithContext(aws.Context, *ram.UpdateResourceShareInput, ...request.Option) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareRequest(*ram.UpdateResourceShareInput) (*request.Request, *ram.UpdateResourceShareOutput)
}

var _ RAMAPI = (*ram.RAM)(nil)
