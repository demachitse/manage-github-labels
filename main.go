package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/demachitse/manage-github-labels/cmd"
	"github.com/demachitse/manage-github-labels/config"
	"github.com/spf13/viper"

	"github.com/google/go-github/github"
)

var client *github.Client
var ctx context.Context
var owner string
var repo string
var command string

func main() {
	var input string
	if len(os.Args) <= 1 {
		fmt.Print("owner: ")
		fmt.Scan(&input)
		owner = input
	} else {
		owner = os.Args[1]
		fmt.Println("owner: " + owner)
	}

	if len(os.Args) <= 2 {
		fmt.Print("repository: ")
		fmt.Scan(&input)
		repo = input
	} else {
		repo = os.Args[2]
		fmt.Println("repository: " + repo)
	}

	if len(os.Args) <= 3 {
		fmt.Println("command usages")
		fmt.Println("  r / reset: Delete all labels and create new labels from config.yaml.")
		fmt.Println("  g / get: Get all labels.")
		fmt.Println("  s / save: Save all labels to config.yaml.")
		fmt.Println("  c / create: Create new labels from config.yaml.")
		fmt.Println("  d / delete: Delete all labels.")
		fmt.Print("command: ")
		fmt.Scan(&input)
		command = input
	} else {
		command = os.Args[3]
		fmt.Println("command: " + command)
	}

	fmt.Println("Loadgin config...")
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}
	fmt.Println("  Loaded config")

	ctx = context.Background()
	client = cmd.GithubClient(ctx)

	var labels []*github.Label
	switch command {
	case "r", "reset":
		labels = getLabels()
		deleteLables(labels)
		createLabels(labels)

	case "s", "save":
		labels = getLabels()
		saveLabels(labels)

	case "g", "get":
		getLabels()

	case "c", "create":
		labels = getLabels()
		createLabels(labels)

	case "d", "delete":
		labels = getLabels()
		deleteLables(labels)

	default:
		log.Fatalf("Invalid command: %s", command)
	}

}

func getLabels() []*github.Label {
	// ラベルの取得
	fmt.Println("Getting labels...")
	labels, _, err := client.Issues.ListLabels(ctx, owner, repo, &github.ListOptions{Page: 1, PerPage: 5000})
	if err != nil {
		log.Fatalf("Failed to get labels: %s", err.Error())
	}

	if len(labels) > 0 {
		fmt.Println("  Labels")
		for _, label := range labels {
			fmt.Println("    name: " + *label.Name + "  color: " + *label.Color)
		}
	} else {
		fmt.Println("  There is no labels.")
	}

	return labels
}

func saveLabels(labels []*github.Label) {
	var input string
	if len(labels) > 0 {
		// 上書き確認
		fmt.Print("Saving labels to config.yaml, Are you sure? [Y/N]: ")
		fmt.Scan(&input)

		if input == "Y" || input == "y" {
			// ラベルの上書き
			config.Data.Labels = []github.Label{}
			for _, label := range labels {
				tmpLabel := github.Label{
					Name:  label.Name,
					Color: label.Color,
				}
				config.Data.Labels = append(config.Data.Labels, tmpLabel)
			}

			viper.Set("labels", config.Data.Labels)
			viper.WriteConfig()

			fmt.Println("  Saved labels.")
		} else {
			fmt.Println("  Cancelled saving labels.")
		}
	} else {
		fmt.Println("Saving labels to config.yaml.")
		fmt.Println("  There is no labels.")
	}
}

func createLabels(labels []*github.Label) {
	// ラベルの作成
	if len(config.Data.Labels) > 0 {
		fmt.Println("Creating new labels...")
		for _, newLabel := range config.Data.Labels {
			_, _, err := client.Issues.CreateLabel(ctx, owner, repo, &newLabel)
			if err != nil {
				if err.(*github.ErrorResponse).Errors[0].Code == "already_exists" {
					fmt.Println("  Label of " + *newLabel.Name + " already exists")
				} else {
					log.Fatalf("Failed to create label of %s: %s", *newLabel.Name, err.Error())
				}
			} else {
				fmt.Println("  Created " + *newLabel.Name)
			}
		}
	} else {
		fmt.Println("Creating new labels.")
		fmt.Println("  There is no new labels.")
	}
}

func deleteLables(labels []*github.Label) {
	var input string
	if len(labels) > 0 {
		// 削除確認
		fmt.Print("Deleting all labels, Are you sure? [Y/N]: ")
		fmt.Scan(&input)

		if input == "Y" || input == "y" {
			// ラベルの削除
			for _, label := range labels {
				_, err := client.Issues.DeleteLabel(ctx, owner, repo, *label.Name)
				if err != nil {
					log.Fatalf("Failed to delete label of %s: %s", *label.Name, err.Error())
				} else {
					fmt.Println("  Deleted " + *label.Name)
				}
			}
		} else {
			fmt.Println("  Cancelled deleting all labels.")
		}
	} else {
		fmt.Println("Deleting all labels.")
		fmt.Println("  There is no labels.")
	}
}
