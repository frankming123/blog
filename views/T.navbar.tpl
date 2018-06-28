{{define "navbar"}}
<header>
  <nav class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      <a class="navbar-brand" href="/">我的博客</a>
      <ul class="nav navbar-nav">
        <li {{if .IsHome}}class="active" {{end}}>
          <a href="/">首页</a>
        </li>
        <li {{if .IsCategory}}class="active" {{end}}>
          <a href="/category">分类</a>
        </li>
        <li {{if .IsTopic}}class="active" {{end}}>
          <a href="/topic">文章</a>
        </li>
      </ul>
      <ul class="nav navbar-nav pull-right">
        <li {{if .IsLoginPage}}class="active" {{end}}>
          {{if .IsLogin}}
          <a href="/login?exit=true">退出</a>
          {{else}}
          <a href="/login">登录</a>
          {{end}}
        </li>
      </ul>
    </div>
  </nav>
</header>
{{end}}