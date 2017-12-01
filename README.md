# GitHub-Notify

GitHub-Notify is a very simple tool for sending a GitHub status notification via a GitHub Pull Request Webhook.

It is designed to function as a 12-factor app, receiving configuration via environment variables. To keep things simple, it is rigid in what it allows you to set. More complex bots might want to customize this a little more.

This tool uses the [GitHub API]](https://api.github.com/) to update a pull request's status.

Before you can use this tool, you need to log into your GitHub account and configure this.

## Usage

Running `github-notify` in a shell prompt goes like this:

```console
$ cat << EOF > .env
GITHUB_ACCESS_TOKEN=xxxxxxxx
GITHUB_OWNER=org
GITHUB_REPO=repo
GITHUB_REF=commit
GITHUB_STATE=success
GITHUB_TARGET_URL=https://example.com
GITHUB_DESCRIPTION="The build succeeded!"
GITHUB_CONTEXT=brigade
EOF
$ source .env
$ github-notify
```

Running the Docker container goes like this:

```console
$ cat << EOF > .env
GITHUB_ACCESS_TOKEN=xxxxxxxx
GITHUB_OWNER=org
GITHUB_REPO=repo
GITHUB_REF=commit
GITHUB_STATE=success
GITHUB_TARGET_URL=https://example.com
GITHUB_DESCRIPTION="The build succeeded!"
GITHUB_CONTEXT=brigade
EOF
$ docker run --env-file=.env bacongobbler/github-notify
```

### In Brigade

You can easily use this inside of brigade hooks.

```javascript
const {events, Job} = require("brigadier")

events.on("pull_request", (e, p) => {
  var payload = JSON.parse(e.payload)

  if (e.provider == "github") {

    var gh = new Job("github-notify")
    gh.image = "bacongobbler/github-notify:latest"
    gh.env = {
      // It's best to store the github auth token in a project's secrets.
      GITHUB_ACCESS_TOKEN: p.secrets.GITHUB_ACCESS_TOKEN,
      GITHUB_OWNER: p.secrets.GITHUB_OWNER,
      GITHUB_REPO: p.secrets.GITHUB_REPO,
      GITHUB_REF: p.secrets.GITHUB_REF,
      GITHUB_STATE: p.secrets.GITHUB_STATE,
      GITHUB_TARGET_URL: p.secrets.GITHUB_TARGET_URL,
      GITHUB_DESCRIPTION: p.secrets.GITHUB_DESCRIPTION,
      GITHUB_CONTEXT: p.secrets.GITHUB_CONTEXT
    }
    gh.run()
  }
})
```

## Build It

Configure:

```
make bootstrap
```

Compile:

```
make build
```

Publish to DockerHub

```
make docker-build docker-push
```
