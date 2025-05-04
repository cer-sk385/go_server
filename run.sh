#!/bin/bash

# エラー時に停止
set -e

echo "🚀 Goアプリケーションをコンテナで実行します..."

# Podman Machineが起動しているか確認
echo "🔍 Podman Machineの状態を確認しています..."
if ! podman machine list | grep -q "Currently running"; then
  echo "🔄 Podman Machineを起動しています..."
  podman machine start || echo "既に起動している可能性があります。続行します..."
else
  echo "✅ Podman Machineは既に起動しています"
fi

# 既存のコンテナを確認して停止
echo "🔍 既存のgo-serverコンテナを確認しています..."
if podman ps -a | grep -q "go-server"; then
  echo "🧹 既存のgo-serverコンテナを削除しています..."
  podman rm -f go-server
fi

# ポート8080が使用中かどうか確認
echo "🔍 ポート8080の使用状況を確認しています..."
if lsof -i :8080 > /dev/null 2>&1; then
  echo "⚠️ ポート8080は既に使用されています。プロセスを終了します..."
  lsof -i :8080 -t | xargs kill
  sleep 1
fi

# イメージをビルド
echo "🏗️ Podmanイメージをビルドしています..."
podman build -t go-server .

# コンテナを起動
echo "🚀 コンテナを起動しています..."
podman run -d -p 8080:8080 --name go-server go-server

echo "✅ go-serverが起動しました！"
echo "📋 ログを表示するには: podman logs -f go-server"
echo "🌐 アプリにアクセスするには: http://localhost:8080"
echo "🛑 停止するには: podman stop go-server" 