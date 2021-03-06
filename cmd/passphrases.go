// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import "github.com/spf13/viper"

// getStorePassphrases() fetches the store passphrase supplied by the user.
func getStorePassphrase() string {
	return viper.GetString("store-passphrase")
}

// getWalletPassphrases() fetches the wallet passphrase supplied by the user.
func getWalletPassphrase() string {
	return viper.GetString("wallet-passphrase")
}

// getPassphrases() fetches the passphrases supplied by the user.
func getPassphrases() []string {
	return viper.GetStringSlice("passphrase")
}

// getPassphrase fetches the passphrase supplied by the user.
func getPassphrase() string {
	passphrases := getPassphrases()
	assert(len(passphrases) != 0, "passphrase is required")
	assert(len(passphrases) == 1, "multiple passphrases supplied; cannot continue")
	return passphrases[0]
}
