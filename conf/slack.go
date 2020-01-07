package conf

type Slack struct {
	WebHookURL string `default:"webhook_url_goes_here" envconfig:"WEBHOOK_URL"`
}
