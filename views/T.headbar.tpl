{{define "headbar"}}
<html>
  	<head>
    	<title>THWL-wuziqi</title>
		<link rel="shortcut icon" href="/static/img/lyf.png" />
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

		 <!-- Stylesheets -->
		<link href="/static/css/bootstrap.min.css" rel="stylesheet" />
  	</head>
	
	<body>
		<div class="navbar navbar-default navbar-fixed-top">		
			<div class="container">
				
				<div>
					<ul class="nav navbar-nav">
						<li {{if .IsTHWL}}class="active"{{end}}><a class="navbar-brand" href="/THWL">Tentcoo</a></li>
						<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
						<li {{if .IsWuziqi}}class="active"{{end}} ><a href="/Wuziqi/PVC" >五子棋-人机对战</a></li>
						<li {{if .IsWuziqiPVP}}class="active"{{end}} ><a href="/Wuziqi/PVP" >五子棋-人人对战</a></li>
						<li {{if .IsResult}}class="active"{{end}} ><a href="/Result" >战绩</a></li>
					</ul>
				</div>
				
			</div>

		</div>
		<div>
		<div class="container">
				<p align="right">作者：{{.Author}}  </p>
				<h5 align="right">时间：{{.Time}}    </h5>
		
		</div>
		</div>
{{end}}
