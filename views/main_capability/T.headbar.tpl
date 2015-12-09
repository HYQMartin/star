{{define "headbar"}}

<div class="container" data-toggle="clingify">           
    <div  role="navigation" id="navbar-collapse">
        <ul class="nav navbar-nav">
            <li class="active"><a href="/">首页</a></li>
            <li><a href="/capability">能力地图</a></li>
            <li><a href="/training">培训交流</a></li>
            <li><a href="/newemployee/">新员工培养</a></li>
            
        </ul>
        
        <div class="pull-right">
			<ul class="nav navbar-nav">
				{{if $.IsLogin}}
					<li><a href="/logout">{{.UserName}},您好！退出登录</a></li>
				{{else}}
					<li><a href="/login">登录</a></li>
				{{end}}
			</ul>
		</div>
    </div>
       
</div>
{{end}}