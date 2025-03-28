package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Subproject struct to hold subproject data.
type Subproject struct {
	Name    string  `yaml:"name"`
	Contact Contact `yaml:"contact"`
}

// Person struct to hold person data, including company.
type Person struct {
	Name    string `yaml:"name"`
	GitHub  string `yaml:"github"`
	Company string `yaml:"company,omitempty"`
}

// Leadership struct to hold all leadership roles.
type Leadership struct {
	Chairs          []Person `yaml:"chairs"`
	TechLeads       []Person `yaml:"tech_leads"`
	SubprojectLeads []Person `yaml:"Subpproject Leads"`
}

// Meeting struct to hold meeting data.
type Meeting struct {
	Description   string `yaml:"description"`
	RecordingsURL string `yaml:"recordings_url"`
	TagCalendar   string `yaml:"tag_calendar,omitempty"`
}

// Contact struct to hold contact information.
type Contact struct {
	Slack       string `yaml:"slack"`
	MailingList string `yaml:"mailing_list"`
	TOCLiaison  Person `yaml:"toc_liaison"`
}

// Tag struct to hold tag data, including CharterLink and Subprojects.
type Tag struct {
	Dir              string       `yaml:"dir"`
	Name             string       `yaml:"name"`
	MissionStatement string       `yaml:"mission_statement"`
	Leadership       Leadership   `yaml:"leadership"`
	Meetings         []Meeting    `yaml:"meetings"`
	Contact          Contact      `yaml:"contact"`
	TagSubprojects   []Subproject `yaml:"tag_subprojects"`
	CharterLink      string       `yaml:"charter_link"`
	TagInitiatives   string       `yaml:"tag_initiatives"` // ADDED THIS LINE
}

// TOCSubproject struct to hold TOC subproject data, including CharterLink.
type TOCSubproject struct {
	Dir                   string     `yaml:"dir"`
	Name                  string     `yaml:"name"`
	MissionStatement      string     `yaml:"mission_statement"`
	Leadership            Leadership `yaml:"leadership"`
	Meetings              []Meeting  `yaml:"meetings"`
	Contact               Contact    `yaml:"contact"`
	CharterLink           string     `yaml:"charter_link"`
	SubprojectInitiatives string     `yaml:"subproject_initiatives"` // ADDED THIS LINE
}

// Config struct to hold the entire configuration.
type Config struct {
	Tags           []Tag           `yaml:"tags"`
	TOCSubprojects []TOCSubproject `yaml:"toc_subprojects"`
}

func main() {
	// Define paths.
	configPath := filepath.Join("..", "tags.yaml")
	tagsDir := filepath.Join("..", "tags")
	tocDir := filepath.Join("..", "toc_subprojects")
	tagTemplatePath := filepath.Join("..", "generator", "tag_readme.tmpl")
	tocTemplatePath := filepath.Join("..", "generator", "toc_subproject_readme.tmpl")

	// Read YAML file.
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read tags.yaml: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Load templates.
	tagTmplContent, err := os.ReadFile(tagTemplatePath)
	if err != nil {
		log.Fatalf("Failed to read tag template: %v", err)
	}
	tagTmpl, err := template.New("tag_readme").Parse(string(tagTmplContent))
	if err != nil {
		log.Fatalf("Failed to parse tag template: %v", err)
	}

	tocTmplContent, err := os.ReadFile(tocTemplatePath)
	if err != nil {
		log.Fatalf("Failed to read toc template: %v", err)
	}
	tocTmpl, err := template.New("toc_readme").Parse(string(tocTmplContent))
	if err != nil {
		log.Fatalf("Failed to parse toc template: %v", err)
	}

	// Ensure directories exist.
	if err := ensureDir(tagsDir); err != nil {
		log.Fatalf("Failed to create tags directory: %v", err)
	}
	if err := ensureDir(tocDir); err != nil {
		log.Fatalf("Failed to create TOC Subprojects directory: %v", err)
	}

	// Process each tag.
	for _, tag := range config.Tags {
		tagFolderPath := filepath.Join(tagsDir, tag.Dir)
		if err := ensureDir(tagFolderPath); err != nil {
			log.Fatalf("Failed to create folder for tag %s: %v", tag.Name, err)
		}

		var tagBuf bytes.Buffer
		if err := tagTmpl.Execute(&tagBuf, tag); err != nil {
			log.Fatalf("Tag template execution failed: %v", err)
		}

		tagFilePath := filepath.Join(tagFolderPath, "README.md")
		if err := os.WriteFile(tagFilePath, tagBuf.Bytes(), 0644); err != nil {
			log.Fatalf("Failed to write %s: %v", tagFilePath, err)
		}
		// create charter.md
		if tag.CharterLink != "" {
			charterPath := filepath.Join(tagFolderPath, "charter.md")
			if err := os.WriteFile(charterPath, []byte("Charter content here"), 0644); err != nil {
				log.Fatalf("Failed to write %s: %v", charterPath, err)
			}
		}
	}

	// Process each TOC subproject.
	for _, tocSubproject := range config.TOCSubprojects {
		tocSubprojectFolderPath := filepath.Join(tocDir, tocSubproject.Dir)
		if err := ensureDir(tocSubprojectFolderPath); err != nil {
			log.Fatalf("Failed to create folder for TOC subproject %s: %v", tocSubproject.Name, err)
		}

		var tocBuf bytes.Buffer
		if err := tocTmpl.Execute(&tocBuf, tocSubproject); err != nil {
			log.Fatalf("TOC template execution failed: %v", err)
		}

		tocReadmePath := filepath.Join(tocSubprojectFolderPath, "README.md")
		if err := os.WriteFile(tocReadmePath, tocBuf.Bytes(), 0644); err != nil {
			log.Fatalf("Failed to write %s: %v", tocReadmePath, err)
		}
		//Create charter.md
		if tocSubproject.CharterLink != "" {
			charterPath := filepath.Join(tocSubprojectFolderPath, "charter.md")
			if err := os.WriteFile(charterPath, []byte("Charter content here"), 0644); err != nil {
				log.Fatalf("Failed to write %s: %v", charterPath, err)
			}
		}
	}

	log.Println("Tag, Subproject, and README files have been generated.")
}

// ensureDir ensures the directory exists.
func ensureDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", dirPath, err)
	}
	return nil
}
