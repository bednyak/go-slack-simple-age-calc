# Go-Slack-Bot-File-Uploader

### Preparation

1. Create an account in Slack
2. Go to https://api.slack.com/apps?new_app=1 and create slack app ![setup-1.png](./docs/images/setup-1.png) ![setup-2.png](./docs/images/setup-2.png) ![setup-3.png](./docs/images/setup-3.png)
3. Enable sockets ![setup-4.png](./docs/images/setup-4.png) ![setup-5.png](./docs/images/setup-5.png) ![setup-16.png](./docs/images/setup-16.png)
4. Get Slack application token and use it in SLACK_APP_TOKEN
5. Enable event subscriptions ![setup-6.png](./docs/images/setup-6.png)
6. Configure Slack Event Subscriptions ![setup-17.png](./docs/images/setup-17.png)
7. Configure Slack OAuth and Permissions with adding scopes ![setup-7.png](./docs/images/setup-7.png) ![setup-8.png](./docs/images/setup-8.png) ![setup-9.png](./docs/images/setup-9.png)
8. Get Slack Bot User OAuth Token and set SLACK_BOT_TOKEN in config file ![setup-11.png](./docs/images/setup-11.png)
10. Add file-bot to the channel in Slack ![setup-14.png](./docs/images/setup-14.png) ![setup-15.png](./docs/images/setup-15.png)

### Run

```
go run main.go
```

Go to Slack channel where age-bot was added and ask question with next structure
```
@age-bot My year of birth is {{YOUR_YEAR}}
```
![result.png](./docs/images/result.png)