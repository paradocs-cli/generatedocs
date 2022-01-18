package generate_docs

import (
	"fmt"
	"html/template"
	"os"
)

//WriteMarkdownTerra generates a README which recursively documents all terraform code
//WriteMarkdownTerra does template execution to write the file
func WriteMarkdownTerra(w Stats) os.File {
	tmpls := template.New("md")
	template.Must(tmpls.Parse(TerraDocGeneric))

	f, err := os.Create(fmt.Sprintf("TERRAFORM.%s", tmpls.Name()))
	if err != nil {
		if err != nil {
			fmt.Println("Unable to create file for Markdown")
			os.Exit(1)
		}
	}

	err = tmpls.Execute(f, w)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return *f
}

//WriteMarkdownRepo generates a README which recursively documents statistics about the repo from the root down
//WriteMarkdownRepo does template execution to write the file
func WriteMarkdownRepo(w RepoInfo) os.File {
	tmpls := template.New("md")
	template.Must(tmpls.Parse(RepoDocGeneric))

	f, err := os.Create(fmt.Sprintf("REPO.%s", tmpls.Name()))
	if err != nil {
		if err != nil {
			fmt.Println("Unable to create file for Markdown")
			os.Exit(1)
		}
	}

	err = tmpls.Execute(f, w)
	if err != nil {
		fmt.Println("Unable to execute template for WriteMarkdown")
		os.Exit(1)
	}

	return *f
}