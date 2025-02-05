package insecure_services_v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mtls_bot_root/authentication"
	"regexp"
)

const (
	// error_id_required - The ID is required error
	errorIDRequired = "ID_REQUIRED"
	// error_id_invalid_format - The ID is in an invalid format error
	errorIDInvalidFormat = "ID_INVALID_FORMAT"
	// error_id_in_use - The ID is already in use error
	errorIDInUse = "ID_IN_USE"
)

type AuthenticationService struct{}

// Start - Start an authentication flow
func (service *AuthenticationService) Start(ctx context.Context, in *AuthenticationStartRequest) (*AuthenticationStartResponse, error) {

	// Ensure that the id is valid.
	// ----------------------------------------
	// 1. The ID cannot be empty.
	// 2. The ID must be a valid UUID.
	// 3. The ID is not already in use.
	//
	if in.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, errorIDRequired)
	}

	uuidPattern := `^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$`
	if !regexp.MustCompile(uuidPattern).MatchString(in.Id) {
		return nil, status.Errorf(codes.InvalidArgument, errorIDInvalidFormat)
	}

	if authentication.IsIdUnique(in.Id) {
		return nil, status.Errorf(codes.InvalidArgument, errorIDInUse)
	}

	// Validate the public key of the client.
	// ----------------------------------------
	//

	// todo: Validate the public key of the client.

	// Initialize the authentication flow.
	// ----------------------------------------
	//

	// todo: Initialize the authentication flow.

	return &AuthenticationStartResponse{
		Success: false,
	}, nil
}

// Cancel - Cancel an authentication flow
func (service *AuthenticationService) Cancel(ctx context.Context, in *AuthenticationFlowCancelRequest) (*AuthenticationFlowCancelResponse, error) {
	return &AuthenticationFlowCancelResponse{
		Success: false,
	}, nil
}

func (service *AuthenticationService) mustEmbedUnimplementedAuthenticationServer() {}
