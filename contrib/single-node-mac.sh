#!/bin/sh

rm -rf ~/.fanfuryd

set -o errexit -o nounset

# Build genesis file incl account for passed address
fanfuryd init --chain-id test test
fanfuryd keys add validator --keyring-backend="test"
fanfuryd add-genesis-account $(fanfuryd keys show validator -a --keyring-backend="test") 100000000000000ufury
fanfuryd gentx validator 100000000ufury --keyring-backend="test" --chain-id test
fanfuryd collect-gentxs

# Set proper defaults and change ports
sed -i '' 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.fanfuryd/config/config.toml
sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.fanfuryd/config/config.toml
sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.fanfuryd/config/config.toml
sed -i '' 's/index_all_keys = false/index_all_keys = true/g' ~/.fanfuryd/config/config.toml

# Start fanfury
fanfuryd start --pruning=nothing
