package common

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// AppName ..
const AppName = "release"

// UploadURLUseless ..
const UploadURLUseless = "{?name,label}"

// APIGithub ..
const APIGitHub = "https://api.github.com"

// HostGitHub ..
const HostGitHub = "https://github.com"

// Config ..
const Config = ".release.yml"

// EnvPrefix ..
const EnvPrefix = "RELEASE"

// APIRepoURL ..
func APIRepoURL() string {
	return fmt.Sprintf("%s/repos/%s/%s",
		APIGitHub, viper.GetString("Username"), viper.GetString("Repo"))
}

// HostRepoURL ..
func HostRepoURL() string {
	return fmt.Sprintf("%s/%s/%s",
		HostGitHub, viper.GetString("Username"), viper.GetString("Repo"))
}

// OAuthClientQueryString ..
func OAuthClientQueryString() string {
	var ClientID = viper.GetString("ClientID")
	var ClientSecret = viper.GetString("ClientSecret")
	if ClientID == "" || ClientSecret == "" {
		return ""
	}
	return fmt.Sprintf("?client_id=%s&client_secret=%s", ClientID, ClientSecret)
}

// Timeout ..
func Timeout() time.Duration {
	return time.Second * time.Duration(viper.GetInt("Runtime.Timeout"))
}

// Token ..
func Token() string {
	return fmt.Sprintf("Bearer %s", viper.GetString("TOKEN"))
}
