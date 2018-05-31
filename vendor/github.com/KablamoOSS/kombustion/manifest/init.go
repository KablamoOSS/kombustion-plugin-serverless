package manifest

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"

	"gopkg.in/AlecAivazis/survey.v1"
)

// InitManifestFlags - Flags that will prevent prompts
// TODO: implment flags
var InitManifestFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "name, n",
		Usage: "Set the name of the project",
	},
	cli.StringFlag{
		Name:  "environments",
		Usage: "comma seperated environments eg: production,development",
	},
}

// InitaliseNewManifestTask - Create a new manifest file, and prompt to fill out
// the default required fields
func InitaliseNewManifestTask(c *cli.Context) error {

	// This funciton is a thin layer between the task, and the cli wrapper
	err := initaliseNewManifest()

	return err
}

// initaliseNewManifest
func initaliseNewManifest() error {
	// TODO: Check if there is a manifest file and exit

	// Load the manifest file from this directory
	_, err := FindAndLoadManifest()
	if err == nil {
		log.Fatal("Sorry we can't create a new kombustion.yaml, one already exists.")
	}

	// Survey the user for required info
	name, environments, err := surveyForInitialManifest()
	if err != nil {
		return err
	}

	manifest := Manifest{
		Name:         name,
		Environments: environments,
	}
	fmt.Print(manifest)

	err = WriteManifestToDisk(manifest)
	if err != nil {
		return err
	}
	return nil
}

// surveyForInitialManifest - Prompt the user to fill out the required fields,
// but check if we have a flag for them
func surveyForInitialManifest() (
	name string,
	environments map[string]ManifestEnvironment,
	err error) {

	// name
	surveyName, err := surveyForName()
	if err != nil {
		return name, environments, err
	}
	name = surveyName

	//environments
	surveyEnvironments, err := surveyForEnvironments()
	if err != nil {
		return name, environments, err
	}
	environments = surveyEnvironments

	return name, environments, nil
}

// surveyForName - Prompt for the name of the project
func surveyForName() (string, error) {
	// the questions to ask
	var surveyQuestions = []*survey.Question{
		{
			Name:     "Name",
			Prompt:   &survey.Input{Message: "What is the name of this project?"},
			Validate: survey.Required,
		},
	}

	// the answers will be written to this struct
	surveyAnswers := struct {
		Name string // survey will match the question and field names
	}{}

	// perform the questions
	err := survey.Ask(surveyQuestions, &surveyAnswers)
	if err != nil {
		return "", err
	}

	// TODO: Add a better transform here, to ensure name is valid to the
	// Cloudformation name spec
	return strings.Replace(surveyAnswers.Name, " ", "", -1), nil
}

// surveyForEnvironments - prompt the user to find out what environments this
// project uses
func surveyForEnvironments() (manifestEnvironments map[string]ManifestEnvironment, err error) {
	// Survey for which environments are used in this project
	environments := []string{}
	manifestEnvironments = map[string]ManifestEnvironment{}

	prompt := &survey.MultiSelect{
		Message: "Which environments does this project deploy to:",
		Help:    "you can add more later",
		Options: []string{"production", "pre-production", "staging", "development"},
	}
	// Prompts the user
	err = survey.AskOne(prompt, &environments, nil)
	if err != nil {
		return manifestEnvironments, err
	}

	// TODO: prompt to ask for whitelisting account ids

	for _, env := range environments {
		manifestEnvironments[env] = ManifestEnvironment{
			AccountIDs: []string{""},
			Parameters: map[string]string{"": ""},
		}
	}

	return manifestEnvironments, err
}
