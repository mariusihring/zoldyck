package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"os"
	"os/exec"
)

func runCommit() {
	var commitType string
	var commitMessage string
	var breakingChanges bool
	var breakingChangesDescription string
	var advancedDescription string
	var commitScope string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select the type of change that you're commiting:").
				Options(
					huh.NewOption("feat: - A new feature", "feat"),
					huh.NewOption("fix: - A bug fix", "fix"),
					huh.NewOption("docs: - Documentation only changes", "docs"),
					huh.NewOption("style: - Changes that do not affect the meaning of the code", "style"),
					huh.NewOption("refactor: - A code change that neither fixes a bug or adds a feature", "refactor"),
					huh.NewOption("perf: - A code change that improves performance", "perf"),
					huh.NewOption("test: - Adding missingn tests", "test"),
					huh.NewOption("chore: - Changes to the build process or auxiliary tools", "chore"),
				).
				Value(&commitType),
			huh.NewInput().Title("Scope").Description("The Scope of the commit").Value(&commitScope),

			huh.NewInput().Title("Commit Message").Value(&commitMessage),
			huh.NewConfirm().Title("Are there Breaking Changes?").Value(&breakingChanges),
		),
		huh.NewGroup(
			huh.NewInput().Title("Specify the Breaking Changes").Value(&breakingChangesDescription),
		).WithHideFunc(func() bool {
			return !breakingChanges
		}),
		huh.NewGroup(
			huh.NewText().Title("Body").Value(&advancedDescription),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	message := fmt.Sprintf("%s (%s): %s", commitType, commitScope, commitMessage)
	body := fmt.Sprintf("Breaking Changes: %s \n", map[bool]string{true: "Yes", false: "No"}[breakingChanges])
	if breakingChanges {
		body += fmt.Sprintf("What are the braking changes: %s \n", breakingChangesDescription)
	}
	if advancedDescription != "" {
		body += fmt.Sprintf("\n%s \n", advancedDescription)
	}
	cmd := exec.Command("git", "commit", "-m", message, "-m", body)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}
