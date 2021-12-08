# Auto bras dial and dynamic DNS tool for Nanjing University

The network system of NJU has a mystery flow control mechanism. Some times
it may be triggered and the bandwidth and visitable domains of a certain machine
will be limited.
The limitation can be bypassed  by change a machine's MAC address. But if we do
this, the auto-dial function of nju bras will not work (because it identifies a machine
by it's MAC address).
So I developed this tool to auto dial the bras to help you stay online even if 
you changed your mac address remotely (by RDC or some other tools).

It can also help you maintain a dynamic dns to your machine and notify
something via tools like bark.

## Usage
- prepare a config like `template.yaml`
- `njuddns -c /path/to/your/config`
- check the log

## Modules
### njubras: a bras dail/status query tool
```go
func NjuBrasTest() {
    loginStatus, err := DoLogin(username, password)
    logoutStatus, err := DoLogout()
    err = AcquirePortalStatus()
}
```
### notify: a notify tool
#### bark: notify via bark
[Bark](https://github.com/Finb/Bark) is an iOS App which allows you to push customized notifications to your iPhone

### provider: dynamic dns providers
#### cloudflare
Using `cloudflare-go` API to dynamically update A record


## TODO & Roadmaps
- Use new logging system to notify correctly
- Support more ddns and notify tools
- Support login via unified login of NJU

## Contribute
This project is still under very early development, PR and issue is welcomed. 