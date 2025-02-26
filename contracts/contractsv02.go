package contracts

import (
	"context"
	_ "embed"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
)

type RPCProvider rpc.Provider

func (p *RPCProvider) declareAndWaitWithWallet(ctx context.Context, compiledClass []byte) (*DeclareOutput, error) {
	panic("This needs updated - Declare transactions have changed significantly over the past few rpc-spec updates")
	// provider := rpc.Provider(*p)
	// class := rpc.DeprecatedContractClass{}
	// if err := json.Unmarshal(compiledClass, &class); err != nil {
	// 	return nil, err
	// }
	// SenderAddress, err := utils.HexToFelt("0x01")
	// if err != nil {
	// 	return nil, err
	// }
	// tx, err := provider.AddDeclareTransaction(ctx, rpc.DeclareTxnV1{
	// 	Type:    "DECLARE",
	// 	MaxFee:  new(felt.Felt).SetUint64(10000),
	// 	Version: "0x01",
	// 	Nonce:   new(felt.Felt).SetUint64(1), // TODO: nonce handling

	// 	// ContractClass: class,
	// 	SenderAddress: SenderAddress, // TODO: constant devnet address
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// status, err := provider.WaitForTransaction(ctx, tx.TransactionHash, 8*time.Second)
	// if err != nil {
	// 	log.Printf("class Hash: %s\n", tx.ClassHash)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, err
	// }
	// if types.TransactionState(status.String()) == types.TransactionRejected {
	// 	log.Printf("class Hash: %s\n", tx.ClassHash)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, errors.New("declare rejected")
	// }
	// return &DeclareOutput{
	// 	classHash:       tx.ClassHash.String(),
	// 	transactionHash: tx.TransactionHash.String(),
	// }, nil
}

func (p *RPCProvider) deployAccountAndWaitNoWallet(ctx context.Context, classHash *felt.Felt, compiledClass []byte, salt string, inputs []string) (*DeployOutput, error) {
	panic("deployAccountAndWaitNoWallet needs updated")

	// provider := rpc.Provider(*p)
	// class := rpc.ContractClass{}
	// if err := json.Unmarshal(compiledClass, &class); err != nil {
	// 	return nil, err
	// }

	// fmt.Printf("classHash %v\n", classHash.String())

	// saltFelt, err := utils.HexToFelt(salt)
	// if err != nil {
	// 	return nil, err
	// }
	// inputsFelt, err := utils.HexArrToFelt(inputs)
	// if err != nil {
	// 	return nil, err
	// }

	// // TODO : Needs updated after DeployTransaction() was deprecated
	// tx, err := provider.AddDeployAccountTransaction(ctx, rpc.DeployAccountTransaction{
	// 	MaxFee:  new(felt.Felt).SetUint64(1),
	// 	Version: "0x01",
	// 	Nonce:   new(felt.Felt).SetUint64(2), // TODO: nonce handling

	// 	ContractAddressSalt: saltFelt,
	// 	ConstructorCalldata: inputsFelt,
	// 	ClassHash:           classHash,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// status, err := provider.WaitForTransaction(ctx, tx.TransactionHash, 8*time.Second)
	// if err != nil {
	// 	log.Printf("contract Address: %s\n", tx.ContractAddress)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, err
	// }
	// if types.TransactionState(status.String()) == types.TransactionRejected {
	// 	log.Printf("contract Address: %s\n", tx.ContractAddress)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, errors.New("deploy rejected")
	// }
	// return &DeployOutput{
	// 	ContractAddress: tx.ContractAddress.String(),
	// 	TransactionHash: tx.TransactionHash.String(),
	// }, nil
}

func (p *RPCProvider) deployAndWaitWithWallet(ctx context.Context, compiledClass []byte, salt string, inputs []string) (*DeployOutput, error) {
	panic("deployAndWaitWithWallet needs updated")

	// provider := rpc.Provider(*p)
	// class := rpc.ContractClass{}
	// if err := json.Unmarshal(compiledClass, &class); err != nil {
	// 	return nil, err
	// }
	// fmt.Println("a")

	// saltFelt, err := utils.HexToFelt(salt)
	// if err != nil {
	// 	return nil, err
	// }
	// inputsFelt, err := utils.HexArrToFelt(inputs)
	// if err != nil {
	// 	return nil, err
	// }

	// // TODO: use UDC via account
	// // TODO : Needs updated after DeployTransaction() was deprecated
	// tx, err := provider.AddDeployTransaction(ctx, rpc.BroadcastedDeployTxn{
	// 	DeployTransactionProperties: rpc.DeployTransactionProperties{
	// 		Version:             rpc.TransactionV1,
	// 		ContractAddressSalt: saltFelt,
	// 		ConstructorCalldata: inputsFelt,
	// 	},
	// 	ContractClass: class,
	// })
	// fmt.Println("b")
	// if err != nil {
	// 	return nil, err
	// }

	// status, err := provider.WaitForTransaction(ctx, tx.TransactionHash, 8*time.Second)
	// fmt.Println("c")
	// if err != nil {
	// 	log.Printf("contract Address: %s\n", tx.ContractAddress)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, err
	// }
	// if types.TransactionState(status.String()) == types.TransactionRejected {
	// 	log.Printf("contract Address: %s\n", tx.ContractAddress)
	// 	log.Printf("transaction Hash: %s\n", tx.TransactionHash)
	// 	return nil, errors.New("deploy rejected")
	// }
	// return &DeployOutput{
	// 	ContractAddress: tx.ContractAddress.String(),
	// 	TransactionHash: tx.TransactionHash.String(),
	// }, nil
}
