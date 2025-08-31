# basic_webhook

This is a basic webhook container for Mythic v3.0.0. This container connects up to RabbitMQ queues for webhooks and posts those messages and additional data to Slack and Discord channels.

## How to install an agent in this format within Mythic

When it's time for you to test out your install or for another user to install your c2 profile, it's pretty simple. Within Mythic you can run the `mythic-cli` binary to install this in one of three ways:

* `sudo ./mythic-cli install github https://github.com/user/repo` to install the main branch
* `sudo ./mythic-cli install github https://github.com/user/repo branchname` to install a specific branch of that repo
* `sudo ./mythic-cli install folder /path/to/local/folder/cloned/from/github` to install from an already cloned down version of an agent repo

Now, you might be wondering _when_ should you or a user do this to properly add your profile to their Mythic instance. There's no wrong answer here, just depends on your preference. The three options are:

* Mythic is already up and going, then you can run the install script and just direct that profile's containers to start (i.e. `sudo ./mythic-cli start profileName`.
* Mythic is already up and going, but you want to minimize your steps, you can just install the profile and run `sudo ./mythic-cli start`. That script will first _stop_ all of your containers, then start everything back up again. This will also bring in the new profile you just installed.
* Mythic isn't running, you can install the script and just run `sudo ./mythic-cli start`. 

## Are you looking to set the actual webhook url to use? You have three options:
* To just change the URL for one operation, you can click the operation name at the top of Mythic's UI and click the edit for your current operation. You'll be able to add the webhook URL there
* If you want to change it more generally, you can set the .env value `WEBHOOK_DEFAULT_URL` (then make sure you do a `sudo ./mythic-cli start basic_webhook` so that the change is pulled in). When you edit that, you'll also see various `WEBHOOK_DEFAULT_*_CHANNEL` so you can change the default channel used with your webhook for alerts, callbacks, feedback, startup, and custom notifications
* You can also edit the code in the basic_webhook where it sends the URL to manually add it there (if you do this, you need to make sure your change is pulled in, so you'd need to set `sudo ./mythic-cli config set basic_webhook_use_build_context true` and `sudo ./mythic-cli config set basic_webhook_use_volume false` then `sudo ./mythic-cli build basic_webhook`
