{{define "navbar"}}
<div><br> </br></div>
<div class="navbar navbar-default navbar-fixed-top">		
	<div class="container">
	
	<div>
		<ul class="nav navbar-nav">
			<li {{if .IsTHWL}}class="active"{{end}}><a class="navbar-brand">Huawei</a></li>
			<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
			<li {{if .IsWuziqi}}class="active"{{end}} ><a href="/map" >能力地图</a></li>
			<li {{if .IsWuziqiPVP}}class="active"{{end}} ><a href="/communicate">新员工交流</a></li>
			<li {{if .IsResult}}class="active"{{end}} ><a href="/study">新员工培养</a></li>
		</ul>
	</div>
	
	<div class="pull-right">
		<ul class="nav navbar-nav">
			{{if .IsLogin}}
				<li><a href="/Wuziqi/PVP/login?exit=true">{{.UserName}},您好！退出登录</a></li>
			{{else}}
				<li><a href="/Wuziqi/PVP/login">登录</a></li>
			{{end}}
		</ul>
	</div>
	
	</div>

</div>


<div><br> </br></div>
{{end}}