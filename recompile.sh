#!/bin/bash
if [[ "$EUID" -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

# EDIT BELOW
# EXAMPLE: /home/kali/source/basic_webhook/C2_Profiles/webhook
WEBHOOKPATH="/path/to/C2_Profiles/webhook"  # Full path to Payload_Type/webhook, see my example
# EXAMPLE: /opt/Mythic
MYTHICPATH="/path/to/Mythic"

cd "/opt/Mythic" || exit 1
sudo "$MYTHICPATH"/mythic-cli remove webhook
sudo rm -rf "$MYTHICPATH"/InstalledServices/webhook
sudo cp -r "$WEBHOOKPATH" "$MYTHICPATH"/InstalledServices
sudo "$MYTHICPATH"/mythic-cli add webhook
sudo "$MYTHICPATH"/mythic-cli build webhook