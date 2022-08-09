txscript
========

[![Build Status](https://travis-ci.org/brsuite/brond.png?branch=master)](https://travis-ci.org/brsuite/brond)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://godoc.org/github.com/brsuite/brond/txscript?status.png)](http://godoc.org/github.com/brsuite/brond/txscript)

Package txscript implements the brocoin transaction script language.  There is
a comprehensive test suite.

This package has intentionally been designed so it can be used as a standalone
package for any projects needing to use or validate brocoin transaction scripts.

## Brocoin Scripts

Brocoin provides a stack-based, FORTH-like language for the scripts in
the brocoin transactions.  This language is not turing complete
although it is still fairly powerful.  A description of the language
can be found at https://en.brocoin.it/wiki/Script

## Installation and Updating

```bash
$ go get -u github.com/brsuite/brond/txscript
```

## Examples

* [Standard Pay-to-pubkey-hash Script](http://godoc.org/github.com/brsuite/brond/txscript#example-PayToAddrScript)  
  Demonstrates creating a script which pays to a brocoin address.  It also
  prints the created script hex and uses the DisasmString function to display
  the disassembled script.

* [Extracting Details from Standard Scripts](http://godoc.org/github.com/brsuite/brond/txscript#example-ExtractPkScriptAddrs)  
  Demonstrates extracting information from a standard public key script.

* [Manually Signing a Transaction Output](http://godoc.org/github.com/brsuite/brond/txscript#example-SignTxOutput)  
  Demonstrates manually creating and signing a redeem transaction.

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the brsuite developers.  To
verify the signature perform the following:

- Download the public key from the Conformal website at
  https://opensource.conformal.com/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

Package txscript is licensed under the [copyfree](http://copyfree.org) ISC
License.
