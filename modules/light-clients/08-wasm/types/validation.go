package types

import errorsmod "cosmossdk.io/errors"

const maxWasmSize = 3 * 1024 * 1024

// ValidateWasmCode valides that the size of the wasm code is in the allowed range
// and that the contents are of a wasm binary.
func ValidateWasmCode(code []byte) error {
	if len(code) == 0 {
		return ErrWasmEmptyCode
	}
	if len(code) > maxWasmSize {
		return ErrWasmCodeTooLarge
	}

	return nil
}

// MaxWasmByteSize returns the maximum allowed number of bytes for wasm bytecode
func MaxWasmByteSize() uint64 {
	return maxWasmSize
}

// ValidateWasmCodeHash validates that the code hash is of the correct length
func ValidateWasmCodeHash(codeHash []byte) error {
	lenCodeHash := len(codeHash)
	if lenCodeHash == 0 {
		return errorsmod.Wrap(ErrInvalidCodeHash, "code hash cannot be empty")
	}
	if lenCodeHash != 32 { // sha256 output is 256 bits long
		return errorsmod.Wrapf(ErrInvalidCodeHash, "expected length of 32 bytes, got %d", lenCodeHash)
	}

	return nil
}
