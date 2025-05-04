#!/bin/bash

echo "🛑 Goアプリケーションコンテナを停止しています..."

# コンテナが存在するか確認
if podman ps -a | grep -q "go-server"; then
  echo "🔍 go-serverコンテナを停止・削除しています..."
  podman stop go-server
  podman rm go-server
  echo "✅ go-serverコンテナを停止・削除しました"
else
  echo "ℹ️ go-serverコンテナは実行されていません"
fi

echo "📋 実行中のコンテナ一覧:"
podman ps 