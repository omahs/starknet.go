package rpc

import (
	"context"
)

func (provider *Provider) AddInvokeTransaction(ctx context.Context, invokeTxn InvokeTxnV1) (*AddInvokeTransactionResponse, error) {
	var output AddInvokeTransactionResponse
	if err := do(ctx, provider.c, "starknet_addInvokeTransaction", &output, invokeTxn); err != nil {
		if unexpectedErr, ok := isErrUnexpectedError(err); ok {
			return nil, unexpectedErr
		}
		return nil, tryUnwrapToRPCErr(
			err,
			ErrInsufficientAccountBalance,
			ErrInsufficientMaxFee,
			ErrInvalidTransactionNonce,
			ErrValidationFailure,
			ErrNonAccount,
			ErrDuplicateTx,
			ErrUnsupportedTxVersion,
		)
	}
	return &output, nil
}

func (provider *Provider) AddDeclareTransaction(ctx context.Context, declareTransaction AddDeclareTxnInput) (*AddDeclareTransactionResponse, error) {
	var result AddDeclareTransactionResponse
	if err := do(ctx, provider.c, "starknet_addDeclareTransaction", &result, declareTransaction); err != nil {
		if unexpectedErr, ok := isErrUnexpectedError(err); ok {
			return nil, unexpectedErr
		}
		return nil, tryUnwrapToRPCErr(
			err,
			ErrClassAlreadyDeclared,
			ErrCompilationFailed,
			ErrCompiledClassHashMismatch,
			ErrInsufficientAccountBalance,
			ErrInsufficientMaxFee,
			ErrInvalidTransactionNonce,
			ErrValidationFailure,
			ErrNonAccount,
			ErrDuplicateTx,
			ErrContractClassSizeTooLarge,
			ErrUnsupportedTxVersion,
			ErrUnsupportedContractClassVersion,
		)
	}
	return &result, nil
}

// AddDeployAccountTransaction manages the DEPLOY_ACCOUNT syscall
func (provider *Provider) AddDeployAccountTransaction(ctx context.Context, deployAccountTransaction DeployAccountTxn) (*AddDeployAccountTransactionResponse, error) {
	var result AddDeployAccountTransactionResponse
	if err := do(ctx, provider.c, "starknet_addDeployAccountTransaction", &result, deployAccountTransaction); err != nil {
		if unexpectedErr, ok := isErrUnexpectedError(err); ok {
			return nil, unexpectedErr
		}

		return nil, tryUnwrapToRPCErr(
			err,
			ErrInsufficientAccountBalance,
			ErrInsufficientMaxFee,
			ErrInvalidTransactionNonce,
			ErrValidationFailure,
			ErrNonAccount,
			ErrClassHashNotFound,
			ErrDuplicateTx,
			ErrUnsupportedTxVersion,
		)
	}
	return &result, nil
}
