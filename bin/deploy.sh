#!/usr/bin/env bash

download_binary() {
  AUTHOR=reggles44
  REPO=semanti-go
  REPO_BASE_URL=https://api.github.com/repos/"$AUTHOR"/"$REPO"

  OS_TYPE=$(uname)
  ARCH=$(uname -m)
  FILE=semanti-go_"$OS_TYPE"_"$ARCH".tar.gz

  DOWNLOAD_URL=$(curl "$REPO_BASE_URL/releases/latest" | jq --arg FILE "$FILE" -r '.assets[] | select(.name | contains($FILE))| .browser_download_url')

  curl -sL "$DOWNLOAD_URL" -o semanti.tar.gz
  tar -xf semanti.tar.gz
  mv semanti /usr/local/bin
}

# setup_service() {
#   SERVICE="semanti.service"
#   SERVICE_FILE="/etc/systemd/system/$SERVICE"
#
#   # Create Service
#
#   # Remove the service file if it exists
#   #
#   if [ -f $SERVICE_FILE ]; then
#     rm $SERVICE_FILE
#   fi
#
#   cat >/etc/systemd/system/"$SERVICE" <<EOL
# [Unit]
# Description=Semanti GO
# After=network.target
# [Service]
# Type=simple
# ExecStart=/usr/local/bin/semanti
# [Install]
# WantedBy=multi-user.target
# EOL
#
#   sudo systemctl enable $SERVICE
#   sudo systemctl start $SERVICE
#
#   if systemctl -q is-active $SERVICE; then
#     echo "$SERVICE is up and running!"
#   else
#     echo "Failed to start $SERVICE."
#   fi
# }

download_binary
# setup_service
