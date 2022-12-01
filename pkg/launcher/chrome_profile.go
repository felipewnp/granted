package launcher

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"runtime"

	"github.com/common-fate/granted/pkg/browser"
)

type ChromeProfile struct {
	// ExecutablePath is the path to the Chrome binary on the system.
	ExecutablePath string
	// UserDataPath is the path to the Chrome user data directory,
	// which we override to put Granted profiles in a specific folder
	// for easy management.
	UserDataPath string

	BrowserType string
}

func (l ChromeProfile) LaunchCommand(url string, profile string) []string {
	profileName, _ := FindBrowserProfile(profile, l.BrowserType)
	// profileName := chromeProfileName(profile)

	return []string{
		l.ExecutablePath,
		// "--args",
		// "--user-data-dir=" + l.UserDataPath,
		"--profile-directory=" + profileName,
		// "--no-first-run",
		"--no-default-browser-check",
		url,
	}
}

//Todo: is this still needed
// func chromeProfileName(profile string) string {

// 	h := fnv.New32a()
// 	h.Write([]byte(profile))

// 	hash := fmt.Sprint(h.Sum32())
// 	return hash
// }

var BravePathMac = "Library/Application Support/BraveSoftware/Brave-Browser/Local State"
var BravePathLinux = ".config/brave-browser/Local State"
var BravePathWindows = `AppData\Local\BraveSoftware\Brave-Browser\Local State`

var ChromePathMac = "Library/Application Support/Google/Chrome/Local State"
var ChromePathLinux = ".config/google-chrome/Local State"
var ChromePathWindows = `AppData\Local\Google\Chrome\User Data/Local State`

var EdgePathMac = `Library/Application Support/Microsoft\ Edge/Local State`
var EdgePathLinux = ".config/microsoft-edge/Local State"
var EdgePathWindows = `AppData\Local\Microsoft Edge\User Data/Local State`

var ChromiumPathMac = "Library/Application Support/Chromium/Local State"
var ChromiumPathLinux = ".config/chromium/Local State"
var ChromiumPathWindows = `AppData\Local\Chromium\User Data/Local State`

func FindBrowserProfile(profile string, browserType string) (string, error) {
	//open Local State file for browser

	//work out which chromium browser we are using

	//todo: make this os compatible
	stateFile, err := getLocalStatePath(browserType)
	if err != nil {
		return "", err
	}

	//read the state file
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return "", err
	}

	//the Local State json blob is a bunch of map[string]interfaces which makes it difficult to unmarshal
	var f map[string]interface{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		return "", err
	}

	//grab the profiles out from the json blob
	profiles := f["profile"].(map[string]interface{})
	//can this be done cleaner with a conversion into a struct?
	for profileName, profileObj := range profiles["info_cache"].(map[string]interface{}) {
		//if the profile name is the same as the profile name we are assuming then we want to use the same profile
		if profileObj.(map[string]interface{})["name"] == profile {
			return profileName, nil
		}

	}

	return profile, nil
}

func getLocalStatePath(browserType string) (stateFile string, err error) {
	stateFile, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}
	switch runtime.GOOS {
	case "windows":
		switch browserType {
		case browser.ChromeKey:
			stateFile = path.Join(stateFile, ChromePathWindows)

		case browser.BraveKey:
			stateFile = path.Join(stateFile, BravePathWindows)

		case browser.EdgeKey:
			stateFile = path.Join(stateFile, EdgePathWindows)

		case browser.ChromiumKey:
			stateFile = path.Join(stateFile, ChromiumPathWindows)
		}

	case "darwin":
		switch browserType {
		case browser.ChromeKey:
			stateFile = path.Join(stateFile, ChromePathMac)

		case browser.BraveKey:
			stateFile = path.Join(stateFile, BravePathMac)

		case browser.EdgeKey:
			stateFile = path.Join(stateFile, EdgePathMac)

		case browser.ChromiumKey:
			stateFile = path.Join(stateFile, ChromiumPathMac)
		}

	case "linux":
		switch browserType {
		case browser.ChromeKey:
			stateFile = path.Join(stateFile, ChromePathLinux)

		case browser.BraveKey:
			stateFile = path.Join(stateFile, BravePathLinux)

		case browser.EdgeKey:
			stateFile = path.Join(stateFile, EdgePathLinux)

		case browser.ChromiumKey:
			stateFile = path.Join(stateFile, ChromiumPathLinux)
		}

	default:
		return "", errors.New("os not supported")
	}
	return stateFile, nil
}

// type Profile struct {
// 	ActiveTime                   float64 `json:"active_time"`
// 	AvatarIcon                   string  `json:"avatar_icon"`
// 	BackgroundApps               bool    `json:"background_apps"`
// 	ForceSigninProfileLocked     bool    `json:"force_signin_profile_locked"`
// 	GaiaID                       string  `json:"gaia_id"`
// 	IsConsentedPrimaryAccount    bool    `json:"is_consented_primary_account"`
// 	IsEphemeral                  bool    `json:"is_ephemeral"`
// 	IsUsingDefaultAvatar         bool    `json:"is_using_default_avatar"`
// 	IsUsingDefaultName           bool    `json:"is_using_default_name"`
// 	ManagedUserID                string  `json:"managed_user_id"`
// 	MetricsBucketIndex           int     `json:"metrics_bucket_index"`
// 	Name                         string  `json:"name"`
// 	SigninWithCredentialProvider bool    `json:"signin.with_credential_provider"`
// 	UserName                     string  `json:"user_name"`
// }
