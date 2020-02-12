# Cloudflare DNS Record Updater
## Welcome
With this tool, you can automatically replace of selected DNS on Cloudflare for systems like servers with dynamic IP. 

## Usage
1. Build project (you can use taskfile).
2. Run it once and edit the generated config file.
3. Run again and see how it works.

You can add this tool to crontab, rc.local, etc. on linux or scheduled task on windows for run at certain times.

## Config File
    {
        "x_auth_email": "", //EMAIL ADRESS ON CLOUDFLARE
        "x_auth_key": "", //CLOUDFLARE AUTH KEY
        "domain": "", //DOMAIN WHICH DNS REGISTRATION WILL BE CHANGED
        "record": "", //DNS RECORD NAME(www, etc.)
        "proxied": false //CLOUDFLARE PROXY
    }

## Copyright and License
**Cloudflare DNS Record Updater** is licensed under the MIT license.
