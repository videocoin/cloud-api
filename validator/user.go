package validator

import (
	"github.com/VideoCoin/cloud-api/rpc"
	"github.com/VideoCoin/cloud-api/users/v1"
)

func ValidateCreateUserRequest(r *v1.CreateUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateLoginUserRequest(r *v1.LoginUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateStartRecoveryUserRequest(r *v1.StartRecoveryUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateRecoverUserRequest(r *v1.RecoverUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}

func ValidateResetPasswordUserRequest(r *v1.ResetPasswordUserRequest) *rpc.MultiValidationError {
	return ValidateRequest(r)
}
