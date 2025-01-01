#!/bin/bash

# 変数の設定
USER_POOL_ID="ap-northeast-1_ZigqohuSQ"
CLIENT_ID="qtnut91bg07iuq16ib6b2mljc"
USERNAME="87347ac8-10f1-7092-bf21-67483b88fa06"
PASSWORD="Reservation_0"
REGION="ap-northeast-1"
PROFILE="private"

# initiate-auth を実行して認証情報を取得
echo "Executing initiate-auth..."
INITIATE_AUTH_RESPONSE=$(aws cognito-idp initiate-auth \
  --auth-flow USER_PASSWORD_AUTH \
  --client-id "$CLIENT_ID" \
  --auth-parameters USERNAME="$USERNAME",PASSWORD="$PASSWORD" \
  --region "$REGION" \
  --profile "$PROFILE" 2>&1)

# コマンドが失敗した場合のエラーチェック
if [ $? -ne 0 ]; then
  echo "ERROR: Failed to execute initiate-auth."
  echo "Response:"
  echo "$INITIATE_AUTH_RESPONSE"
  exit 1
fi

# レスポンスを確認
echo "Initiate Auth Response:"
# echo "$INITIATE_AUTH_RESPONSE"

# IdToken を抽出
ID_TOKEN=$(echo "$INITIATE_AUTH_RESPONSE" | jq -r '.AuthenticationResult.IdToken')

# IdToken の存在を確認
if [ -z "$ID_TOKEN" ] || [ "$ID_TOKEN" == "null" ]; then
  echo "ERROR: Failed to retrieve IdToken."
  echo "Complete Response:"
  echo "$INITIATE_AUTH_RESPONSE"
  exit 1
fi

echo $ID_TOKEN

