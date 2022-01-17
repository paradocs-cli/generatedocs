package generate_docs


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
<td>{{.Type}}</td>
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
<th>Sensitive</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Datas}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.DataType}}</td>
<td>{{.ProviderName}}</td>
<td>{{.ProviderAlias}}</td>
<td>{{.Sensitive}}</td>
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
<td>{{.Name}}</td>
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

const RepoDocGeneric = `
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
<td>{{.Type}}</td>
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
<th>Sensitive</th>
<th>File Name</th>
<th>Line Number</th>
</tr>
</thead>
{{range .Datas}}
<tbody>
<tr>
<td>{{.Name}}</td>
<td>{{.DataType}}</td>
<td>{{.ProviderName}}</td>
<td>{{.ProviderAlias}}</td>
<td>{{.Sensitive}}</td>
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
<td>{{.Name}}</td>
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