package terraform

import (
	"errors"
)

// Destroy invokes Terraform's "destroy" command.
func (terraformer *Terraformer) Destroy() (success bool, err error) {
	success, err = terraformer.RunStreamed("destroy",
		"-force",
		"-input=false",
		"-auto-approve",
		"-no-color",
		"-var-file=tfvars.json",
	)
	if err != nil {
		return
	}
	if !success {
		err = errors.New("Failed to execute 'terraform destroy'")
	}

	return
}
