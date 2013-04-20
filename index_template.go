package main

import "html/template"

const (
	BLOG_INDEX = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
  <head>
    <title>Blog Entries</title>
    <script src='/assets/bundle.js' type='text/javascript'></script>
    <script src='/assets/application.js' type='text/javascript'></script>
    <link href='/assets/bundle.css' rel='stylesheet' type='text/css' />
    <link href='/assets/application.css' rel='stylesheet' type='text/css' />
  </head>
  <body>
    <nav class='top-bar'>
      <ul class='title-area'>
        <li class='name'>
          <h1>
            <a href='/'>Andrew C. Sellers</a>
          </h1>
        </li>
        <li class='toggle-topbar menu-icon'>
          <a href='#'>
            <span></span>
          </a>
        </li>
      </ul>
      <section class='top-bar-section'>
        <ul class='left'>
          <li class='divider'></li>
          <li class="active">
            <a href='/'>Blog</a>
          </li>
        </ul>
        <ul class='right'>
          <li>
            <a href='/about'>About</a>
          </li>
        </ul>
      </section>
    </nav>
    <div class='row'>
      <div class='large-12 columns'>
        {{if .Title}}
          <h1 class='subheader'>{{.Title}}</h1>
        {{else}}
          <h1 class='subheader'>Blog Entries</h1>
        {{end}}
        {{range .Entries}}
          <div class="entry">
            <h3><a href="/page/{{.ID}}">{{.Title}}</a></h3>
            <p>{{.Hook}}</p>
          </div>
        {{end}}
      </div>
    </div>
  </body>
</html>`
)

var indexTemplate = template.Must(template.New("index").Parse(BLOG_INDEX))
