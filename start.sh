#!/bin/bash

# Define variables
NODE_CMD="overlockd"
HOME_DIR="$HOME/.overlockd"
CONFIG_FILE="config.yml"
HASH_FILE="$HOME_DIR/config_hash"
KEYRING_BACKEND="test"
CHAIN_ID="overlock"
MONIKER="overlock-node"
STATE_BACKUP="$HOME_DIR/state_backup.json"
ACCOUNTS_FILE="$HOME_DIR/accounts_list.txt"
KEYRING_DIR="$HOME/.keyring-overlock"
RPC_PORT="26657"
GRPC_PORT="9090"
API_PORT="1317"
export NODE_CMD=$NODE_CMD CHAIN_ID=$CHAIN_ID KEYRING_DIR=$KEYRING_DIR ACCOUNT="bob"

# Compute hash of the config file
CURRENT_HASH=$(sha256sum $CONFIG_FILE | awk '{print $1}')

# Stop existing node if running
if pgrep -f "$NODE_CMD start" > /dev/null; then
    echo "Stopping existing node..."
    pkill -f "$NODE_CMD start"
    sleep 5
fi

# Reinitialize chain if config changed or no blockchain data exists
if [ "$CURRENT_HASH" != "$(cat $HASH_FILE 2>/dev/null)" ] || [ ! -d "$HOME_DIR/data" ]; then
    echo "⚠ Configuration changed. Reinitializing blockchain..."

    # Remove blockchain data and keyring
    rm -rf "$HOME_DIR" "$KEYRING_DIR"
    mkdir -p "$HOME_DIR"

    echo "" > "$ACCOUNTS_FILE"

    # Initialize the chain
    $NODE_CMD init $MONIKER --chain-id $CHAIN_ID --home $HOME_DIR >/dev/null 2>&1

    ACCOUNTS=$(yq -r '.accounts[].name' "$CONFIG_FILE" 2>/dev/null || true)


    ACCOUNTS_JSON="$HOME_DIR/accounts.json"
    echo "{}" > "$ACCOUNTS_JSON" >/dev/null 2>&1


    for ACCOUNT in $ACCOUNTS; do
        echo "🔑 Creating account: $ACCOUNT"
        
        # Ensure key exists
        if ! $NODE_CMD keys show $ACCOUNT --keyring-backend $KEYRING_BACKEND --home $KEYRING_DIR > /dev/null 2>&1; then
            ACCOUNT_DATA=$($NODE_CMD keys add $ACCOUNT --keyring-backend $KEYRING_BACKEND --output json --home $KEYRING_DIR)
            
            ACCOUNT_ADDRESS=$(echo $ACCOUNT_DATA | jq -r '.address')
            MNEMONIC=$(echo "$ACCOUNT_DATA" | jq -r '.mnemonic')

            echo "$ACCOUNT,$ACCOUNT_ADDRESS" >> "$ACCOUNTS_FILE"

            # Store in JSON
            jq --arg name "$ACCOUNT" \
            --arg address "$ACCOUNT_ADDRESS" \
            --arg mnemonic "$MNEMONIC" \
            '. + {($name): {"address": $address, "mnemonic": $mnemonic}}' \
            "$ACCOUNTS_JSON" > "$ACCOUNTS_JSON.tmp" && mv "$ACCOUNTS_JSON.tmp" "$ACCOUNTS_JSON"
        fi

        ACCOUNT_COINS=$(yq -r ".accounts[] | select(.name == \"$ACCOUNT\") | .coins | join(\",\")" $CONFIG_FILE)

        if [ -n "$ACCOUNT_COINS" ]; then
            echo "⚠ Funding account $ACCOUNT with $ACCOUNT_COINS..."

            echo "⚠ Funding account $ACCOUNT with $ACCOUNT_COINS..."
            $NODE_CMD genesis add-genesis-account $ACCOUNT_ADDRESS "$ACCOUNT_COINS" --append --home $HOME_DIR --keyring-backend $KEYRING_BACKEND 
            
        else
            echo "⚠ No coins specified for $ACCOUNT, skipping funding."
        fi
    done

    echo "📂 All account details saved in $ACCOUNTS_JSON"
    VALIDATOR=$(yq '.validators[0].name' $CONFIG_FILE)
    echo "⚖ Setting up validator: $VALIDATOR"

    VALIDATOR_BONDED=$(yq -r ".validators[0].bonded" $CONFIG_FILE)
    VALIDATOR_ADDRESS=$($NODE_CMD keys show $VALIDATOR -a --keyring-backend $KEYRING_BACKEND --home $KEYRING_DIR 2>/dev/null)

    if [ -n "$VALIDATOR_BONDED" ]; then
        echo "⚠ Checking if validator $VALIDATOR already exists in genesis..."
        if jq -e --arg addr "$VALIDATOR_ADDRESS" '.app_state.auth.accounts[] | select(.address == $addr)' "$HOME_DIR/config/genesis.json" > /dev/null; then
            echo "✅ Validator $VALIDATOR already exists. Using gentx with bonded amount."
            echo "📜 Generating gentx for $VALIDATOR"
            $NODE_CMD genesis gentx $VALIDATOR "$VALIDATOR_BONDED" --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --keyring-dir $KEYRING_DIR -y
        else
            echo "⚠ Validator $VALIDATOR does not exist. Adding to genesis and generating gentx."
            $NODE_CMD keys add $VALIDATOR --keyring-backend $KEYRING_BACKEND --home $KEYRING_DIR
            VALIDATOR_ADDRESS=$($NODE_CMD keys show $VALIDATOR -a --keyring-backend $KEYRING_BACKEND --home $KEYRING_DIR)
            $NODE_CMD genesis add-genesis-account $VALIDATOR_ADDRESS $VALIDATOR_BONDED --home $HOME_DIR --append
            echo "📜 Generating gentx for $VALIDATOR"
            $NODE_CMD genesis gentx $VALIDATOR "$VALIDATOR_BONDED" --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --keyring-dir $KEYRING_DIR -y
        fi
    else
        echo "❌ No bonded amount specified for validator $VALIDATOR, skipping gentx."
    fi

    sleep 3

    # Collect gentxs
    echo "📥 Collecting gentx transactions..."
    if [ -d "$HOME_DIR/config/gentx" ] && [ "$(ls -A $HOME_DIR/config/gentx)" ]; then
        $NODE_CMD genesis collect-gentxs --home $HOME_DIR >/dev/null 2>&1
    else
        echo "❌ No gentx files found! Initialization failed."
        exit 1
    fi

    echo "$CURRENT_HASH" | tee "$HASH_FILE" >/dev/null

else
    echo "✅ No changes detected. Using existing blockchain state."
    if [ -f "$STATE_BACKUP" ]; then
        echo "🔄 Restoring previous state..."
        $NODE_CMD unsafe-reset-all --home $HOME_DIR
        $NODE_CMD import < $STATE_BACKUP
    fi
fi

cp config/config.toml $HOME_DIR/config/config.toml

# Start the node with correct API bindings
echo "🚀 Starting Overlock node..."
$NODE_CMD start --rpc.laddr tcp://localhost:$RPC_PORT \
                --grpc.enable \
                --grpc.address localhost:$GRPC_PORT \
                --api.enable \
                --api.enabled-unsafe-cors \
                --api.address tcp://localhost:$API_PORT \
                --minimum-gas-prices 0stake \
                --home $HOME_DIR > overlock_node.log 2>&1 &

sleep 3  # Allow some time for process to start

go run ./faucet/faucet.go &

FAUCET_PID=$(pgrep -f "go run ./faucet/faucet.go")
NODE_PID=$(pgrep -f "$NODE_CMD start")

if [ -z "$NODE_PID" ]; then
    echo "❌ Failed to start Overlock node."
    exit 1
fi

# Display API endpoints
echo "🌍 Tendermint node: http://0.0.0.0:26657"
echo "🌍 Blockchain API: http://0.0.0.0:1317"
echo "🌍 Token faucet: http://0.0.0.0:4500"

# Monitor keypress for 'q' to stop
echo "Press 'q' to stop the node."
while true; do
    read -n 1 key
    if [[ $key == "q" ]]; then
        echo "🛑 Stopping Overlock node..."
        
        # Ensure the node process exists before killing
        if pgrep -f "$NODE_CMD start" > /dev/null; then
            kill $FAUCET_PID
            kill $NODE_PID
            echo "✅ Node stopped."
        else
            echo "⚠️ Node was already stopped."
        fi

        if pgrep -f "go run ./faucet/faucet.go" > /dev/null; then
            kill -9 $FAUCET_PID  
        fi

        if lsof -i :4500 > /dev/null; then
            fuser -k 4500/tcp  
        fi

        exit 0
    fi
done
