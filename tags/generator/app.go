package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	yaml "gopkg.in/yaml.v3"
)

const (
	readmeTemplate   = "readme.tmpl"
	listTemplate     = "list.tmpl"
	liaisonsTemplate = "liaisons.tmpl"
	headerTemplate   = "header.tmpl"

	tagsYamlFile     = "tags.yaml"
	tagListOutput    = "tag-list.md"
	indexFilename    = "README.md"
	liaisonsFilename = "liaisons.md"

	beginCustomMarkdown = "<!-- BEGIN CUSTOM CONTENT -->"
	endCustomMarkdown   = "<!-- END CUSTOM CONTENT -->"
	beginCustomYaml     = "## BEGIN CUSTOM CONTENT"
	endCustomYaml       = "## END CUSTOM CONTENT"
)

var (
	baseGeneratorDir = ""
	templateDir      = "generator"
)

// FoldedString is a string that will be serialized in FoldedStyle by go-yaml
type FoldedString string

// MarshalYAML customizes how FoldedStrings will be serialized by go-yaml
func (x FoldedString) MarshalYAML() (interface{}, error) {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FoldedStyle,
		Value: string(x),
	}, nil
}

// Person represents an individual person holding a role in a group.
type Person struct {
	GitHub  string
	Name    string
	Company string `yaml:"company,omitempty"`
}

// Meeting represents a regular meeting for a group.
type Meeting struct {
	Description   string
	Day           string
	Time          string
	TZ            string
	Frequency     string
	URL           string `yaml:",omitempty"`
	ArchiveURL    string `yaml:"archive_url,omitempty"`
	RecordingsURL string `yaml:"recordings_url,omitempty"`
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack              string       `yaml:",omitempty"`
	MailingList        string       `yaml:"mailing_list,omitempty"`
	PrivateMailingList string       `yaml:"private_mailing_list,omitempty"`
	GithubTeams        []GithubTeam `yaml:"teams,omitempty"`
	Liaisons           []Person     `yaml:"liaisons,omitempty"`
}

// GithubTeam represents a specific Github Team.
type GithubTeam struct {
	Name        string
	Description string `yaml:",omitempty"`
}

// WorkingGroup represents a specific WorkingGroup owned by the group
type WorkingGroup struct {
	Name        string
	Description string    `yaml:",omitempty"`
	CharterLink string    `yaml:"charter_link,omitempty"`
	Contact     *Contact  `yaml:",omitempty"`
	Leads       []Person  `yaml:"leads,omitempty"`
	Meetings    []Meeting `yaml:",omitempty"`
}

// LeadershipGroup represents the different groups of leaders within a group
type LeadershipGroup struct {
	Chairs         []Person
	TechnicalLeads []Person `yaml:"tech_leads,omitempty"`
	EmeritusLeads  []Person `yaml:"emeritus_leads,omitempty"`
}

// PrefixToPersonMap returns a map of prefix to persons, useful for iteration over all persons
func (g *LeadershipGroup) PrefixToPersonMap() map[string][]Person {
	return map[string][]Person{
		"chair":         g.Chairs,
		"tech_lead":     g.TechnicalLeads,
		"emeritus_lead": g.EmeritusLeads,
	}
}

// Leads returns a sorted and de-duped list of Leads for a LeadershipGroup
func (g *LeadershipGroup) Leads() []Person {
	o := append(g.Chairs, g.TechnicalLeads...)

	// Sort
	sort.Slice(o, func(i, j int) bool {
		return o[i].GitHub < o[j].GitHub
	})

	// De-dupe
	seen := make(map[string]struct{}, len(o))
	i := 0
	for _, p := range o {
		if _, ok := seen[p.GitHub]; ok {
			continue
		}
		seen[p.GitHub] = struct{}{}
		o[i] = p
		i++
	}
	return o[:i]
}

// Group represents either a TAG
type Group struct {
	Dir           string
	Prefix        string `yaml:",omitempty"`
	Name          string
	Objective     string          `yaml:"objective,omitempty"`
	CharterLink   string          `yaml:"charter_link,omitempty"`
	Leadership    LeadershipGroup `yaml:"leadership"`
	Meetings      []Meeting
	Contact       Contact
	WorkingGroups []WorkingGroup `yaml:"working_groups,omitempty"`
}

type TAGName string

func (n TAGName) DirName() string {
	return DirName("tag", string(n))
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (tag for TAGs),
// and a formatted version of the group's name (in kebab case).
func (g *Group) DirName(prefix string) string {
	return DirName(prefix, g.Name)
}

func DirName(prefix, name string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(name, " ", "-", -1)))
}

// Context is the context for the tags.yaml file.
type Context struct {
	Tags []Group
}

func index(groups []Group, predicate func(Group) bool) int {
	for i, group := range groups {
		if predicate(group) {
			return i
		}
	}
	return -1
}

// PrefixToGroupMap returns a map of prefix to groups, useful for iteration over all groups
func (c *Context) PrefixToGroupMap() map[string][]Group {
	return map[string][]Group{
		"tag": c.Tags,
	}
}

// Complete populates derived portions of the Context struct
func (c *Context) Complete() {
}

// Sort sorts all lists within the Context struct
func (c *Context) Sort() {
	for _, groups := range c.PrefixToGroupMap() {
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].Dir < groups[j].Dir
		})
		for _, group := range groups {

			for _, people := range [][]Person{
				group.Leadership.Chairs,
				group.Leadership.TechnicalLeads,
				group.Leadership.EmeritusLeads} {
				sort.Slice(people, func(i, j int) bool {
					// This ensures that leads are ordered by github
					return people[i].GitHub < people[j].GitHub
				})
			}
			sort.Slice(group.Meetings, func(i, j int) bool {
				return group.Meetings[i].Description < group.Meetings[j].Description
			})
			sort.Slice(group.Contact.GithubTeams, func(i, j int) bool {
				return group.Contact.GithubTeams[i].Name < group.Contact.GithubTeams[j].Name
			})
			sort.Slice(group.WorkingGroups, func(i, j int) bool {
				return group.WorkingGroups[i].Name < group.WorkingGroups[j].Name
			})
			for _, WorkingGroup := range group.WorkingGroups {
				for _, people := range [][]Person{WorkingGroup.Leads} {
					sort.Slice(people, func(i, j int) bool {
						return people[i].GitHub < people[j].GitHub
					})
				}

				if WorkingGroup.Contact != nil {
					sort.Slice(WorkingGroup.Contact.GithubTeams, func(i, j int) bool {
						return WorkingGroup.Contact.GithubTeams[i].Name < WorkingGroup.Contact.GithubTeams[j].Name
					})
				}
				sort.Slice(WorkingGroup.Meetings, func(i, j int) bool {
					return WorkingGroup.Meetings[i].Description < WorkingGroup.Meetings[j].Description
				})
			}
		}
	}
}

// Validate returns a list of errors encountered while validating a Context
func (c *Context) Validate() []error {
	errors := []error{}
	people := make(map[string]Person)

	for prefix, groups := range c.PrefixToGroupMap() {
		for _, group := range groups {
			expectedDir := group.DirName(prefix)
			if expectedDir != group.Dir {
				errors = append(errors, fmt.Errorf("expected dir: %s, got: %s", expectedDir, group.Dir))
			}

			for prefix, persons := range group.Leadership.PrefixToPersonMap() {
				for _, person := range persons {
					if val, ok := people[person.GitHub]; ok {
						if val.Name != person.Name || (prefix != "emeritus_lead" && val.Company != person.Company) {
							errors = append(errors, fmt.Errorf("%s: %ss: expected person: %v, got: %v", group.Dir, prefix, val, person))
						}
					} else if prefix != "emeritus_lead" {
						people[person.GitHub] = person
					}

					if prefix == "emeritus_lead" && person.Company != "" {
						errors = append(errors, fmt.Errorf("%s: emeritus leads should not have company specified; company specified for: %s", group.Dir, person.Name))
					}
				}
			}

			if prefix == "tag" {
				if group.CharterLink == "" {
					errors = append(errors, fmt.Errorf("%s: has no charter", group.Dir))
				}
				if group.Objective == "" {
					errors = append(errors, fmt.Errorf("%s: has no objective", group.Dir))
				}
			}
		}
	}
	return errors
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createDirIfNotExists(path string) error {
	if !pathExists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func getExistingContent(path string, fileFormat string) (string, error) {
	capture := false
	var captured []string

	beginMarker := ""
	endMarker := ""
	switch fileFormat {
	case "markdown":
		beginMarker = beginCustomMarkdown
		endMarker = endCustomMarkdown
	case "yaml":
		beginMarker = beginCustomYaml
		endMarker = endCustomYaml
	case "":
		return "", nil
	}

	// NOTE: For some reason using bufio.Scanner with existing file pointer prepends
	// a bunch of null ^@ characters, so using to ioutil.ReadFile instead.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if strings.Contains(line, endMarker) {
			capture = false
		}
		if capture {
			captured = append(captured, line)
		}
		if strings.Contains(line, beginMarker) {
			capture = true
		}
	}

	return strings.Join(captured, "\n"), nil
}

var funcMap = template.FuncMap{
	"tzUrlEncode": tzURLEncode,
	"trimSpace":   strings.TrimSpace,
	"trimSuffix":  strings.TrimSuffix,
	"now":         time.Now,
	"lastYear":    lastYear,
	"toUpper":     strings.ToUpper,
}

// lastYear returns the last year as a string
func lastYear() string {
	return time.Now().AddDate(-1, 0, 0).Format("2006")
}

// tzUrlEncode returns a url encoded string without the + shortcut. This is
// required as the timezone conversion site we are using doesn't recognize + as
// a valid url escape character.
func tzURLEncode(tz string) string {
	return strings.Replace(url.QueryEscape(tz), "+", "%20", -1)
}

func writeTemplate(templatePath, outputPath string, fileFormat string, data interface{}) error {
	// set up template
	t, err := template.New(filepath.Base(templatePath)).
		Funcs(funcMap).
		ParseFiles(templatePath, filepath.Join(baseGeneratorDir, templateDir, headerTemplate))
	if err != nil {
		return err
	}

	// create if not exists
	if !pathExists(outputPath) {
		_, err = os.Create(outputPath)
		if err != nil {
			return err
		}
	}

	// open file and truncate
	f, err := os.OpenFile(outputPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// get any existing content
	content, err := getExistingContent(outputPath, fileFormat)
	if err != nil {
		return err
	}

	// ensure file is empty
	f.Truncate(0)

	// generated content
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	writeCustomContentBlock(f, content, fileFormat)

	return nil
}

func writeCustomContentBlock(f *os.File, content string, fileFormat string) {
	beginMarker := ""
	endMarker := ""
	switch fileFormat {
	case "markdown":
		beginMarker = beginCustomMarkdown
		endMarker = endCustomMarkdown
	case "yaml":
		beginMarker = beginCustomYaml
		endMarker = endCustomYaml
	case "":
		return
	}

	lines := []string{beginMarker, "\n", content, "\n", endMarker, "\n"}
	for _, line := range lines {
		f.Write([]byte(line))
	}
}

func createGroupReadme(groups []Group, prefix string) error {
	// figure out if the user wants to generate one group
	var selectedGroupName *string
	if envVal, ok := os.LookupEnv("WHAT"); ok {
		selectedGroupName = &envVal
	}

	for _, group := range groups {
		// skip generation if the user specified only one group
		if selectedGroupName != nil && strings.HasSuffix(group.Dir, *selectedGroupName) == false {
			fmt.Printf("Skipping %s/README.md\n", group.Dir)
			continue
		}

		fmt.Printf("Generating %s/README.md\n", group.Dir)

		outputDir := filepath.Join(baseGeneratorDir, group.Dir)
		if err := createDirIfNotExists(outputDir); err != nil {
			return err
		}

		outputPath := filepath.Join(outputDir, indexFilename)
		readmePath := filepath.Join(baseGeneratorDir, templateDir, fmt.Sprintf("%s_%s", prefix, readmeTemplate))
		if err := writeTemplate(readmePath, outputPath, "markdown", group); err != nil {
			return err
		}
	}

	return nil
}

// readYaml decodes yaml stored in a file at path into the
// specified yaml.Node
func readYaml(path string, data interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	return decoder.Decode(data)
}

// writeYaml writes the specified data to a file at path
// indent is set to 2 spaces
func writeYaml(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := yaml.NewEncoder(file)
	enc.SetIndent(2)
	return enc.Encode(data)
}

func main() {
	yamlPath := filepath.Join(baseGeneratorDir, tagsYamlFile)
	var ctx Context

	err := readYaml(yamlPath, &ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Complete()

	ctx.Sort()

	fmt.Printf("Validating %s\n", yamlPath)
	errs := ctx.Validate()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		}
		os.Exit(1)
	}

	// Write the Context struct back to yaml to enforce formatting
	err = writeYaml(&ctx, yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating group READMEs")
	for prefix, groups := range ctx.PrefixToGroupMap() {
		err = createGroupReadme(groups, prefix)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Generating tag-list.md")
	outputPath := filepath.Join(baseGeneratorDir, tagListOutput)
	err = writeTemplate(filepath.Join(baseGeneratorDir, templateDir, listTemplate), outputPath, "markdown", ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating liaisons.md")
	outputPath = filepath.Join(baseGeneratorDir, liaisonsFilename)
	err = writeTemplate(filepath.Join(baseGeneratorDir, templateDir, liaisonsTemplate), outputPath, "markdown", ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished generation!")
}
