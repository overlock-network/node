#!/bin/bash

GENESIS_FILE="$1"
CREATOR_ADDRESS="$2"

CW721_URL="https://github.com/overlock-network/cw-nfts/releases/download/v0.18.5/cw721_base.wasm"
CW2981_URL="https://github.com/overlock-network/cw-nfts/releases/download/v0.18.5/cw2981_configuration.wasm"

CW721_FILE="cw721_base.wasm"
CW2981_FILE="cw2981_royalties.wasm"

if [ -z "$GENESIS_FILE" ]; then
    echo "Error: Path to genesis.json not provided."
    exit 1
fi

if [ -z "$CREATOR_ADDRESS" ]; then
    echo "Error: Creator address not provided."
    exit 1
fi

download_wasm_if_missing() {
    local url=$1
    local file=$2

    if [ ! -f "$file" ]; then
        echo "Downloading $file..."
        wget -q -O "$file" "$url"
        if [ $? -ne 0 ]; then
            echo "Error downloading $file. Aborting."
            exit 1
        fi
    fi
}

wasm_entry_json() {
    local file=$1
    local creator=$2
    local code_id=$3

    local code_hash_hex=$(sha256sum "$file" | awk '{print $1}')
    local code_hash_base64=$(echo -n "$code_hash_hex" | xxd -r -p | base64 -w 0)
    local code_bytes_base64=$(base64 -w 0 "$file")

    cat <<EOF
{
  "code_id": "$code_id",
  "code_info": {
    "code_hash": "$code_hash_base64",
    "creator": "$creator",
    "instantiate_config": {
      "permission": "Everybody",
      "addresses": []
    }
  },
  "code_bytes": "$code_bytes_base64"
}
EOF
}

download_wasm_if_missing "$CW721_URL" "$CW721_FILE"
download_wasm_if_missing "$CW2981_URL" "$CW2981_FILE"

LAST_CODE_ID=$(jq -r '.app_state.wasm.sequences[]? | select(.id_key == "BGxhc3RDb2RlSWQ=") | .value' "$GENESIS_FILE")
if [ -z "$LAST_CODE_ID" ]; then
    CODE_ID_1=1
else
    CODE_ID_1=$((LAST_CODE_ID + 1))
fi
CODE_ID_2=$((CODE_ID_1 + 1))
NEW_SEQ=$((CODE_ID_2 + 1))

wasm_entry_json "$CW721_FILE" "$CREATOR_ADDRESS" "$CODE_ID_1" > "$GENESIS_FILE.tmp_code1"
wasm_entry_json "$CW2981_FILE" "$CREATOR_ADDRESS" "$CODE_ID_2" > "$GENESIS_FILE.tmp_code2"

jq \
  --slurpfile code1 "$GENESIS_FILE.tmp_code1" \
  --slurpfile code2 "$GENESIS_FILE.tmp_code2" \
  --argjson new_seq "$NEW_SEQ" \
  '
  .app_state.wasm.codes += ($code1 + $code2)
  | .app_state.wasm.sequences = [
      {
        id_key: "BGxhc3RDb250cmFjdElk",
        value: ($new_seq | tostring)
      },
      {
        id_key: "BGxhc3RDb2RlSWQ=",
        value: ($new_seq | tostring)
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

rm -f "$GENESIS_FILE.tmp_code1" "$GENESIS_FILE.tmp_code2"
rm -f "$CW721_FILE" "$CW2981_FILE"

echo "Genesis updated with CW721 and CW2981"
