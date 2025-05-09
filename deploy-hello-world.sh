#!/bin/bash
# deploy-hello-world.sh - Script untuk deploy aplikasi Hello World ke PaaS

# Variabel konfigurasi
APP_NAME="hello-world"
BASE_PORT=9000
CONTAINERS=3
DESCRIPTION="Aplikasi Hello World contoh untuk PaaS Container"
PAAS_API="http://localhost:8080/api"

echo "Deploying Hello World app ke PaaS Container Platform..."

# 1. Buat aplikasi baru
echo "Membuat aplikasi baru..."
create_response=$(curl -s -X POST "${PAAS_API}/apps" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "'"${APP_NAME}"'",
    "containers": '"${CONTAINERS}"',
    "base_port": '"${BASE_PORT}"',
    "description": "'"${DESCRIPTION}"'"
  }')

# Ekstrak app_id dari response
APP_ID=$(echo $create_response | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)

if [ -z "$APP_ID" ]; then
  echo "Error: Gagal membuat aplikasi. Response: $create_response"
  exit 1
fi

echo "Aplikasi berhasil dibuat dengan ID: $APP_ID"

# 2. Deploy aplikasi
echo "Deploying aplikasi..."
deploy_response=$(curl -s -X POST "${PAAS_API}/apps/${APP_ID}/deploy")

deploy_status=$(echo $deploy_response | grep -o '"status":"[^"]*"' | cut -d'"' -f4)

if [ "$deploy_status" != "success" ]; then
  echo "Error: Gagal deploy aplikasi. Response: $deploy_response"
  exit 1
fi

echo "Aplikasi berhasil di-deploy!"

# 3. Update DNS records (opsional, jika DNS manager berjalan)
echo "Mengupdate DNS records..."
curl -s -X POST "http://localhost:8053/update" > /dev/null

# 4. Tampilkan informasi akses
echo ""
echo "Aplikasi Hello World berhasil di-deploy!"
echo "-----------------------------------------"
echo "Nama aplikasi: $APP_NAME"
echo "ID aplikasi: $APP_ID"
echo "Jumlah container: $CONTAINERS"
echo "Base port: $BASE_PORT"
echo ""
echo "Akses aplikasi:"
echo "- Via direct port: http://localhost:$BASE_PORT"
echo "- Via PaaS proxy: http://localhost/api/proxy/$APP_ID"
echo "- Via hostname (jika DNS setup): http://$APP_NAME.local"
echo ""
echo "API manajemen aplikasi:"
echo "- Informasi aplikasi: ${PAAS_API}/apps/${APP_ID}"
echo "- Daftar container: ${PAAS_API}/apps/${APP_ID}/containers"
echo "- Scale aplikasi: curl -X POST \"${PAAS_API}/apps/${APP_ID}/scale\" -H \"Content-Type: application/json\" -d '{\"containers\": 5}'"
echo "-----------------------------------------"
