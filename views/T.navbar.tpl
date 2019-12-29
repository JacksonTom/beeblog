{{define "navbar"}}
<a class="navbar-brand" href="/">Jackson Tom</a>
    <div>
        <ul class="nav navbar-nav">
            <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
			<li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
			<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
        </ul>
    </div>

    <div class="pull-right">
        <ul class="nav navbar-nav">
            <li>
                <a href="http://www.beian.miit.gov.cn" target="_Blank">
                    苏ICP备19012905号
                </a>
            </li>
            <li>
                <a href="https://github.com/JacksonTom/beeblog" target="_Blank">
				　　<img  src="/static/img/Github.png" >
			    </a>
            </li>
            {{if .IsLogin}}
                <li><a href="/exit">退出</a></li>
            {{else}}
                <li><a href="/login">登录</a></li>
            {{end}}
        </ul>
    </div>
{{end}}