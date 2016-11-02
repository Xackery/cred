//Package ini is a credentials config package that loads the ~/<path>/credentials file,
//or can be configured to load CREDENTIALS_FILE environment variable, if set.
package ini

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
)

type Profile struct {
	path        string
	filepath    string
	lastProfile string
	section     *ini.Section
}

//Load a new profile based on provided details
func Load(path string, profile string) (p *Profile, err error) {
	p = &Profile{}
	if err = p.Load(path, profile); err != nil {
		return
	}
	return
}

//Save the currently configured profile.
func Save(path string, filename string) (err error) {
	err = fmt.Errorf("Save is not yet implemented")
	return
}

//Load the credentials with provided profile context
func (p *Profile) Load(path string, profile string) (err error) {
	p.path = path
	if profile == "" {
		if p.lastProfile != "" {
			profile = p.lastProfile
		} else {
			profile = "default"
		}
	}
	if path == "" {
		path = ".profile"
	}
	p.lastProfile = profile

	if err = p.getFilepath(); err != nil {
		return
	}

	var config *ini.File
	if config, err = ini.Load(p.filepath); err != nil {
		return
	}
	if p.section, err = config.GetSection(profile); err != nil {
		return
	}
	return
}

//Add a custom key to the loaded credentials. Not recommended except for edge cases.
func (p *Profile) SetValue(key string, value string) (err error) {
	if err = p.prepareProfile(); err != nil {
		return
	}

	if p.section == nil {
		err = fmt.Errorf("profile is not loaded")
		return
	}

	if _, err = p.section.NewKey(key, value); err != nil {
		return
	}
	return
}

//Get a key from loaded credentials as string
func (p *Profile) String(key string) (value string) {
	var err error
	if err = p.prepareProfile(); err != nil {
		return
	}

	var k *ini.Key
	if k, err = p.section.GetKey(key); err != nil {
		return
	}
	value = k.String()
	return
}

//Get a key from loaded credentials as string
func (p *Profile) Duration(key string) (value time.Duration) {
	var err error
	if err = p.prepareProfile(); err != nil {
		return
	}

	var k *ini.Key
	if k, err = p.section.GetKey(key); err != nil {
		return
	}
	value, err = k.Duration()
	return
}

//Get a key from loaded credentials as string
func (p *Profile) Int(key string) (value int) {
	var err error
	if err = p.prepareProfile(); err != nil {
		return
	}

	var k *ini.Key
	if k, err = p.section.GetKey(key); err != nil {
		return
	}

	value, err = k.Int()
	return
}

func (p *Profile) Int64(key string) (value int64) {
	var err error
	if err = p.prepareProfile(); err != nil {
		return
	}

	var k *ini.Key
	if k, err = p.section.GetKey(key); err != nil {
		return
	}

	value, err = k.Int64()
	return
}

func (p *Profile) prepareProfile() (err error) {
	if p.section == nil {
		err = p.Load("", "")
	}
	return
}

//Determine the credential file path, based on $HOME/.bridgevine/credentials
func (p *Profile) getFilepath() (err error) {

	//You can override the default path using this environment variable
	if p.filepath = os.Getenv("CREDENTIALS_FILE"); p.filepath != "" {
		return
	}

	homeDir := os.Getenv("HOME") //*nix
	if homeDir == "" {           //windows
		homeDir = os.Getenv("USERPROFILE")
	}
	if homeDir == "" {
		err = fmt.Errorf("Credentials home not found")
		return
	}

	if p.path == "" {
		p.path = filepath.Join(homeDir, "/.profile/")
	}

	p.filepath = filepath.Join(p.path, "credentials")
	return
}
