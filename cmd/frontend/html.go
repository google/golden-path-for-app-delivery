package main

type InstanceMetadata struct {
	ID         string
	Name       string
	Version    string
	Hostname   string
	Zone       string
	Project    string
	InternalIP string
	ExternalIP string
	LBRequest  string
	ClientIP   string
	Error      string
}

const (
	html = `<!doctype html>
<html>
<head>
<!-- Compiled and minified CSS -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/css/materialize.min.css">

<!-- Compiled and minified JavaScript -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/js/materialize.min.js"></script>
<title>Frontend Web Server</title>
</head>
<body>
<div class="container">
<div class="row">
<div class="col s2">&nbsp;</div>
<div class="col s8">


<div class="card blue"><!-- color -->
<div class="card-content white-text">
<div class="card-title">Backend that serviced this request</div>
</div>
<div class="card-content white">
<table class="bordered">
  <tbody>
	<tr>
	  <td>Name</td>
	  <td>{{.Name}}</td>
	</tr>
	<tr>
	  <td>Version</td>
	  <td>{{.Version}}</td>
	</tr>
	<tr>
	  <td>Instance ID</td>
	  <td>{{.ID}}</td>
	</tr>
	<tr>
	  <td>Hostname</td>
	  <td>{{.Hostname}}</td>
	</tr>
	<tr>
	  <td>Zone</td>
	  <td>{{.Zone}}</td>
	</tr>
	<tr>
	  <td>Project</td>
	  <td>{{.Project}}</td>
	</tr>
	<tr>
	  <td>Internal IP</td>
	  <td>{{.InternalIP}}</td>
	</tr>
	<tr>
	  <td>External IP</td>
	  <td>{{.ExternalIP}}</td>
	</tr>
  </tbody>
</table>
</div>
</div>

<div class="card blue">
<div class="card-content white-text">
<div class="card-title">Proxy that handled this request</div>
</div>
<div class="card-content white">
<table class="bordered">
  <tbody>
	<tr>
	  <td>Request Headers</td>
	  <td><pre>{{.LBRequest}}</pre></td>
	</tr>
	<tr>
		<td>Error</td>
		<td>{{.Error}}</td>
	</tr>
</tbody>
</table>
</div>

</div>
</div>
<div class="col s2">&nbsp;</div>
</div>
</div>
</html>`
)
