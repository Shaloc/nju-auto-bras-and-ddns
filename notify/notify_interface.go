package notify

type NotifierConfig struct {
	ApiUrl string
	Topic string
}

type INotifier interface {
	Initialize(conf *NotifierConfig)
	DoNotify(message string)
}
