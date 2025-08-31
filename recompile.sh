#!/bin/bash
if [[ "$EUID" -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

# THIS FILE IS USED FOR LOCAL RECOMPILATION IN CASE YOU NEED IT. 
# SIMPLY MAKE SURE TO SPECIFY THE CORRECT PATHS AND THEN RUN IT AS ROOT.

# EDIT BELOW
# EXAMPLE: /home/kali/source/basic_webhook/C2_Profiles/webhook
WEBHOOKPATH="/path/to/C2_Profiles/webhook"  # Full path to Payload_Type/webhook, see my example
# EXAMPLE: /opt/Mythic
MYTHICPATH="/path/to/Mythic"

cd "$MYTHICPATH" || exit 1
sudo "$MYTHICPATH"/mythic-cli remove webhook
sudo rm -rf "$MYTHICPATH"/InstalledServices/webhook
sudo cp -r "$WEBHOOKPATH" "$MYTHICPATH"/InstalledServices
sudo "$MYTHICPATH"/mythic-cli add webhook
sudo "$MYTHICPATH"/mythic-cli build webhook