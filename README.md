# Yet Another Cloudflare DDNS Updater (yacddns)

Updates a Cloudflare DNS record with your current IP address. There are tons of these already, but it's so simple I wrote my own.

## Features
- Single executable, no runtime dependencies. Copy the binary and go!
- All config applied with command line switches
- "Single-shot" run once command, suitable for use with Cron
- Check multiple sources for the current IP address, only update if both match
- Stateless, no files created or modified

## Usage
```
Usage of ./yacddns:
  -apikey string
    	Cloudflare API key, find it in your profile
  -email string
    	Cloudflare account email
  -record string
    	Name of the record to update, e.g. www.example.com
  -zone string
    	Name of the cloudflare zone, e.g. example.com
```

## Limitations
- It currently just wallops the cloudflare API every run. It could be a little nicer, and only touch cloudflare if the IP changes but this requires saving the IP somewhere between runs, and 3 API calls even once a minute is not a lot all things considered. (For reference, Cloudflare's rate limit is set at 1200 every 5 mins)
- Only expects to update 1 A record. If you need to update multiple this should be easily modified.
- Only looks for and updates A records. No ipv6 support yet....
