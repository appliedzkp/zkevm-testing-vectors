package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm/runtime"

	"gethutil"
)

func main() {
	address := common.BytesToAddress([]byte{0xff})
	asm := gethutil.NewAssembly().Add(0xdeadbeef, 0xcafeb0ba).Sub(0xfaceb00c, 0xb0bacafe)
	contracts := []gethutil.Contract{{Address: address, Bytecode: asm.Bytecode}}

	logs, err := gethutil.TraceTx(address, nil, &runtime.Config{GasLimit: 100}, contracts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to trace tx, err: %v\n", err)
	}

	bytes, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal logs, err: %v\n", err)
	}

	fmt.Fprintln(os.Stdout, string(bytes))
}
