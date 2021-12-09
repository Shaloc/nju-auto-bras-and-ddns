package placeholder

import "nju_auto_ddns/notify"

type NotifierNil struct {
}

func (n *NotifierNil) Initialize(_ *notify.NotifierConfig) {
}

func (n *NotifierNil) DoNotify(_ string) {
}

