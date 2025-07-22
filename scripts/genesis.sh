#!/bin/bash

GENESIS_FILE="$1"
CW2981_URL="https://github.com/public-awesome/cw-nfts/releases/download/v0.18.0/cw721_base.wasm"
WASM_FILE="cw721_base.wasm"
CREATOR_ADDRESS="$2"

if [ -z "$GENESIS_FILE" ]; then
    echo "Error: Path to genesis.json not provided."
    exit 1
fi

if [ -z "$CREATOR_ADDRESS" ]; then
    echo "Error: Creator address not provided."
    exit 1
fi

if [ ! -f "$WASM_FILE" ]; then
    wget -q -O "$WASM_FILE" "$CW2981_URL"
    if [ $? -ne 0 ]; then
        echo "Error downloading WASM file. Aborting."
        exit 1
    fi
fi

CODE_HASH_HEX=$(sha256sum "$WASM_FILE" | awk '{print $1}')
CODE_HASH_BASE64=$(echo -n "$CODE_HASH_HEX" | xxd -r -p | base64 -w 0)
CODE_BYTES_BASE64=$(base64 -w 0 "$WASM_FILE")

LAST_CODE_ID=$(jq -r '.app_state.wasm.sequences[]? | select(.id_key == "BGxhc3RDb2RlSWQ=") | .value' "$GENESIS_FILE")
if [ -z "$LAST_CODE_ID" ]; then
    NEW_CODE_ID=1
else
    NEW_CODE_ID=$((LAST_CODE_ID + 1))
fi

cat > "$GENESIS_FILE.tmp_code" <<EOF
{
  "code_id": "$NEW_CODE_ID",
  "code_info": {
    "code_hash": "$CODE_HASH_BASE64",
    "creator": "$CREATOR_ADDRESS",
    "instantiate_config": {
      "permission": "Everybody",
      "addresses": []
    }
  },
  "code_bytes": "$CODE_BYTES_BASE64"
}
EOF

jq \
  --slurpfile newcode "$GENESIS_FILE.tmp_code" \
  --argjson new_id "$NEW_CODE_ID" \
  '
  .app_state.wasm.codes += $newcode
  | .app_state.wasm.sequences = [
      {
        id_key: "BGxhc3RDb250cmFjdElk",
        value: ($new_id + 1 | tostring)
      },
      {
        id_key: "BGxhc3RDb2RlSWQ=",
        value: ($new_id + 1 | tostring)
      }
    ]
  | .app_state.wasm.params = {
      code_upload_access: {
        permission: "Everybody",
        addresses: []
      },
      instantiate_default_permission: "Everybody"
    }
  ' "$GENESIS_FILE" > "$GENESIS_FILE.tmp_final" && mv "$GENESIS_FILE.tmp_final" "$GENESIS_FILE"

rm "$GENESIS_FILE.tmp_code"
rm -f "$WASM_FILE"

echo "Genesis updated"
