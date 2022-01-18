package generate_docs

import "time"


type Stats struct {
	Vars      []Var      `json:"variables,omitempty"`
	Resources []Resource `json:"resources,omitempty"`
	Modules   []Module   `json:"modules,omitempty"`
	Outputs   []Output   `json:"outputs,omitempty"`
	Datas     []Data     `json:"datas,omitempty"`
	Providers []Provider `json:"providers,omitempty"`
}
type Var struct {
	VarName                string      `json:"name,omitempty"`
	VarType                string      `json:"type,omitempty"`
	VarDescription         string      `json:"description,omitempty"`
	VarDefault             interface{} `json:"default_value,omitempty"`
	VarRequired            bool        `json:"required,omitempty"`
	VarSensitive           bool        `json:"sensitive,omitempty"`
	SourcePositionFileName string      `json:"source_file_name,omitempty"`
	SourcePositionLine     string      `json:"source_position_line,omitempty"`
}
type Resource struct {
	Mode                   string `json:"mode,omitempty"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	ProviderName           string `json:"provider_name,omitempty"`
	ProviderAlias          string `json:"provider_alias,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Module struct {
	Name                   string `json:"name,omitempty"`
	ModSource              string `json:"module_source,omitempty"`
	Version                string `json:"module_version,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
}

type Output struct {
	Name                   string `json:"name,omitempty"`
	Description            string `json:"description,omitempty"`
	Sensitive              bool   `json:"module_version,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Data struct {
	Name                   string `json:"name,omitempty"`
	DataType               string `json:"data_type,omitempty"`
	ProviderName           string `json:"provider_name,omitempty"`
	ProviderAlias          string `json:"provider_alias,omitempty"`
	Senstive               bool   `json:"module_version,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Provider struct {
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type Dirs struct {
	Name string
	ModificationTime time.Time
	IsTerraDir bool
}

type File struct {
	Name string
	ModificationTime time.Time
	IsTfFile bool
}

type RepoInfo struct {
	Directories []Dirs
	Files []File
}