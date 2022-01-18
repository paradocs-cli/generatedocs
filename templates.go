package generate_docs

//TerraDocGeneric is a template for generating terraform documentation
const TerraDocGeneric = `
<h1>Terraform Documentation</h1> 
<h2>Variables:</h2>
<table>
<thead>
<tr>
<th>Variable Name</th>
<th>Variable Type</th>
<th>Variable Default</th>
<th>Variable Description</th>
<th>Variable Required</th>
<th>Variable Sensitive</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Vars}}
<tbody>
<tr>
<td>{{.VarName}}</td>
<td>{{.VarType}}</td>
<td>{{.VarDescription}}</td>
<td>{{.VarDefault}}</td>
<td>{{.VarRequired}}</td>
<td>{{.VarSensitive}}</td>
<td>{{.SourcePositionFileName}}</td>
<td>{{.SourcePositionLine}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Resources:</h2>
<table>
<thead>
<tr>
<th>Mode</th>
<th>Type</th>
<th>Name</th>
<th>Provider Name</th>
<th>Provider Alias</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Resources}}
<tbody>
<tr>
<td>{{.Mode}}</td>
<td><a href="{{.Link}}">{{.Type}}</td>
<td>{{.Name}}</td>
<td>{{.ProviderName}}</td>
<td>{{.ProviderAlias}}</td>
<td>{{.SourcePositionFileName}}</td>
<td>{{.SourcePositionLine}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Modules:</h2>
<table>
<thead>
<tr>
<th>Name</th>
<th>Module Source</th>
<th>Version</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Modules}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.ModSource}}</td>
<td>{{.Version}}</td>
<td>{{.SourcePositionFileName}}</td>
<td>{{.SourcePositionLine}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Data Resources:</h2>
<table>
<thead>
<tr>
<th>Name</th>
<th>Data Type</th>
<th>Provider Name</th>
<th>Provider Alias</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Datas}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td><a href="{{.Link}}">{{.DataType}}</a></td>
<td>{{.ProviderName}}</td>
<td>{{.ProviderAlias}}</td>
<td>{{.SourcePositionFileName}}</td>
<td>{{.SourcePositionLine}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Providers:</h2>
<table>
<thead>
<tr>
<th>Name</th>
<th>Alias</th>
</tr>
</thead>
{{range .Providers}}
<tbody>
<tr>
<td><a href="{{.Link}}">{{.Name}}</a></td>
<td>{{.Alias}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Outputs:</h2>
<table>
<thead>
<tr>
<th>Name</th>
<th>Description</th>
<th>Sensitive</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Outputs}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.Description}}</td>
<td>{{.Sensitive}}</td>
<td>{{.SourcePositionFileName}}</td>
<td>{{.SourcePositionLine}}</td>
</tr>
</tbody>
{{end}}
</table>
`

//RepoDocGeneric is a template for generating repository documentation
const RepoDocGeneric = `
<h1>Repo Documentation</h1> 
<h2>Directories:</h2>
<table>
<thead>
<tr>
<th>Directory Name</th>
<th>Terraform Directory</th>
<th>Modification Time</th>
</tr>
</thead>
{{range .Directories}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.IsTerraDir}}</td>
<td>{{.ModificationTime}}</td>
</tr>
</tbody>
{{end}}
</table>

<h2>Files:</h2>
<table>
<thead>
<tr>
<th>Name</th>
<th>Terraform File</th>
<th>Modification Time</th>
</tr>
</thead>
{{range .Files}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.IsTfFile}}</td>
<td>{{.ModificationTime}}</td>
</tr>
</tbody>
{{end}}
</table>
`