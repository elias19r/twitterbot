twitterbot
==========

This is a simple twitter bot implemented in Golang and currently relying on
[github.com/ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)
lib to interact with Twitter API v1.1.

__NOTE__: as of June, 2018, Twitter plans to shutdown User streams which this
bot relies on and replace it by [Account Activity API](https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/overview)
via webhooks.

Config, build, and run
----------------------

Set your bot username, replies and behavior settings in file `./internal/config/config.go`

Then build the twitterbot binary:

```sh
go build ./cmd/twitterbot
```

Set the following Twitter API keys in your environment variables. Such keys you
get from your app page at https://apps.twitter.com/.

__NOTE__: in order to receive DM messages, your keys' access level must be:
"Read, write, and direct messages".

```sh
export CONSUMER_KEY=your-app-consumer-key
export CONSUMER_SECRET=your-app-consumer-secret
export ACCESS_TOKEN=your-app-access-token
export ACCESS_TOKEN_SECRET=your-app-access-token-secret
```

And then run the bot:

```sh
./twitterbot
```

When a `console >` prompts, type `bhvr startall` to start all behaviors or
`help` for more info.

Build docker image
------------------

Create a local `.env` file with your Twitter API keys:

```sh
echo CONSUMER_KEY=your-app-consumer-key >> .env
echo CONSUMER_SECRET=your-app-consumer-secret >> .env
echo ACCESS_TOKEN=your-app-access-token >> .env
echo ACCESS_TOKEN_SECRET=your-app-access-token-secret >> .env
```

Simply run the script:

```sh
./build-docker-image.sh
```

It will compile a linux-amd64 static binary, without debug info, and copy it
into an image based on `amd64/alpine:latest` Linux according to `Dockerfile`.

In order to run the container, execute:

```sh
docker run -it twitterbot
```

TODO
----

- Promote "SDV da Semana" of top 10 user with best ratio following/followers.
- Replace User streams with a webhook implementation.
- console config|cfg set NAME VALUES , config|cfg ls