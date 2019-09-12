package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cecotw/strut-cli/internal/app/product"
	"github.com/cecotw/strut-cli/internal/pkg/file"
	"github.com/fatih/color"
)

func createPrompt(name string) (*product.Product, *file.Type) {
	answers := struct {
		*product.Product
		Extension string
	}{}
	prompt := []*survey.Question{}
	if name == "" {
		prompt = append(prompt, &survey.Question{
			Name:   "name",
			Prompt: &survey.Input{Message: "Enter new strut product name:"},
		})
	} else {
		answers.Name = name
	}

	prompt = append(prompt, []*survey.Question{
		{
			Name:   "description",
			Prompt: &survey.Input{Message: "Enter new product description:"},
		},
		{
			Name: "extension",
			Prompt: &survey.Select{
				Message: "Select file type:",
				Options: []string{
					file.Types.YAML.Extension,
					file.Types.JSON.Extension,
				},
				Default: file.Types.YAML.Extension,
			},
		},
	}...)

	err := survey.Ask(prompt, answers)
	if err != nil {
		fmt.Println(err.Error())
	}
	ft := file.Types.YAML
	for _, fileType := range file.TypeList {
		if fileType.Extension == answers.Extension {
			ft = fileType
			break
		}
	}
	return &product.Product{
		Name:        answers.Name,
		Description: answers.Description,
	}, ft
}

func addApplicationPrompt() *product.Application {
	color.Yellow("Lets add an application to your product.")
	answers := &product.Application{
		Version:     "0.0.0",
		LocalConfig: &product.LocalConfig{},
	}
	err := survey.Ask([]*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "Enter application name:"},
		},
	}, answers)
	if err != nil {
		fmt.Println(err.Error())
	}
	include := struct{ hasRepo bool }{}
	err = survey.Ask([]*survey.Question{
		{
			Name:   "hasRepo",
			Prompt: &survey.Confirm{Message: "Include Repo?"},
		},
	}, include)
	if include.hasRepo {
		err = survey.Ask([]*survey.Question{
			{
				Name:   "url",
				Prompt: &survey.Input{Message: "Provide the remote URL to the app code:"},
			},
			{
				Name: "type",
				Prompt: &survey.Select{
					Message: "Please select your VCS:",
					Options: []string{"git", "SVN", "mercurial"},
				},
			},
		}, answers.Repository)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	err = survey.Ask([]*survey.Question{
		{
			Name:   "path",
			Prompt: &survey.Input{Message: "Provide the local path to the application code:"},
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Please select your VCS:",
				Options: []string{"git", "SVN", "mercurial"},
			},
		},
	}, answers.LocalConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	return answers
}
