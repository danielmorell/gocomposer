package gocomposer

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	doubleQuote  = 34
	leftBracket  = 91
	rightBracket = 93

	TypeComposer = "composer"
	TypeVCS      = "vcs"
	TypePath     = "path"
	TypeArtifact = "artifact"
	TypePear     = "pear"
	TypePackage  = "package"
)

// ComposerJSON is a complete representation of the composer.json file data.
type ComposerJSON struct {
	// ComposerJSON name, including 'vendor-name/' prefix.
	Name string `json:"name"`

	// Short package description.
	Description string `json:"description"`

	// License name. Or an array of license names.
	License StringOrSlice `json:"license"`

	// ComposerJSON type, either 'library' for common packages, 'composer-plugin' for
	// plugins, 'metapackage' for empty packages, or a custom type ([a-z0-9-]+) defined
	// by whatever project this package applies to.
	Type string `json:"type"`

	// Indicates whether this package has been abandoned, t can be boolean or a package
	// name/URL pointing to a recommended alternative. Defaults to false.
	Abandoned StringOrBool `json:"abandoned,omitempty"`

	// ComposerJSON version, see https://getcomposer.org/doc/04-schema.md#version for more
	// info on valid schemes.
	Version string `json:"version,omitempty"`

	// Internal use only, do not specify this in composer.json. Indicates whether this
	// version is the default branch of the linked VCS repository. Defaults to false.
	DefaultBranch bool `json:"default-branch,omitempty"`

	// A set of string or regex patterns for non-numeric branch names that will not be
	// handled as feature branches.
	NonFeatureBranches []string `json:"non-feature-branches,omitempty"`

	// A tag/keyword that this package relates to.
	Keywords []string `json:"keywords,omitempty"`

	// Relative path to the readme document.
	Readme string `json:"readme,omitempty"`

	// ComposerJSON release date, in 'YYYY-MM-DD', 'YYYY-MM-DD HH:MM:SS' or
	// 'YYYY-MM-DDTHH:MM:SSZ' format.
	Time time.Time `json:"time,omitempty"`

	// List of authors that contributed to the package. This is typically the main
	// maintainers, not the full list.
	Authors []Authors `json:"authors,omitempty"`

	// Homepage URL for the project.
	Homepage string `json:"homepage,omitempty"`

	// Support channels for the package
	Support Support `json:"support,omitempty"`

	// A list of options to fund the development and maintenance of the package.
	Funding []Funding `json:"funding"`

	// Source code details.
	Source Source `json:"source"`

	Dist Dist `json:"dist,omitempty"`

	// A key to store comments in.
	Comment StringOrSlice `json:"_comment,omitempty"`

	// This is an object of package name (keys) and version constraints (values) that
	// are required to run this package.
	Require map[string]string `json:"require,omitempty"`

	// This is an object of package name (keys) and version constraints (values) that
	// this package requires for developing it (testing tools and such).
	RequireDev map[string]string `json:"require-dev,omitempty"`

	// This is an object of package name (keys) and version constraints (values) that
	// can be replaced by this package.
	Replace map[string]string `json:"replace,omitempty"`

	// This is an object of package name (keys) and version constraints (values) that
	// conflict with this package.
	Conflict map[string]string `json:"conflict,omitempty"`

	// This is an object of package name (keys) and version constraints (values) that
	// this package provides in addition to this package's name.
	Provide map[string]string `json:"provide,omitempty"`

	// This is an object of package name (keys) and descriptions (values) that this
	// package suggests work well with it (this will be suggested to the user during
	// installation).
	Suggest map[string]string `json:"suggest,omitempty"`

	// A set of additional repositories where packages can be found.
	Repositories Repositories `json:"repositories,omitempty"`

	// The minimum stability the packages must have to be install-able. Possible values
	// are: dev, alpha, beta, RC, stable.
	MinimumStability string `json:"minimum-stability,omitempty"`

	// If set to true, stable packages will be preferred to dev packages when possible,
	// even if the minimum-stability allows unstable packages.
	PreferStable bool `json:"prefer-stable,omitempty"`

	// Description of how the package can be autoloaded.
	Autoload Autoload `json:"autoload,omitempty"`

	// Description of additional autoload rules for development purpose (eg. a test
	// suite).
	AutoloadDev AutoloadDev `json:"autoload-dev,omitempty"`

	// Forces the package to be installed into the given subdirectory path. This is
	// used for autoloading PSR-0 packages that do not contain their full path. Use
	// forward slashes for cross-platform compatibility.
	//
	// Deprecated: Not used with Composer 2.0
	TargetDir string `json:"target-dir,omitempty"`

	// A list of directories which should get added to PHP's include path. This is only
	// present to support legacy projects, and all new code should preferably use
	// autoloading.
	//
	// Deprecated: Not used with Composer 2.0
	IncludePath string `json:"include-path,omitempty"`

	// A set of files, or a single file, that should be treated as binaries and
	// symlinked into bin-dir (from config).
	Bin StringOrSlice `json:"bin,omitempty"`

	// Options for creating package archives for distribution.
	Archive Archive `json:"archive,omitempty"`

	// Composer options.
	Config Config `json:"config,omitempty"`

	Extra map[string]interface{} `json:"extra,omitempty"`
}

// List of authors that contributed to the package. This is typically the main
// maintainers, not the full list.
type Authors struct {
	// Full name of the author.
	Name string `json:"name"`
	// Email address of the author.
	Email string `json:"email,omitempty"`
	// Homepage URL for the author.
	Homepage string `json:"homepage,omitempty"`
	// Author's role in the project.
	Role string `json:"role,omitempty"`
}

// Description of how the package can be autoloaded.
type Autoload struct {
	// This is an array of files that are always required on every request.
	Files []string `json:"files,omitempty"`

	// This is an object of namespaces (keys) and the PSR-4 directories they can map to
	// (values, can be arrays of paths) by the autoloader.
	PSR4 map[string]StringOrSlice `json:"psr-4,omitempty"`

	// This is an object of namespaces (keys) and the directories they can be found in
	// (values, can be arrays of paths) by the autoloader.
	PSR0 map[string]StringOrSlice `json:"psr-0,omitempty"`

	// This is an array of paths that contain classes to be included in the class-map
	// generation process.
	ClassMap []string `json:"classmap,omitempty"`

	// This is an array of patterns to exclude from autoload classmap generation.
	// (e.g. "exclude-from-classmap": ["/test/", "/tests/", "/Tests/"]
	ExcludeFromClassMap []string `json:"exclude-from-classmap,omitempty"`
}

// Description of additional autoload rules for development purpose (eg. a test suite).
type AutoloadDev struct {
	// This is an array of files that are always required on every request.
	Files []string `json:"files,omitempty"`

	// This is an object of namespaces (keys) and the PSR-4 directories they can map to
	// (values, can be arrays of paths) by the autoloader.
	PSR4 map[string]StringOrSlice `json:"psr-4,omitempty"`

	// This is an object of namespaces (keys) and the directories they can be found in
	// (values, can be arrays of paths) by the autoloader.
	PSR0 map[string]StringOrSlice `json:"psr-0,omitempty"`

	// This is an array of paths that contain classes to be included in the class-map
	// generation process.
	ClassMap []string `json:"classmap,omitempty"`
}

// Options for creating package archives for distribution.
type Archive struct {
	// A base name for archive.
	Name string `json:"name,omitempty"`

	// A list of patterns for paths to exclude or include if prefixed with an
	// exclamation mark. Use like you would in a .gitignore file.
	Exclude []string `json:"exclude,omitempty"`
}

// Composer options.
type Config struct {
	// This is an object of package name (keys) and version (values) that will be used
	// to mock the platform packages on this machine, the version can be set to false to
	// make it appear like the package is not present.
	Platform map[string]StringOrBool `json:"platform,omitempty"`

	// This is an object of {"pattern": true|false} with packages which are allowed to
	// be loaded as plugins, or true to allow all, false to allow none. Defaults to {}
	// which prompts when an unknown plugin is added.
	AllowPlugins map[string]bool `json:"allow-plugins,omitempty"`

	// The timeout in seconds for process executions, defaults to 300 (5mins).
	ProcessTimeout int `json:"process-timeout,omitempty"`

	// If true, the Composer autoloader will also look for classes in the PHP include
	// path.
	UseIncludePath bool `json:"use-include-path,omitempty"`

	// When running Composer in a directory where there is no composer.json, if there is
	// one present in a directory above Composer will by default ask you whether you
	// want to use that directory's composer.json instead. One of: true (always use
	// parent if needed), false (never ask or use it) or "prompt" (ask every time),
	// defaults to prompt.
	UseParentDir StringOrBool `json:"use-parent-dir,omitempty"`

	// The install method Composer will prefer to use, defaults to auto and can be any
	// of source, dist, auto, or an object of {"pattern": "preference"}.
	PreferredInstall PreferredInstall `json:"preferred-install,omitempty"`
}

type Dist struct {
	URL       string `json:"url"`
	Type      string `json:"type"`
	Reference string `json:"reference,omitempty"`
	ShaSum    string `json:"shasum,omitempty"`
	Mirrors   string `json:"mirrors,omitempty"`
}

// Funding method to support the development and maintenance of the package.
type Funding struct {
	// Type of funding or platform through which funding is possible.
	Type string `json:"type"`
	// URL to a website with details on funding and a way to fund the package.
	URL string `json:"url"`
}

// A set of additional repositories where packages can be found.
type Repositories struct {
	object bool
	Array  []Repository
}

func (r Repositories) IsObject() bool {
	return r.object
}

// GetRepo looks for a Repository with a given name in the collection of Repositories.
func (r Repositories) GetRepo(name string) (Repository, bool) {
	for _, r := range r.Array {
		if r.Name == name {
			return r, true
		}
	}
	return Repository{}, false
}

func (r *Repositories) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if isObject(data) {
		temp := make(map[string]Repository)
		err := json.Unmarshal(data, &temp)
		if err != nil {
			return err
		}
		r.Array = make([]Repository, 0, 1)
		for name, rep := range temp {
			rep.Name = name
			r.Array = append(r.Array, rep)
		}
	}
	if isArray(data) {
		r.Array = make([]Repository, 0, 1)
		temp := make([]Repository, 0, 1)
		err := json.Unmarshal(data, &temp)
		if err != nil {
			return err
		}
		for _, rep := range temp {
			if rep.Disabled && rep.Name == "" {
				continue
			}
			r.Array = append(r.Array, rep)
		}
	}
	return nil
}

func (r Repositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Array)
}

type Repository struct {
	// Type of Repository used. This is derived from the actual repository type.
	// Valid options are [composer, vcs, path, artifact, pear, package].
	Type     string
	Name     string
	Disabled bool
	Composer ComposerRepository
	VCS      VCSRepository
	Path     PathRepository
	Artifact ArtifactRepository
	Pear     PearRepository
	Package  PackageRepository
}

func (r *Repository) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "false" {
		r.Disabled = true
		return nil
	}
	if !isObject(data) {
		return fmt.Errorf("invalid JSON repository")
	}

	// Check to see if this is a disabling repository e.g. {"packagist.org": false}.
	// Because "type" is required on all other repository types and prohibited on a
	// disabling repository we can check for the existence of "type" to determine if
	// this should be a disabling repo.
	disabled := make(map[string]interface{}, 0)
	err := json.Unmarshal(data, &disabled)
	if err != nil {
		return err
	}
	repoType, exists := disabled["type"]

	// If "type" does not exist, we know this is a disabling repository.
	if !exists {
		first := true
		for repo := range disabled {
			if !first {
				return fmt.Errorf("a disabling repo should only have one key, the name")
			}
			r.Disabled = true
			r.Name = repo
			first = false
		}
		return nil
	}

	t, ok := repoType.(string)
	if !ok {
		return fmt.Errorf("repository \"type\" must be a string")
	}

	switch t {
	case "composer":
		r.Type = TypeComposer
		return json.Unmarshal(data, &r.Composer)
	case "vcs", "github", "git", "gitlab", "bitbucket", "git-bitbucket", "hg", "fossil", "perforce", "svn":
		r.Type = TypeVCS
		return json.Unmarshal(data, &r.VCS)
	case "path":
		r.Type = TypePath
		return json.Unmarshal(data, &r.Path)
	case "artifact":
		r.Type = TypeArtifact
		return json.Unmarshal(data, &r.Artifact)
	case "pear":
		r.Type = TypePear
		return json.Unmarshal(data, &r.Pear)
	case "package":
		r.Type = TypePackage
		return json.Unmarshal(data, &r.Package)
	default:
		return fmt.Errorf(`repository type "%s" not recognized`, t)
	}
}

func (r Repository) MarshalJSON() ([]byte, error) {
	// Check if we should just return something like {"packagist.org": false}
	if r.Disabled {
		return json.Marshal(map[string]bool{r.Name: !r.Disabled})
	}
	switch r.Type {
	case TypeComposer:
		return json.Marshal(r.Composer)
	case TypeVCS:
		return json.Marshal(r.VCS)
	case TypePath:
		return json.Marshal(r.Path)
	case TypeArtifact:
		return json.Marshal(r.Artifact)
	case TypePear:
		return json.Marshal(r.Pear)
	case TypePackage:
		return json.Marshal(r.Package)
	default:
		return []byte{}, fmt.Errorf(`repository type "%s" not recognized`, r.Type)
	}
}

type ComposerRepository struct {
	Type               string                 `json:"type"`
	URL                string                 `json:"url"`
	Canonical          bool                   `json:"canonical,omitempty"`
	Only               []string               `json:"only,omitempty"`
	Exclude            []string               `json:"exclude,omitempty"`
	Options            map[string]interface{} `json:"extra,omitempty"`
	AllowSSLDowngrade  bool                   `json:"allow_ssl_downgrade,omitempty"`
	ForceLazyProviders bool                   `json:"force-lazy-providers,omitempty"`
}

type VCSRepository struct {
	Type                     string       `json:"type"`
	URL                      string       `json:"url"`
	Canonical                bool         `json:"canonical,omitempty"`
	Only                     []string     `json:"only,omitempty"`
	Exclude                  []string     `json:"exclude,omitempty"`
	NoAPI                    bool         `json:"no-api,omitempty"`
	SecureHTTP               bool         `json:"secure-http,omitempty"`
	SVNCacheCredentials      bool         `json:"svn-cache-credentials,omitempty"`
	TrunkPath                StringOrBool `json:"trunk-path,omitempty"`
	BranchesPath             StringOrBool `json:"branches-path,omitempty"`
	TagsPath                 StringOrBool `json:"tags-path,omitempty"`
	PackagePath              string       `json:"package-path,omitempty"`
	Depot                    string       `json:"depot,omitempty"`
	Branch                   string       `json:"branch,omitempty"`
	UniquePerforceClientName string       `json:"unique_perforce_client_name,omitempty"`
	P4User                   string       `json:"p4user,omitempty"`
	P4Password               string       `json:"p4password,omitempty"`
}

type PathRepository struct {
	Type      string                 `json:"type"`
	URL       string                 `json:"url"`
	Canonical bool                   `json:"canonical,omitempty"`
	Only      []string               `json:"only,omitempty"`
	Exclude   []string               `json:"exclude,omitempty"`
	Options   map[string]interface{} `json:"extra,omitempty"`
}

type ArtifactRepository struct {
	Type      string   `json:"type"`
	URL       string   `json:"url"`
	Canonical bool     `json:"canonical,omitempty"`
	Only      []string `json:"only,omitempty"`
	Exclude   []string `json:"exclude,omitempty"`
}

type PearRepository struct {
	Type        string   `json:"type"`
	URL         string   `json:"url"`
	Canonical   bool     `json:"canonical,omitempty"`
	Only        []string `json:"only,omitempty"`
	Exclude     []string `json:"exclude,omitempty"`
	VendorAlias string   `json:"vendor-alias,omitempty"`
}

type PackageRepository struct {
	Type      string         `json:"type"`
	Packages  PackageOrSlice `json:"package"`
	Canonical bool           `json:"canonical,omitempty"`
	Only      []string       `json:"only,omitempty"`
	Exclude   []string       `json:"exclude,omitempty"`
}

type PackageOrSlice []InlinePackage

func (p *PackageOrSlice) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if isObject(data) {
		temp := InlinePackage{}
		err := json.Unmarshal(data, &temp)
		if err != nil {
			return err
		}
		(*p) = append((*p), temp)
		return nil
	}
	if isArray(data) {
		err := json.Unmarshal(data, &p)
		return err
	}
	return nil
}

func (p PackageOrSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

// PreferredInstall for the project and its dependencies. Because order matters this is
// a slice of arrays with the pattern as the first item in the array and the install
// preferrence as the second item in the array.
type PreferredInstall [][2]string

func (p *PreferredInstall) UnmarshalJSON(data []byte) error {
	if 0 == len(data) {
		return nil
	}
	if isString(data) {
		value := ""
		err := json.Unmarshal(data, &value)
		if err != nil {
			return err
		}

		(*p) = append((*p), [2]string{"", value})
	}
	if isObject(data) {
		temp := make(map[string]string)
		err := json.Unmarshal(data, &temp)
		if err != nil {
			return err
		}
		for pattern, install := range temp {
			(*p) = append((*p), [2]string{pattern, install})
		}
	}
	return errors.New("preferred-install must be a string or an object")
}

func (p PreferredInstall) MarshalJSON() ([]byte, error) {
	if len(p) == 0 {
		return []byte{}, nil
	}
	if len(p) == 1 && p[0][0] == "" {
		value := ""
		for _, value = range p {
			break
		}
		return json.Marshal(value)
	}

	// Order matters so we serialize the JSON ourselves instead of transforming to a
	// map[string]string and marshaling, even though this would be simpler.
	buf := strings.Builder{}
	buf.WriteRune('{')
	for i, v := range p {
		pattern, err := json.Marshal([]byte(v[0]))
		if err != nil {
			return []byte{}, err
		}
		install, err := json.Marshal([]byte(v[1]))
		if err != nil {
			return []byte{}, err
		}

		buf.WriteString(fmt.Sprintf(`%s:%s`, pattern, install))
		if i < len(p) {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune('}')
	return []byte(buf.String()), nil
}

type InlinePackage struct {
	// TODO fill this out
}

type Source struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	Reference string `json:"reference"`
	Mirrors   string `json:"mirrors,omitempty"`
}

type StringOrSlice []string

func (s *StringOrSlice) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if isString(data) {
		(*s) = []string{string(data[1 : len(data)-1])}
		return nil
	}
	if isArray(data) {
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s StringOrSlice) MarshalJSON() ([]byte, error) {
	if len(s) == 0 {
		return []byte{}, nil
	}
	if len(s) == 1 {
		return json.Marshal(s[0])
	}
	return json.Marshal(s)
}

type StringOrBool struct {
	isBool      bool
	boolValue   bool
	stringValue string
}

func (s *StringOrBool) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	strData := string(data)
	if strData == "false" {
		s.boolValue = false
		s.isBool = true
		return nil
	}
	if strData == "true" {
		s.isBool = true
		s.boolValue = true
		return nil
	}
	if data[0] == doubleQuote && data[len(data)-1] == doubleQuote {
		s.isBool = false
		s.stringValue = string(data[1 : len(data)-1])
		return nil
	}
	return errors.New("invalid value, must be type string or bool")
}

func (s StringOrBool) MarshalJSON() ([]byte, error) {
	if s.isBool {
		return json.Marshal(s.boolValue)
	}
	return json.Marshal(s.stringValue)
}

// Support channels available for the package.
type Support struct {
	// Email address for support.
	Email string `json:"email"`
	// Issues tracker URL.
	Issues string `json:"issues"`
	// Forum URL.
	Forum string `json:"forum"`
	// Wiki URL.
	Wiki string `json:"wiki"`
	// IRC chanel for support, as irc://server/channel.
	IRC string `json:"irc"`
	// Chat support URL.
	Chat string `json:"chat"`
	// Source code browse or download URL.
	Source string `json:"source"`
	// Docs URL.
	Docs string `json:"docs"`
	// RSS feed URL.
	RSS string `json:"rss"`
}
