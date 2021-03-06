// Copyright 2020 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package config // import "github.com/wabarc/wayback/config"

const (
	defDebug    = false
	defLogTime  = true
	defIPFSHost = "127.0.0.1"
	defIPFSPort = 4001
	defIPFSMode = "pinner"

	defOverTor = false

	defEnabledIA = false
	defEnabledIS = false
	defEnabledIP = false
	defEnabledPH = false

	defTelegramToken   = ""
	defTelegramChannel = ""
	defGitHubToken     = ""
	defGitHubOwner     = ""
	defGitHubRepo      = ""

	defTorPrivateKey = ""
	defTorLocalPort  = 0
	defTorrcFile     = "/etc/tor/torrc"
)

var (
	defTorRemotePorts = []int{80}
)

type Options struct {
	debug   bool
	logTime bool
	overTor bool

	ipfs     *ipfs
	slots    map[string]bool
	telegram *telegram
	github   *github
	tor      *tor
}

type ipfs struct {
	host string
	port uint
	mode string
}

type slots struct {
	ia bool
	is bool
	ip bool
	ph bool
}

type telegram struct {
	token   string
	channel string
}

type github struct {
	token string
	owner string
	repo  string
}

type tor struct {
	pvk string

	localPort   int
	remotePorts []int
	torrcFile   string
}

// NewOptions returns Options with default values.
func NewOptions() *Options {
	opts := &Options{
		debug:   defDebug,
		logTime: defLogTime,
		overTor: defOverTor,
		ipfs: &ipfs{
			host: defIPFSHost,
			port: defIPFSPort,
			mode: defIPFSMode,
		},
		slots: map[string]bool{
			SLOT_IA: defEnabledIA,
			SLOT_IS: defEnabledIS,
			SLOT_IP: defEnabledIP,
			SLOT_PH: defEnabledPH,
		},
		telegram: &telegram{
			token:   defTelegramToken,
			channel: defTelegramChannel,
		},
		github: &github{
			token: defGitHubToken,
			owner: defGitHubOwner,
			repo:  defGitHubRepo,
		},
		tor: &tor{
			pvk:         defTorPrivateKey,
			localPort:   defTorLocalPort,
			remotePorts: defTorRemotePorts,
			torrcFile:   defTorrcFile,
		},
	}

	return opts
}

// HasDebugMode returns true if debug mode is enabled.
func (o *Options) HasDebugMode() bool {
	return o.debug
}

// LogTime returns if the time should be displayed in log messages.
func (o *Options) LogTime() bool {
	return o.logTime
}

// IPFSHost returns the host of IPFS daemon service.
func (o *Options) IPFSHost() string {
	return o.ipfs.host
}

// IPFSPort returns the port of IPFS daemon service.
func (o *Options) IPFSPort() uint {
	return o.ipfs.port
}

// IPFSMode returns the mode to using IPFS.
func (o *Options) IPFSMode() string {
	return o.ipfs.mode
}

// UseTor returns whether to use the Tor proxy when snapshot webpage.
func (o *Options) UseTor() bool {
	return o.overTor
}

// Slots returns configurations of wayback service, e.g. Internet Archive
func (o *Options) Slots() map[string]bool {
	return o.slots
}

// TelegramToken returns the token of Telegram Bot.
func (o *Options) TelegramToken() string {
	return o.telegram.token
}

// TelegramChannel returns the Telegram Channel name.
func (o *Options) TelegramChannel() string {
	return o.telegram.channel
}

// PublishToChannel returns whether to publish results to Telegram Channel.
func (o *Options) PublishToChannel() bool {
	return o.telegram.token != "" && o.telegram.channel != ""
}

// GitHubToken returns the personal access token of GitHub account.
func (o *Options) GitHubToken() string {
	return o.github.token
}

// GitHubOwner returns the user id of GitHub account.
func (o *Options) GitHubOwner() string {
	return o.github.owner
}

// GitHubRepo returns the GitHub repository which to publish results.
func (o *Options) GitHubRepo() string {
	return o.github.repo
}

// PublishToIssues returns whether to publish results to GitHub issues.
func (o *Options) PublishToIssues() bool {
	return o.github.token != "" && o.github.owner != "" && o.github.repo != ""
}

// TorPrivKey returns the private key of Tor service.
func (o *Options) TorPrivKey() string {
	return o.tor.pvk
}

// TorLocalPort returns the local port to a TCP listener on.
func (o *Options) TorLocalPort() int {
	return o.tor.localPort
}

// TorRemotePorts returns the remote ports to serve the Tor hidden service on.
func (o *Options) TorRemotePorts() []int {
	return o.tor.remotePorts
}

// TorrcFile returns path of the torrc file to set on start Tor Hidden Service.
func (o *Options) TorrcFile() string {
	return o.tor.torrcFile
}
