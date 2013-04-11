package main

import "html/template"

const (
	BLOG_ENTRY = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
  <head>
    <title>{{ .Entry.Title }}</title>
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
          <li>
            <a href='/'>Blog</a>
          </li>
          <li class="active">
            <a href='/page/{{.Entry.ID}}'>{{.Entry.Title}}</a>
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
      <div class='large-9 columns'>
        <h1 class='subheader'>{{.Entry.Title}}</h1>
        <p>{{.Entry.Content}}</p>
      </div>
      <div class='large-3 columns'>
        <div class='panel'>
          <ul class='circle'>
            <li>
              Recent Articles
              <ul>
                {{range .NewEntries}}
                  <li>
                    <a href='/page/{{.ID}}'>{{.Title}}</a>
                  </li>
                {{end}}
              </ul>
            </li>
            <li>
              Popular Articles
              <ul>
                {{range .TopEntries}}
                  <li>
                    <a href='/page/{{.ID}}'>{{.Title}}</a>
                  </li>
                {{end}}
              </ul>
            </li>
            <li>
              Categories
              <ul>
                {{range .Categories}}
                  <li>
                    <a href='/{{.ID}}'>{{.Name}}</a>
                  </li>
                {{end}}
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </body>
</html>`
)

var entryTemplate = template.Must(template.New("entry").Parse(BLOG_ENTRY))
