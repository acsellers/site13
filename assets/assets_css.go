package assets

const (
	BUNDLE_CSS = `*,
*:before,
*:after {
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box; }

html,
body {
  font-size: 100%; }

body {
  background: white;
  color: #222222;
  padding: 0;
  margin: 0;
  font-family: "Helvetica Neue", "Helvetica", Helvetica, Arial, sans-serif;
  font-weight: normal;
  font-style: normal;
  line-height: 1;
  position: relative; }

a:focus {
  outline: none; }

img,
object,
embed {
  max-width: 100%;
  height: auto; }

object,
embed {
  height: 100%; }

img {
  -ms-interpolation-mode: bicubic; }

#map_canvas img,
#map_canvas embed,
#map_canvas object,
.map_canvas img,
.map_canvas embed,
.map_canvas object {
  max-width: none !important; }

.left {
  float: left !important; }

.right {
  float: right !important; }

.text-left {
  text-align: left !important; }

.text-right {
  text-align: right !important; }

.text-center {
  text-align: center !important; }

.text-justify {
  text-align: justify !important; }

.hide {
  display: none; }

.antialiased {
  -webkit-font-smoothing: antialiased; }

img {
  display: inline-block; }

textarea {
  height: auto;
  min-height: 50px; }

select {
  width: 100%; }

/* Grid HTML Classes */
.row {
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  margin-top: 0;
  margin-bottom: 0;
  max-width: 62.5em;
  *zoom: 1; }
  .row:before, .row:after {
    content: " ";
    display: table; }
  .row:after {
    clear: both; }
  .row .column,
  .row .columns {
    position: relative;
    padding-left: 0.9375em;
    padding-right: 0.9375em;
    width: 100%;
    float: left; }
  .row.collapse .column,
  .row.collapse .columns {
    position: relative;
    padding-left: 0;
    padding-right: 0;
    float: left; }
  .row .row {
    width: auto;
    margin-left: -0.9375em;
    margin-right: -0.9375em;
    margin-top: 0;
    margin-bottom: 0;
    max-width: none;
    *zoom: 1; }
    .row .row:before, .row .row:after {
      content: " ";
      display: table; }
    .row .row:after {
      clear: both; }
    .row .row.collapse {
      width: auto;
      margin: 0;
      max-width: none;
      *zoom: 1; }
      .row .row.collapse:before, .row .row.collapse:after {
        content: " ";
        display: table; }
      .row .row.collapse:after {
        clear: both; }

@media only screen {
  .row .column,
  .row .columns {
    position: relative;
    padding-left: 0.9375em;
    padding-right: 0.9375em;
    float: left; }

  .row .small-1 {
    position: relative;
    width: 8.33333%; }

  .row .small-2 {
    position: relative;
    width: 16.66667%; }

  .row .small-3 {
    position: relative;
    width: 25%; }

  .row .small-4 {
    position: relative;
    width: 33.33333%; }

  .row .small-5 {
    position: relative;
    width: 41.66667%; }

  .row .small-6 {
    position: relative;
    width: 50%; }

  .row .small-7 {
    position: relative;
    width: 58.33333%; }

  .row .small-8 {
    position: relative;
    width: 66.66667%; }

  .row .small-9 {
    position: relative;
    width: 75%; }

  .row .small-10 {
    position: relative;
    width: 83.33333%; }

  .row .small-11 {
    position: relative;
    width: 91.66667%; }

  .row .small-12 {
    position: relative;
    width: 100%; }

  .row .small-offset-1 {
    position: relative;
    margin-left: 8.33333%; }

  .row .small-offset-2 {
    position: relative;
    margin-left: 16.66667%; }

  .row .small-offset-3 {
    position: relative;
    margin-left: 25%; }

  .row .small-offset-4 {
    position: relative;
    margin-left: 33.33333%; }

  .row .small-offset-5 {
    position: relative;
    margin-left: 41.66667%; }

  .row .small-offset-6 {
    position: relative;
    margin-left: 50%; }

  .row .small-offset-7 {
    position: relative;
    margin-left: 58.33333%; }

  .row .small-offset-8 {
    position: relative;
    margin-left: 66.66667%; }

  .row .small-offset-9 {
    position: relative;
    margin-left: 75%; }

  .row .small-offset-10 {
    position: relative;
    margin-left: 83.33333%; }

  [class*="column"] + [class*="column"]:last-child {
    float: right; }

  [class*="column"] + [class*="column"].end {
    float: left; }

  .column.small-centered,
  .columns.small-centered {
    position: relative;
    margin-left: auto;
    margin-right: auto;
    float: none !important; } }
/* Styles for screens that are atleast 768px; */
@media only screen and (min-width: 48em) {
  .row .large-1 {
    position: relative;
    width: 8.33333%; }

  .row .large-2 {
    position: relative;
    width: 16.66667%; }

  .row .large-3 {
    position: relative;
    width: 25%; }

  .row .large-4 {
    position: relative;
    width: 33.33333%; }

  .row .large-5 {
    position: relative;
    width: 41.66667%; }

  .row .large-6 {
    position: relative;
    width: 50%; }

  .row .large-7 {
    position: relative;
    width: 58.33333%; }

  .row .large-8 {
    position: relative;
    width: 66.66667%; }

  .row .large-9 {
    position: relative;
    width: 75%; }

  .row .large-10 {
    position: relative;
    width: 83.33333%; }

  .row .large-11 {
    position: relative;
    width: 91.66667%; }

  .row .large-12 {
    position: relative;
    width: 100%; }

  .row .large-offset-1 {
    position: relative;
    margin-left: 8.33333%; }

  .row .large-offset-2 {
    position: relative;
    margin-left: 16.66667%; }

  .row .large-offset-3 {
    position: relative;
    margin-left: 25%; }

  .row .large-offset-4 {
    position: relative;
    margin-left: 33.33333%; }

  .row .large-offset-5 {
    position: relative;
    margin-left: 41.66667%; }

  .row .large-offset-6 {
    position: relative;
    margin-left: 50%; }

  .row .large-offset-7 {
    position: relative;
    margin-left: 58.33333%; }

  .row .large-offset-8 {
    position: relative;
    margin-left: 66.66667%; }

  .row .large-offset-9 {
    position: relative;
    margin-left: 75%; }

  .row .large-offset-10 {
    position: relative;
    margin-left: 83.33333%; }

  .push-2 {
    position: relative;
    left: 16.66667%;
    right: auto; }

  .pull-2 {
    position: relative;
    right: 16.66667%;
    left: auto; }

  .push-3 {
    position: relative;
    left: 25%;
    right: auto; }

  .pull-3 {
    position: relative;
    right: 25%;
    left: auto; }

  .push-4 {
    position: relative;
    left: 33.33333%;
    right: auto; }

  .pull-4 {
    position: relative;
    right: 33.33333%;
    left: auto; }

  .push-5 {
    position: relative;
    left: 41.66667%;
    right: auto; }

  .pull-5 {
    position: relative;
    right: 41.66667%;
    left: auto; }

  .push-6 {
    position: relative;
    left: 50%;
    right: auto; }

  .pull-6 {
    position: relative;
    right: 50%;
    left: auto; }

  .push-7 {
    position: relative;
    left: 58.33333%;
    right: auto; }

  .pull-7 {
    position: relative;
    right: 58.33333%;
    left: auto; }

  .push-8 {
    position: relative;
    left: 66.66667%;
    right: auto; }

  .pull-8 {
    position: relative;
    right: 66.66667%;
    left: auto; }

  .push-9 {
    position: relative;
    left: 75%;
    right: auto; }

  .pull-9 {
    position: relative;
    right: 75%;
    left: auto; }

  .push-10 {
    position: relative;
    left: 83.33333%;
    right: auto; }

  .pull-10 {
    position: relative;
    right: 83.33333%;
    left: auto; }

  .small-push-2 {
    left: inherit; }

  .small-pull-2 {
    right: inherit; }

  .small-push-3 {
    left: inherit; }

  .small-pull-3 {
    right: inherit; }

  .small-push-4 {
    left: inherit; }

  .small-pull-4 {
    right: inherit; }

  .small-push-5 {
    left: inherit; }

  .small-pull-5 {
    right: inherit; }

  .small-push-6 {
    left: inherit; }

  .small-pull-6 {
    right: inherit; }

  .small-push-7 {
    left: inherit; }

  .small-pull-7 {
    right: inherit; }

  .small-push-8 {
    left: inherit; }

  .small-pull-8 {
    right: inherit; }

  .small-push-9 {
    left: inherit; }

  .small-pull-9 {
    right: inherit; }

  .small-push-10 {
    left: inherit; }

  .small-pull-10 {
    right: inherit; }

  .column.large-centered,
  .columns.large-centered {
    position: relative;
    margin-left: auto;
    margin-right: auto;
    float: none !important; } }
/* Foundation Visibility HTML Classes */
.show-for-small,
.show-for-medium-down,
.show-for-large-down {
  display: inherit !important; }

.show-for-medium,
.show-for-medium-up,
.show-for-large,
.show-for-large-up,
.show-for-xlarge {
  display: none !important; }

.hide-for-medium,
.hide-for-medium-up,
.hide-for-large,
.hide-for-large-up,
.hide-for-xlarge {
  display: inherit !important; }

.hide-for-small,
.hide-for-medium-down,
.hide-for-large-down {
  display: none !important; }

/* Specific visilbity for tables */
table.show-for-small, table.show-for-medium-down, table.show-for-large-down, table.hide-for-medium, table.hide-for-medium-up, table.hide-for-large, table.hide-for-large-up, table.hide-for-xlarge {
  display: table; }

thead.show-for-small, thead.show-for-medium-down, thead.show-for-large-down, thead.hide-for-medium, thead.hide-for-medium-up, thead.hide-for-large, thead.hide-for-large-up, thead.hide-for-xlarge {
  display: table-header-group !important; }

tbody.show-for-small, tbody.show-for-medium-down, tbody.show-for-large-down, tbody.hide-for-medium, tbody.hide-for-medium-up, tbody.hide-for-large, tbody.hide-for-large-up, tbody.hide-for-xlarge {
  display: table-row-group !important; }

tr.show-for-small, tr.show-for-medium-down, tr.show-for-large-down, tr.hide-for-medium, tr.hide-for-medium-up, tr.hide-for-large, tr.hide-for-large-up, tr.hide-for-xlarge {
  display: table-row !important; }

td.show-for-small, td.show-for-medium-down, td.show-for-large-down, td.hide-for-medium, td.hide-for-medium-up, td.hide-for-large, td.hide-for-large-up, td.hide-for-xlarge,
th.show-for-small,
th.show-for-medium-down,
th.show-for-large-down,
th.hide-for-medium,
th.hide-for-medium-up,
th.hide-for-large,
th.hide-for-large-up,
th.hide-for-xlarge {
  display: table-cell !important; }

/* Medium Displays: 768px - 1279px */
@media only screen and (min-width: 48em) {
  .show-for-medium,
  .show-for-medium-up {
    display: inherit !important; }

  .show-for-small {
    display: none !important; }

  .hide-for-small {
    display: inherit !important; }

  .hide-for-medium,
  .hide-for-medium-up {
    display: none !important; }

  /* Specific visilbity for tables */
  table.show-for-medium, table.show-for-medium-up, table.hide-for-small {
    display: table; }

  thead.show-for-medium, thead.show-for-medium-up, thead.hide-for-small {
    display: table-header-group !important; }

  tbody.show-for-medium, tbody.show-for-medium-up, tbody.hide-for-small {
    display: table-row-group !important; }

  tr.show-for-medium, tr.show-for-medium-up, tr.hide-for-small {
    display: table-row !important; }

  td.show-for-medium, td.show-for-medium-up, td.hide-for-small,
  th.show-for-medium,
  th.show-for-medium-up,
  th.hide-for-small {
    display: table-cell !important; } }
/* Large Displays: 1280px - 1440px */
@media only screen and (min-width: 80em) {
  .show-for-large,
  .show-for-large-up {
    display: inherit !important; }

  .show-for-medium,
  .show-for-medium-down {
    display: none !important; }

  .hide-for-medium,
  .hide-for-medium-down {
    display: inherit !important; }

  .hide-for-large,
  .hide-for-large-up {
    display: none !important; }

  /* Specific visilbity for tables */
  table.show-for-large, table.show-for-large-up, table.hide-for-medium, table.hide-for-medium-down {
    display: table; }

  thead.show-for-large, thead.show-for-large-up, thead.hide-for-medium, thead.hide-for-medium-down {
    display: table-header-group !important; }

  tbody.show-for-large, tbody.show-for-large-up, tbody.hide-for-medium, tbody.hide-for-medium-down {
    display: table-row-group !important; }

  tr.show-for-large, tr.show-for-large-up, tr.hide-for-medium, tr.hide-for-medium-down {
    display: table-row !important; }

  td.show-for-large, td.show-for-large-up, td.hide-for-medium, td.hide-for-medium-down,
  th.show-for-large,
  th.show-for-large-up,
  th.hide-for-medium,
  th.hide-for-medium-down {
    display: table-cell !important; } }
/* X-Large Displays: 1400px and up */
@media only screen and (min-width: 90em) {
  .show-for-xlarge {
    display: inherit !important; }

  .show-for-large,
  .show-for-large-down {
    display: none !important; }

  .hide-for-large,
  .hide-for-large-down {
    display: inherit !important; }

  .hide-for-xlarge {
    display: none !important; }

  /* Specific visilbity for tables */
  table.show-for-xlarge, table.hide-for-large, table.hide-for-large-down {
    display: table; }

  thead.show-for-xlarge, thead.hide-for-large, thead.hide-for-large-down {
    display: table-header-group !important; }

  tbody.show-for-xlarge, tbody.hide-for-large, tbody.hide-for-large-down {
    display: table-row-group !important; }

  tr.show-for-xlarge, tr.hide-for-large, tr.hide-for-large-down {
    display: table-row !important; }

  td.show-for-xlarge, td.hide-for-large, td.hide-for-large-down,
  th.show-for-xlarge,
  th.hide-for-large,
  th.hide-for-large-down {
    display: table-cell !important; } }
/* Orientation targeting */
.show-for-landscape,
.hide-for-portrait {
  display: inherit !important; }

.hide-for-landscape,
.show-for-portrait {
  display: none !important; }

/* Specific visilbity for tables */
table.hide-for-landscape, table.show-for-portrait {
  display: table; }

thead.hide-for-landscape, thead.show-for-portrait {
  display: table-header-group !important; }

tbody.hide-for-landscape, tbody.show-for-portrait {
  display: table-row-group !important; }

tr.hide-for-landscape, tr.show-for-portrait {
  display: table-row !important; }

td.hide-for-landscape, td.show-for-portrait,
th.hide-for-landscape,
th.show-for-portrait {
  display: table-cell !important; }

@media only screen and (orientation: landscape) {
  .show-for-landscape,
  .hide-for-portrait {
    display: inherit !important; }

  .hide-for-landscape,
  .show-for-portrait {
    display: none !important; }

  /* Specific visilbity for tables */
  table.show-for-landscape, table.hide-for-portrait {
    display: table; }

  thead.show-for-landscape, thead.hide-for-portrait {
    display: table-header-group !important; }

  tbody.show-for-landscape, tbody.hide-for-portrait {
    display: table-row-group !important; }

  tr.show-for-landscape, tr.hide-for-portrait {
    display: table-row !important; }

  td.show-for-landscape, td.hide-for-portrait,
  th.show-for-landscape,
  th.hide-for-portrait {
    display: table-cell !important; } }
@media only screen and (orientation: portrait) {
  .show-for-portrait,
  .hide-for-landscape {
    display: inherit !important; }

  .hide-for-portrait,
  .show-for-landscape {
    display: none !important; }

  /* Specific visilbity for tables */
  table.show-for-portrait, table.hide-for-landscape {
    display: table; }

  thead.show-for-portrait, thead.hide-for-landscape {
    display: table-header-group !important; }

  tbody.show-for-portrait, tbody.hide-for-landscape {
    display: table-row-group !important; }

  tr.show-for-portrait, tr.hide-for-landscape {
    display: table-row !important; }

  td.show-for-portrait, td.hide-for-landscape,
  th.show-for-portrait,
  th.hide-for-landscape {
    display: table-cell !important; } }
/* Touch-enabled device targeting */
.show-for-touch {
  display: none !important; }

.hide-for-touch {
  display: inherit !important; }

.touch .show-for-touch {
  display: inherit !important; }

.touch .hide-for-touch {
  display: none !important; }

/* Specific visilbity for tables */
table.hide-for-touch {
  display: table; }

.touch table.show-for-touch {
  display: table; }

thead.hide-for-touch {
  display: table-header-group !important; }

.touch thead.show-for-touch {
  display: table-header-group !important; }

tbody.hide-for-touch {
  display: table-row-group !important; }

.touch tbody.show-for-touch {
  display: table-row-group !important; }

tr.hide-for-touch {
  display: table-row !important; }

.touch tr.show-for-touch {
  display: table-row !important; }

td.hide-for-touch {
  display: table-cell !important; }

.touch td.show-for-touch {
  display: table-cell !important; }

th.hide-for-touch {
  display: table-cell !important; }

.touch th.show-for-touch {
  display: table-cell !important; }

/* Foundation Block Grids for below small breakpoint */
@media only screen {
  [class*="block-grid-"] {
    display: block;
    padding: 0;
    margin: 0 -10px;
    *zoom: 1; }
    [class*="block-grid-"]:before, [class*="block-grid-"]:after {
      content: " ";
      display: table; }
    [class*="block-grid-"]:after {
      clear: both; }
    [class*="block-grid-"] > li {
      display: block;
      height: auto;
      float: left;
      padding: 0 10px 10px; }

  .small-block-grid-1 > li {
    width: 100%;
    padding: 0 10px 10px; }
    .small-block-grid-1 > li:nth-of-type(1n+1) {
      clear: both; }

  .small-block-grid-2 > li {
    width: 50%;
    padding: 0 10px 10px; }
    .small-block-grid-2 > li:nth-of-type(2n+1) {
      clear: both; }

  .small-block-grid-3 > li {
    width: 33.33333%;
    padding: 0 10px 10px; }
    .small-block-grid-3 > li:nth-of-type(3n+1) {
      clear: both; }

  .small-block-grid-4 > li {
    width: 25%;
    padding: 0 10px 10px; }
    .small-block-grid-4 > li:nth-of-type(4n+1) {
      clear: both; }

  .small-block-grid-5 > li {
    width: 20%;
    padding: 0 10px 10px; }
    .small-block-grid-5 > li:nth-of-type(5n+1) {
      clear: both; }

  .small-block-grid-6 > li {
    width: 16.66667%;
    padding: 0 10px 10px; }
    .small-block-grid-6 > li:nth-of-type(6n+1) {
      clear: both; }

  .small-block-grid-7 > li {
    width: 14.28571%;
    padding: 0 10px 10px; }
    .small-block-grid-7 > li:nth-of-type(7n+1) {
      clear: both; }

  .small-block-grid-8 > li {
    width: 12.5%;
    padding: 0 10px 10px; }
    .small-block-grid-8 > li:nth-of-type(8n+1) {
      clear: both; }

  .small-block-grid-9 > li {
    width: 11.11111%;
    padding: 0 10px 10px; }
    .small-block-grid-9 > li:nth-of-type(9n+1) {
      clear: both; }

  .small-block-grid-10 > li {
    width: 10%;
    padding: 0 10px 10px; }
    .small-block-grid-10 > li:nth-of-type(10n+1) {
      clear: both; }

  .small-block-grid-11 > li {
    width: 9.09091%;
    padding: 0 10px 10px; }
    .small-block-grid-11 > li:nth-of-type(11n+1) {
      clear: both; }

  .small-block-grid-12 > li {
    width: 8.33333%;
    padding: 0 10px 10px; }
    .small-block-grid-12 > li:nth-of-type(12n+1) {
      clear: both; } }
/* Foundation Block Grids for above small breakpoint */
@media only screen and (min-width: 48em) {
  .large-block-grid-1 > li {
    width: 100%;
    padding: 0 10px 10px; }
    .large-block-grid-1 > li:nth-of-type(1n+1) {
      clear: both; }

  .large-block-grid-2 > li {
    width: 50%;
    padding: 0 10px 10px; }
    .large-block-grid-2 > li:nth-of-type(2n+1) {
      clear: both; }

  .large-block-grid-3 > li {
    width: 33.33333%;
    padding: 0 10px 10px; }
    .large-block-grid-3 > li:nth-of-type(3n+1) {
      clear: both; }

  .large-block-grid-4 > li {
    width: 25%;
    padding: 0 10px 10px; }
    .large-block-grid-4 > li:nth-of-type(4n+1) {
      clear: both; }

  .large-block-grid-5 > li {
    width: 20%;
    padding: 0 10px 10px; }
    .large-block-grid-5 > li:nth-of-type(5n+1) {
      clear: both; }

  .large-block-grid-6 > li {
    width: 16.66667%;
    padding: 0 10px 10px; }
    .large-block-grid-6 > li:nth-of-type(6n+1) {
      clear: both; }

  .large-block-grid-7 > li {
    width: 14.28571%;
    padding: 0 10px 10px; }
    .large-block-grid-7 > li:nth-of-type(7n+1) {
      clear: both; }

  .large-block-grid-8 > li {
    width: 12.5%;
    padding: 0 10px 10px; }
    .large-block-grid-8 > li:nth-of-type(8n+1) {
      clear: both; }

  .large-block-grid-9 > li {
    width: 11.11111%;
    padding: 0 10px 10px; }
    .large-block-grid-9 > li:nth-of-type(9n+1) {
      clear: both; }

  .large-block-grid-10 > li {
    width: 10%;
    padding: 0 10px 10px; }
    .large-block-grid-10 > li:nth-of-type(10n+1) {
      clear: both; }

  .large-block-grid-11 > li {
    width: 9.09091%;
    padding: 0 10px 10px; }
    .large-block-grid-11 > li:nth-of-type(11n+1) {
      clear: both; }

  .large-block-grid-12 > li {
    width: 8.33333%;
    padding: 0 10px 10px; }
    .large-block-grid-12 > li:nth-of-type(12n+1) {
      clear: both; }

  [class*="small-block-grid-"] > li {
    clear: none !important; } }
p.lead {
  font-size: 1.21875em;
  line-height: 1.6; }

.subheader {
  line-height: 1.4;
  color: #6f6f6f;
  font-weight: 300;
  margin-top: 0.2em;
  margin-bottom: 0.5em; }

/* Typography resets */
div,
dl,
dt,
dd,
ul,
ol,
li,
h1,
h2,
h3,
h4,
h5,
h6,
pre,
form,
p,
blockquote,
th,
td {
  margin: 0;
  padding: 0;
  direction: ltr; }

/* Default Link Styles */
a {
  color: #2ba6cb;
  text-decoration: none;
  line-height: inherit; }
  a:hover, a:focus {
    color: #2795b6; }
  a img {
    border: none; }

/* Default paragraph styles */
p {
  font-family: inherit;
  font-weight: normal;
  font-size: 1em;
  line-height: 1.6;
  margin-bottom: 1.25em;
  text-rendering: optimizeLegibility; }
  p aside {
    font-size: 0.875em;
    line-height: 1.35;
    font-style: italic; }

/* Default header styles */
h1, h2, h3, h4, h5, h6 {
  font-family: "Helvetica Neue", "Helvetica", Helvetica, Arial, sans-serif;
  font-weight: bold;
  font-style: normal;
  color: #222222;
  text-rendering: optimizeLegibility;
  margin-top: 0.2em;
  margin-bottom: 0.5em;
  line-height: 1.2125em; }
  h1 small, h2 small, h3 small, h4 small, h5 small, h6 small {
    font-size: 60%;
    color: #6f6f6f;
    line-height: 0; }

h1 {
  font-size: 2.125em; }

h2 {
  font-size: 1.6875em; }

h3 {
  font-size: 1.375em; }

h4 {
  font-size: 1.125em; }

h5 {
  font-size: 1.125em; }

h6 {
  font-size: 1em; }

hr {
  border: solid #dddddd;
  border-width: 1px 0 0;
  clear: both;
  margin: 1.25em 0 1.1875em;
  height: 0; }

/* Helpful Typography Defaults */
em,
i {
  font-style: italic;
  line-height: inherit; }

strong,
b {
  font-weight: bold;
  line-height: inherit; }

small {
  font-size: 60%;
  line-height: inherit; }

code {
  font-family: Consolas, "Liberation Mono", Courier, monospace;
  font-weight: bold;
  color: #7f0a0c; }

/* Lists */
ul,
ol,
dl {
  font-size: 1em;
  line-height: 1.6;
  margin-bottom: 1.25em;
  list-style-position: outside;
  font-family: inherit; }

/* Unordered Lists */
ul li ul,
ul li ol {
  margin-left: 1.25em;
  margin-bottom: 0;
  font-size: 1em;
  /* Override nested font-size change */ }
ul.square li ul, ul.circle li ul, ul.disc li ul {
  list-style: inherit; }
ul.square {
  list-style-type: square; }
ul.circle {
  list-style-type: circle; }
ul.disc {
  list-style-type: disc; }
ul.no-bullet {
  list-style: none; }

/* Ordered Lists */
ol li ul,
ol li ol {
  margin-left: 1.25em;
  margin-bottom: 0; }

/* Definition Lists */
dl dt {
  margin-bottom: 0.3em;
  font-weight: bold; }
dl dd {
  margin-bottom: 0.75em; }

/* Abbreviations */
abbr,
acronym {
  text-transform: uppercase;
  font-size: 90%;
  color: #222222;
  border-bottom: 1px dotted #dddddd;
  cursor: help; }

abbr {
  text-transform: none; }

/* Blockquotes */
blockquote {
  margin: 0 0 1.25em;
  padding: 0.5625em 1.25em 0 1.1875em;
  border-left: 1px solid #dddddd; }
  blockquote cite {
    display: block;
    font-size: 0.8125em;
    color: #555555; }
    blockquote cite:before {
      content: "\2014 \0020"; }
    blockquote cite a,
    blockquote cite a:visited {
      color: #555555; }

blockquote,
blockquote p {
  line-height: 1.6;
  color: #6f6f6f; }

/* Microformats */
.vcard {
  display: inline-block;
  margin: 0 0 1.25em 0;
  border: 1px solid #dddddd;
  padding: 0.625em 0.75em; }
  .vcard li {
    margin: 0;
    display: block; }
  .vcard .fn {
    font-weight: bold;
    font-size: 0.9375em; }

.vevent .summary {
  font-weight: bold; }
.vevent abbr {
  cursor: default;
  text-decoration: none;
  font-weight: bold;
  border: none;
  padding: 0 0.0625em; }

@media only screen and (min-width: 48em) {
  h1, h2, h3, h4, h5, h6 {
    line-height: 1.4; }

  h1 {
    font-size: 2.75em; }

  h2 {
    font-size: 2.3125em; }

  h3 {
    font-size: 1.6875em; }

  h4 {
    font-size: 1.4375em; } }
/*
 * Print styles.
 *
 * Inlined to avoid required HTTP connection: www.phpied.com/delay-loading-your-print-css/
 * Credit to Paul Irish and HTML5 Boilerplate (html5boilerplate.com)
*/
.print-only {
  display: none !important; }

@media print {
  * {
    background: transparent !important;
    color: black !important;
    /* Black prints faster: h5bp.com/s */
    box-shadow: none !important;
    text-shadow: none !important; }

  a,
  a:visited {
    text-decoration: underline; }

  a[href]:after {
    content: " (" attr(href) ")"; }

  abbr[title]:after {
    content: " (" attr(title) ")"; }

  .ir a:after,
  a[href^="javascript:"]:after,
  a[href^="#"]:after {
    content: ""; }

  pre,
  blockquote {
    border: 1px solid #999999;
    page-break-inside: avoid; }

  thead {
    display: table-header-group;
    /* h5bp.com/t */ }

  tr,
  img {
    page-break-inside: avoid; }

  img {
    max-width: 100% !important; }

  @page {
    margin: 0.5cm; }

  p,
  h2,
  h3 {
    orphans: 3;
    widows: 3; }

  h2,
  h3 {
    page-break-after: avoid; }

  .hide-on-print {
    display: none !important; }

  .print-only {
    display: block !important; }

  .hide-for-print {
    display: none !important; }

  .show-for-print {
    display: inherit !important; } }
button, .button {
  border-style: solid;
  border-width: 1px;
  cursor: pointer;
  font-family: inherit;
  font-weight: bold;
  line-height: 1;
  margin: 0 0 1.25em;
  position: relative;
  text-decoration: none;
  text-align: center;
  display: inline-block;
  padding-top: 0.75em;
  padding-right: 1.5em;
  padding-bottom: 0.8125em;
  padding-left: 1.5em;
  font-size: 1em;
  background-color: #2ba6cb;
  border-color: #2284a1;
  color: white; }
  button:hover, button:focus, .button:hover, .button:focus {
    background-color: #2284a1; }
  button:hover, button:focus, .button:hover, .button:focus {
    color: white; }
  button.secondary, .button.secondary {
    background-color: #e9e9e9;
    border-color: #d0d0d0;
    color: #333333; }
    button.secondary:hover, button.secondary:focus, .button.secondary:hover, .button.secondary:focus {
      background-color: #d0d0d0; }
    button.secondary:hover, button.secondary:focus, .button.secondary:hover, .button.secondary:focus {
      color: #333333; }
  button.success, .button.success {
    background-color: #5da423;
    border-color: #457a1a;
    color: white; }
    button.success:hover, button.success:focus, .button.success:hover, .button.success:focus {
      background-color: #457a1a; }
    button.success:hover, button.success:focus, .button.success:hover, .button.success:focus {
      color: white; }
  button.alert, .button.alert {
    background-color: #c60f13;
    border-color: #970b0e;
    color: white; }
    button.alert:hover, button.alert:focus, .button.alert:hover, .button.alert:focus {
      background-color: #970b0e; }
    button.alert:hover, button.alert:focus, .button.alert:hover, .button.alert:focus {
      color: white; }
  button.large, .button.large {
    padding-top: 1em;
    padding-right: 2em;
    padding-bottom: 1.0625em;
    padding-left: 2em;
    font-size: 1.25em; }
  button.small, .button.small {
    padding-top: 0.5625em;
    padding-right: 1.125em;
    padding-bottom: 0.625em;
    padding-left: 1.125em;
    font-size: 0.8125em; }
  button.tiny, .button.tiny {
    padding-top: 0.4375em;
    padding-right: 0.875em;
    padding-bottom: 0.5em;
    padding-left: 0.875em;
    font-size: 0.6875em; }
  button.expand, .button.expand {
    padding-top: false;
    padding-right: 0px;
    padding-bottom: false0.0625em;
    padding-left: 0px;
    width: 100%; }
  button.left-align, .button.left-align {
    text-align: left;
    text-indent: 0.75em; }
  button.right-align, .button.right-align {
    text-align: right;
    padding-right: 0.75em; }
  button.disabled, button[disabled], .button.disabled, .button[disabled] {
    background-color: #2ba6cb;
    border-color: #2284a1;
    color: white;
    cursor: default;
    opacity: 0.6;
    -webkit-box-shadow: none;
    box-shadow: none; }
    button.disabled:hover, button.disabled:focus, button[disabled]:hover, button[disabled]:focus, .button.disabled:hover, .button.disabled:focus, .button[disabled]:hover, .button[disabled]:focus {
      background-color: #2284a1; }
    button.disabled:hover, button.disabled:focus, button[disabled]:hover, button[disabled]:focus, .button.disabled:hover, .button.disabled:focus, .button[disabled]:hover, .button[disabled]:focus {
      color: white; }
    button.disabled:hover, button.disabled:focus, button[disabled]:hover, button[disabled]:focus, .button.disabled:hover, .button.disabled:focus, .button[disabled]:hover, .button[disabled]:focus {
      background-color: #2ba6cb; }
    button.disabled.secondary, button[disabled].secondary, .button.disabled.secondary, .button[disabled].secondary {
      background-color: #e9e9e9;
      border-color: #d0d0d0;
      color: #333333;
      cursor: default;
      opacity: 0.6;
      -webkit-box-shadow: none;
      box-shadow: none; }
      button.disabled.secondary:hover, button.disabled.secondary:focus, button[disabled].secondary:hover, button[disabled].secondary:focus, .button.disabled.secondary:hover, .button.disabled.secondary:focus, .button[disabled].secondary:hover, .button[disabled].secondary:focus {
        background-color: #d0d0d0; }
      button.disabled.secondary:hover, button.disabled.secondary:focus, button[disabled].secondary:hover, button[disabled].secondary:focus, .button.disabled.secondary:hover, .button.disabled.secondary:focus, .button[disabled].secondary:hover, .button[disabled].secondary:focus {
        color: #333333; }
      button.disabled.secondary:hover, button.disabled.secondary:focus, button[disabled].secondary:hover, button[disabled].secondary:focus, .button.disabled.secondary:hover, .button.disabled.secondary:focus, .button[disabled].secondary:hover, .button[disabled].secondary:focus {
        background-color: #e9e9e9; }
    button.disabled.success, button[disabled].success, .button.disabled.success, .button[disabled].success {
      background-color: #5da423;
      border-color: #457a1a;
      color: white;
      cursor: default;
      opacity: 0.6;
      -webkit-box-shadow: none;
      box-shadow: none; }
      button.disabled.success:hover, button.disabled.success:focus, button[disabled].success:hover, button[disabled].success:focus, .button.disabled.success:hover, .button.disabled.success:focus, .button[disabled].success:hover, .button[disabled].success:focus {
        background-color: #457a1a; }
      button.disabled.success:hover, button.disabled.success:focus, button[disabled].success:hover, button[disabled].success:focus, .button.disabled.success:hover, .button.disabled.success:focus, .button[disabled].success:hover, .button[disabled].success:focus {
        color: white; }
      button.disabled.success:hover, button.disabled.success:focus, button[disabled].success:hover, button[disabled].success:focus, .button.disabled.success:hover, .button.disabled.success:focus, .button[disabled].success:hover, .button[disabled].success:focus {
        background-color: #5da423; }
    button.disabled.alert, button[disabled].alert, .button.disabled.alert, .button[disabled].alert {
      background-color: #c60f13;
      border-color: #970b0e;
      color: white;
      cursor: default;
      opacity: 0.6;
      -webkit-box-shadow: none;
      box-shadow: none; }
      button.disabled.alert:hover, button.disabled.alert:focus, button[disabled].alert:hover, button[disabled].alert:focus, .button.disabled.alert:hover, .button.disabled.alert:focus, .button[disabled].alert:hover, .button[disabled].alert:focus {
        background-color: #970b0e; }
      button.disabled.alert:hover, button.disabled.alert:focus, button[disabled].alert:hover, button[disabled].alert:focus, .button.disabled.alert:hover, .button.disabled.alert:focus, .button[disabled].alert:hover, .button[disabled].alert:focus {
        color: white; }
      button.disabled.alert:hover, button.disabled.alert:focus, button[disabled].alert:hover, button[disabled].alert:focus, .button.disabled.alert:hover, .button.disabled.alert:focus, .button[disabled].alert:hover, .button[disabled].alert:focus {
        background-color: #c60f13; }

input.button,
button.button {
  padding-top: 0.8125em;
  padding-bottom: 0.75em; }
  input.button.tiny,
  button.button.tiny {
    padding-top: 0.5em;
    padding-bottom: 0.4375em; }
  input.button.small,
  button.button.small {
    padding-top: 0.625em;
    padding-bottom: 0.5625em; }
  input.button.large,
  button.button.large {
    padding-top: 1.03125em;
    padding-bottom: 1.03125em; }

@media only screen {
  .button {
    -webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.5) inset;
    box-shadow: 0 1px 0 rgba(255, 255, 255, 0.5) inset;
    -webkit-transition: background-color 300ms ease-out;
    -moz-transition: background-color 300ms ease-out;
    transition: background-color 300ms ease-out; }
    .button:active {
      -webkit-box-shadow: 0 1px 0 rgba(0, 0, 0, 0.2) inset;
      box-shadow: 0 1px 0 rgba(0, 0, 0, 0.2) inset; }
    .button.radius {
      -webkit-border-radius: 3px;
      border-radius: 3px; }
    .button.round {
      -webkit-border-radius: 1000px;
      border-radius: 1000px; } }
@media only screen and (min-width: 48em) {
  .button {
    display: inline-block; } }
/* Standard Forms */
form {
  margin: 0 0 1em; }

/* Using forms within rows, we need to set some defaults */
form .row .row {
  margin: -0.5em; }
  form .row .row .column,
  form .row .row .columns {
    padding: 0 0.5em; }
  form .row .row.collapse {
    margin: 0; }
    form .row .row.collapse .column,
    form .row .row.collapse .columns {
      padding: 0; }
form .row input.column,
form .row input.columns {
  padding-left: 0.5em; }

/* Label Styles */
label {
  font-size: 0.875em;
  color: #4d4d4d;
  cursor: pointer;
  display: block;
  font-weight: 500;
  margin-bottom: 0.1875em; }
  label.right {
    float: none;
    text-align: right; }
  label.inline {
    margin: 0 0 1em 0;
    padding: 0.625em 0; }

/* Attach elements to the beginning or end of an input */
.prefix,
.postfix {
  display: block;
  position: relative;
  z-index: 2;
  text-align: center;
  width: 100%;
  padding-top: 0;
  padding-bottom: 0;
  border-style: solid;
  border-width: 1px;
  overflow: hidden;
  font-size: 0.875em;
  height: 2.3125em;
  line-height: 2.3125em; }

/* Adjust padding, alignment and radius if pre/post element is a button */
.postfix.button {
  padding-left: 0;
  padding-right: 0;
  padding-top: 0;
  padding-bottom: 0;
  text-align: center;
  line-height: 2.125em; }

.prefix.button {
  padding-left: 0;
  padding-right: 0;
  padding-top: 0;
  padding-bottom: 0;
  text-align: center;
  line-height: 2.125em; }

.prefix.button.radius {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-bottomleft: 3px;
  -moz-border-radius-topleft: 3px;
  -webkit-border-bottom-left-radius: 3px;
  -webkit-border-top-left-radius: 3px;
  border-bottom-left-radius: 3px;
  border-top-left-radius: 3px; }

.postfix.button.radius {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-topright: 3px;
  -moz-border-radius-bottomright: 3px;
  -webkit-border-top-right-radius: 3px;
  -webkit-border-bottom-right-radius: 3px;
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px; }

.prefix.button.round {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-bottomleft: 1000px;
  -moz-border-radius-topleft: 1000px;
  -webkit-border-bottom-left-radius: 1000px;
  -webkit-border-top-left-radius: 1000px;
  border-bottom-left-radius: 1000px;
  border-top-left-radius: 1000px; }

.postfix.button.round {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-topright: 1000px;
  -moz-border-radius-bottomright: 1000px;
  -webkit-border-top-right-radius: 1000px;
  -webkit-border-bottom-right-radius: 1000px;
  border-top-right-radius: 1000px;
  border-bottom-right-radius: 1000px; }

/* Separate prefix and postfix styles when on span so buttons keep their own */
span.prefix {
  background: #f2f2f2;
  border-color: #d9d9d9;
  border-right: none;
  color: #333333; }
  span.prefix.radius {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-bottomleft: 3px;
    -moz-border-radius-topleft: 3px;
    -webkit-border-bottom-left-radius: 3px;
    -webkit-border-top-left-radius: 3px;
    border-bottom-left-radius: 3px;
    border-top-left-radius: 3px; }

span.postfix {
  background: #f2f2f2;
  border-color: #cccccc;
  border-left: none;
  color: #333333; }
  span.postfix.radius {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-topright: 3px;
    -moz-border-radius-bottomright: 3px;
    -webkit-border-top-right-radius: 3px;
    -webkit-border-bottom-right-radius: 3px;
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px; }

/* Input groups will automatically style first and last elements of the group */
.input-group.radius > *:first-child, .input-group.radius > *:first-child * {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-bottomleft: 3px;
  -moz-border-radius-topleft: 3px;
  -webkit-border-bottom-left-radius: 3px;
  -webkit-border-top-left-radius: 3px;
  border-bottom-left-radius: 3px;
  border-top-left-radius: 3px; }
.input-group.radius > *:last-child, .input-group.radius > *:last-child * {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-topright: 3px;
  -moz-border-radius-bottomright: 3px;
  -webkit-border-top-right-radius: 3px;
  -webkit-border-bottom-right-radius: 3px;
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px; }
.input-group.round > *:first-child, .input-group.round > *:first-child * {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-bottomleft: 1000px;
  -moz-border-radius-topleft: 1000px;
  -webkit-border-bottom-left-radius: 1000px;
  -webkit-border-top-left-radius: 1000px;
  border-bottom-left-radius: 1000px;
  border-top-left-radius: 1000px; }
.input-group.round > *:last-child, .input-group.round > *:last-child * {
  -webkit-border-radius: 0;
  border-radius: 0;
  -moz-border-radius-topright: 1000px;
  -moz-border-radius-bottomright: 1000px;
  -webkit-border-top-right-radius: 1000px;
  -webkit-border-bottom-right-radius: 1000px;
  border-top-right-radius: 1000px;
  border-bottom-right-radius: 1000px; }

/* We use this to get basic styling on all basic form elements */
input[type="text"],
input[type="password"],
input[type="date"],
input[type="datetime"],
input[type="datetime-local"],
input[type="month"],
input[type="week"],
input[type="email"],
input[type="number"],
input[type="search"],
input[type="tel"],
input[type="time"],
input[type="url"],
textarea {
  background-color: white;
  font-family: inherit;
  border: 1px solid #cccccc;
  -webkit-box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
  color: rgba(0, 0, 0, 0.75);
  display: block;
  font-size: 0.875em;
  margin: 0 0 1em 0;
  padding: 0.5em;
  height: 2.3125em;
  width: 100%;
  -moz-box-sizing: border-box;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  -webkit-transition: all 0.15s linear;
  -moz-transition: all 0.15s linear;
  transition: all 0.15s linear; }
  input[type="text"]:focus,
  input[type="password"]:focus,
  input[type="date"]:focus,
  input[type="datetime"]:focus,
  input[type="datetime-local"]:focus,
  input[type="month"]:focus,
  input[type="week"]:focus,
  input[type="email"]:focus,
  input[type="number"]:focus,
  input[type="search"]:focus,
  input[type="tel"]:focus,
  input[type="time"]:focus,
  input[type="url"]:focus,
  textarea:focus {
    background: #fafafa;
    border-color: #999999;
    outline: none; }
  input[type="text"][disabled],
  input[type="password"][disabled],
  input[type="date"][disabled],
  input[type="datetime"][disabled],
  input[type="datetime-local"][disabled],
  input[type="month"][disabled],
  input[type="week"][disabled],
  input[type="email"][disabled],
  input[type="number"][disabled],
  input[type="search"][disabled],
  input[type="tel"][disabled],
  input[type="time"][disabled],
  input[type="url"][disabled],
  textarea[disabled] {
    background-color: #dddddd; }

/* We add basic fieldset styling */
fieldset {
  border: solid 1px #dddddd;
  padding: 1.25em;
  margin: 1.125em 0; }
  fieldset legend {
    font-weight: bold;
    background: white;
    padding: 0 0.1875em;
    margin: 0;
    margin-left: -0.1875em; }

/* Error Handling */
.error input,
input.error,
.error textarea,
textarea.error {
  border-color: #c60f13;
  background-color: rgba(198, 15, 19, 0.1); }
  .error input:focus,
  input.error:focus,
  .error textarea:focus,
  textarea.error:focus {
    background: #fafafa;
    border-color: #999999; }

.error label,
label.error {
  color: #c60f13; }

.error small,
small.error {
  display: block;
  padding: 0.375em 0.25em;
  margin-top: -1.3125em;
  margin-bottom: 1em;
  font-size: 0.75em;
  font-weight: bold;
  background: #c60f13;
  color: white; }

/* Custom Checkbox and Radio Inputs */
form.custom .custom {
  display: inline-block;
  width: 16px;
  height: 16px;
  position: relative;
  top: 2px;
  border: solid 1px #cccccc;
  background: white; }
  form.custom .custom.radio {
    -webkit-border-radius: 1000px;
    border-radius: 1000px; }
  form.custom .custom.checkbox:before {
    content: "";
    display: block;
    line-height: 0.8;
    height: 14px;
    width: 14px;
    text-align: center;
    position: absolute;
    top: 0;
    left: 0;
    font-size: 14px;
    color: #fff; }
  form.custom .custom.radio.checked:before {
    content: "";
    display: block;
    width: 8px;
    height: 8px;
    -webkit-border-radius: 1000px;
    border-radius: 1000px;
    background: #222222;
    position: relative;
    top: 3px;
    left: 3px; }
  form.custom .custom.checkbox.checked:before {
    content: "\00d7";
    color: #222222; }

/* Custom Select Options and Dropdowns */
form.custom {
  /* Custom input, disabled */ }
  form.custom .custom.dropdown {
    display: block;
    position: relative;
    top: 0;
    height: 2.3125em;
    margin-bottom: 1.25em;
    margin-top: 0px;
    padding: 0px;
    width: 100%;
    background: white;
    background: -moz-linear-gradient(top, white 0%, #f3f3f3 100%);
    background: -webkit-linear-gradient(top, white 0%, #f3f3f3 100%);
    background: linear-gradient(to bottom, white 0%, #f3f3f3 100%);
    -webkit-box-shadow: none;
    box-shadow: none;
    font-size: 0.875em;
    vertical-align: top; }
    form.custom .custom.dropdown ul {
      overflow-y: auto;
      max-height: 200px; }
    form.custom .custom.dropdown .current {
      cursor: default;
      white-space: nowrap;
      line-height: 2.25em;
      color: rgba(0, 0, 0, 0.75);
      text-decoration: none;
      overflow: hidden;
      display: block;
      margin-left: 0.5em;
      margin-right: 2.3125em; }
    form.custom .custom.dropdown .selector {
      cursor: default;
      position: absolute;
      width: 2.5em;
      height: 2.3125em;
      display: block;
      right: 0;
      top: 0; }
      form.custom .custom.dropdown .selector:after {
        content: "";
        display: block;
        content: "";
        display: block;
        width: 0;
        height: 0;
        border: solid 5px;
        border-color: #aaaaaa transparent transparent transparent;
        position: absolute;
        left: 0.9375em;
        top: 50%;
        margin-top: -3px; }
    form.custom .custom.dropdown:hover a.selector:after, form.custom .custom.dropdown.open a.selector:after {
      content: "";
      display: block;
      width: 0;
      height: 0;
      border: solid 5px;
      border-color: #222222 transparent transparent transparent; }
    form.custom .custom.dropdown .disabled {
      color: #888888; }
      form.custom .custom.dropdown .disabled:hover {
        background: transparent;
        color: #888888; }
        form.custom .custom.dropdown .disabled:hover:after {
          display: none; }
    form.custom .custom.dropdown.open ul {
      display: block;
      z-index: 10;
      min-width: 100%;
      -moz-box-sizing: content-box;
      -webkit-box-sizing: content-box;
      box-sizing: content-box; }
    form.custom .custom.dropdown.small {
      max-width: 134px; }
    form.custom .custom.dropdown.medium {
      max-width: 254px; }
    form.custom .custom.dropdown.large {
      max-width: 434px; }
    form.custom .custom.dropdown.expand {
      width: 100% !important; }
    form.custom .custom.dropdown.open.small ul {
      min-width: 134px;
      -moz-box-sizing: border-box;
      -webkit-box-sizing: border-box;
      box-sizing: border-box; }
    form.custom .custom.dropdown.open.medium ul {
      min-width: 254px;
      -moz-box-sizing: border-box;
      -webkit-box-sizing: border-box;
      box-sizing: border-box; }
    form.custom .custom.dropdown.open.large ul {
      min-width: 434px;
      -moz-box-sizing: border-box;
      -webkit-box-sizing: border-box;
      box-sizing: border-box; }
  form.custom .custom.dropdown ul {
    position: absolute;
    width: auto;
    display: none;
    margin: 0;
    left: -1px;
    top: auto;
    -webkit-box-shadow: 0 2px 2px 0px rgba(0, 0, 0, 0.1);
    box-shadow: 0 2px 2px 0px rgba(0, 0, 0, 0.1);
    margin: 0;
    padding: 0;
    background: white;
    border: solid 1px #cccccc;
    font-size: 16px; }
    form.custom .custom.dropdown ul li {
      color: #555555;
      font-size: 0.875em;
      cursor: default;
      padding-top: 0.25em;
      padding-bottom: 0.25em;
      padding-left: 0.375em;
      padding-right: 2.375em;
      min-height: 1.5em;
      line-height: 1.5em;
      margin: 0;
      white-space: nowrap;
      list-style: none; }
      form.custom .custom.dropdown ul li.selected {
        background: #eeeeee;
        color: black; }
      form.custom .custom.dropdown ul li:hover {
        background-color: #e4e4e4;
        color: black; }
      form.custom .custom.dropdown ul li.selected:hover {
        background: #eeeeee;
        cursor: default;
        color: black; }
    form.custom .custom.dropdown ul.show {
      display: block; }
  form.custom .custom.disabled {
    background-color: #dddddd; }

/* Button Groups */
.button-group {
  list-style: none;
  margin: 0;
  *zoom: 1; }
  .button-group:before, .button-group:after {
    content: " ";
    display: table; }
  .button-group:after {
    clear: both; }
  .button-group > * {
    margin: 0 0 0 -1px;
    float: left; }
    .button-group > *:first-child {
      margin-left: 0; }
  .button-group.radius > *:first-child, .button-group.radius > *:first-child > a, .button-group.radius > *:first-child > button {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-bottomleft: 3px;
    -moz-border-radius-topleft: 3px;
    -webkit-border-bottom-left-radius: 3px;
    -webkit-border-top-left-radius: 3px;
    border-bottom-left-radius: 3px;
    border-top-left-radius: 3px; }
  .button-group.radius > *:last-child, .button-group.radius > *:last-child > a, .button-group.radius > *:last-child > button {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-topright: 3px;
    -moz-border-radius-bottomright: 3px;
    -webkit-border-top-right-radius: 3px;
    -webkit-border-bottom-right-radius: 3px;
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px; }
  .button-group.round > *:first-child, .button-group.round > *:first-child > a, .button-group.round > *:first-child > button {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-bottomleft: 1000px;
    -moz-border-radius-topleft: 1000px;
    -webkit-border-bottom-left-radius: 1000px;
    -webkit-border-top-left-radius: 1000px;
    border-bottom-left-radius: 1000px;
    border-top-left-radius: 1000px; }
  .button-group.round > *:last-child, .button-group.round > *:last-child > a, .button-group.round > *:last-child > button {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-topright: 1000px;
    -moz-border-radius-bottomright: 1000px;
    -webkit-border-top-right-radius: 1000px;
    -webkit-border-bottom-right-radius: 1000px;
    border-top-right-radius: 1000px;
    border-bottom-right-radius: 1000px; }
  .button-group.even-2 li {
    width: 50%; }
    .button-group.even-2 li .button {
      width: 100%; }
  .button-group.even-3 li {
    width: 33.33333%; }
    .button-group.even-3 li .button {
      width: 100%; }
  .button-group.even-4 li {
    width: 25%; }
    .button-group.even-4 li .button {
      width: 100%; }
  .button-group.even-5 li {
    width: 20%; }
    .button-group.even-5 li .button {
      width: 100%; }
  .button-group.even-6 li {
    width: 16.66667%; }
    .button-group.even-6 li .button {
      width: 100%; }
  .button-group.even-7 li {
    width: 14.28571%; }
    .button-group.even-7 li .button {
      width: 100%; }
  .button-group.even-8 li {
    width: 12.5%; }
    .button-group.even-8 li .button {
      width: 100%; }

.button-bar {
  *zoom: 1; }
  .button-bar:before, .button-bar:after {
    content: " ";
    display: table; }
  .button-bar:after {
    clear: both; }
  .button-bar .button-group {
    float: left;
    margin-right: 0.625em; }
    .button-bar .button-group div {
      overflow: hidden; }

/* Dropdown Button */
.dropdown.button {
  position: relative;
  padding-right: 3.1875em; }
  .dropdown.button:before {
    position: absolute;
    content: "";
    width: 0;
    height: 0;
    display: block;
    border-style: solid;
    border-color: white transparent transparent transparent;
    top: 50%; }
  .dropdown.button:before {
    border-width: 0.5625em;
    right: 1.5em;
    margin-top: -0.25em; }
  .dropdown.button:before {
    border-color: white transparent transparent transparent; }
  .dropdown.button.tiny {
    padding-right: 2.1875em; }
    .dropdown.button.tiny:before {
      border-width: 0.4375em;
      right: 0.875em;
      margin-top: -0.15625em; }
    .dropdown.button.tiny:before {
      border-color: white transparent transparent transparent; }
  .dropdown.button.small {
    padding-right: 2.8125em; }
    .dropdown.button.small:before {
      border-width: 0.5625em;
      right: 1.125em;
      margin-top: -0.21875em; }
    .dropdown.button.small:before {
      border-color: white transparent transparent transparent; }
  .dropdown.button.large {
    padding-right: 4em; }
    .dropdown.button.large:before {
      border-width: 0.625em;
      right: 1.75em;
      margin-top: -0.3125em; }
    .dropdown.button.large:before {
      border-color: white transparent transparent transparent; }
  .dropdown.button.secondary:before {
    border-color: #333333 transparent transparent transparent; }

/* Split Buttons */
.split.button {
  position: relative;
  padding-right: 4.8em; }
  .split.button span {
    display: block;
    height: 100%;
    position: absolute;
    right: 0;
    top: 0;
    border-left: solid 1px; }
    .split.button span:before {
      position: absolute;
      content: "";
      width: 0;
      height: 0;
      display: block;
      border-style: solid;
      left: 50%; }
    .split.button span:active {
      background-color: rgba(0, 0, 0, 0.1); }
  .split.button span {
    border-left-color: #1e728c; }
  .split.button span {
    width: 3em; }
    .split.button span:before {
      border-width: 0.5625em;
      top: 1.125em;
      margin-left: -0.5625em; }
  .split.button span:before {
    border-color: white transparent transparent transparent; }
  .split.button.secondary span {
    border-left-color: #c3c3c3; }
  .split.button.secondary span:before {
    border-color: white transparent transparent transparent; }
  .split.button.alert span {
    border-left-color: #7f0a0c; }
  .split.button.success span {
    border-left-color: #396516; }
  .split.button.tiny {
    padding-right: 3.9375em; }
    .split.button.tiny span {
      width: 2.84375em; }
      .split.button.tiny span:before {
        border-width: 0.4375em;
        top: 0.875em;
        margin-left: -0.3125em; }
  .split.button.small {
    padding-right: 3.9375em; }
    .split.button.small span {
      width: 2.8125em; }
      .split.button.small span:before {
        border-width: 0.5625em;
        top: 0.84375em;
        margin-left: -0.5625em; }
  .split.button.large {
    padding-right: 6em; }
    .split.button.large span {
      width: 3.75em; }
      .split.button.large span:before {
        border-width: 0.625em;
        top: 1.3125em;
        margin-left: -0.5625em; }
  .split.button.secondary span:before {
    border-color: #333333 transparent transparent transparent; }
  .split.button.radius span {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-topright: 3px;
    -moz-border-radius-bottomright: 3px;
    -webkit-border-top-right-radius: 3px;
    -webkit-border-bottom-right-radius: 3px;
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px; }
  .split.button.round span {
    -webkit-border-radius: 0;
    border-radius: 0;
    -moz-border-radius-topright: 1000px;
    -moz-border-radius-bottomright: 1000px;
    -webkit-border-top-right-radius: 1000px;
    -webkit-border-bottom-right-radius: 1000px;
    border-top-right-radius: 1000px;
    border-bottom-right-radius: 1000px; }

/* Flex Video */
.flex-video {
  position: relative;
  padding-top: 1.5625em;
  padding-bottom: 67.5%;
  height: 0;
  margin-bottom: 1em;
  overflow: hidden; }
  .flex-video.widescreen {
    padding-bottom: 57.25%; }
  .flex-video.vimeo {
    padding-top: 0; }
  .flex-video iframe,
  .flex-video object,
  .flex-video embed,
  .flex-video video {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%; }

/* Sections */
.section-container, .section-container.auto {
  width: 100%;
  display: block;
  margin-bottom: 1.25em;
  border: 1px solid #cccccc;
  border-top: none; }
  .section-container section,
  .section-container .section, .section-container.auto section,
  .section-container.auto .section {
    border-top: 1px solid #cccccc;
    position: relative; }
    .section-container section .title,
    .section-container .section .title, .section-container.auto section .title,
    .section-container.auto .section .title {
      top: 0;
      cursor: pointer;
      width: 100%;
      margin: 0;
      background-color: #efefef; }
      .section-container section .title a,
      .section-container .section .title a, .section-container.auto section .title a,
      .section-container.auto .section .title a {
        padding: 0.9375em;
        display: inline-block;
        color: #333333;
        font-size: 0.875em;
        white-space: nowrap;
        width: 100%; }
      .section-container section .title:hover,
      .section-container .section .title:hover, .section-container.auto section .title:hover,
      .section-container.auto .section .title:hover {
        background-color: #e2e2e2; }
    .section-container section .content,
    .section-container .section .content, .section-container.auto section .content,
    .section-container.auto .section .content {
      display: none;
      padding: 0.9375em;
      background-color: white; }
      .section-container section .content > *:last-child,
      .section-container .section .content > *:last-child, .section-container.auto section .content > *:last-child,
      .section-container.auto .section .content > *:last-child {
        margin-bottom: 0; }
      .section-container section .content > *:first-child,
      .section-container .section .content > *:first-child, .section-container.auto section .content > *:first-child,
      .section-container.auto .section .content > *:first-child {
        padding-top: 0; }
      .section-container section .content > *:last-child,
      .section-container .section .content > *:last-child, .section-container.auto section .content > *:last-child,
      .section-container.auto .section .content > *:last-child {
        padding-bottom: 0; }
    .section-container section.active .content,
    .section-container .section.active .content, .section-container.auto section.active .content,
    .section-container.auto .section.active .content {
      display: block; }
    .section-container section.active .title,
    .section-container .section.active .title, .section-container.auto section.active .title,
    .section-container.auto .section.active .title {
      background: #d5d5d5; }

.section-container.tabs {
  border: 0;
  position: relative; }
  .section-container.tabs section,
  .section-container.tabs .section {
    padding-top: 0;
    border: 0;
    position: static; }
    .section-container.tabs section .title,
    .section-container.tabs .section .title {
      width: auto;
      border: 1px solid #cccccc;
      border-right: 0;
      border-bottom: 0;
      position: absolute;
      z-index: 1; }
      .section-container.tabs section .title a,
      .section-container.tabs .section .title a {
        width: 100%; }
    .section-container.tabs section:last-child .title,
    .section-container.tabs .section:last-child .title {
      border-right: 1px solid #cccccc; }
    .section-container.tabs section .content,
    .section-container.tabs .section .content {
      border: 1px solid #cccccc;
      position: absolute;
      z-index: 10;
      top: -1px; }
    .section-container.tabs section.active .title,
    .section-container.tabs .section.active .title {
      background-color: white;
      z-index: 11;
      border-bottom: 0; }
    .section-container.tabs section.active .content,
    .section-container.tabs .section.active .content {
      position: relative; }

@media only screen and (min-width: 48em) {
  .section-container.auto {
    border: 0;
    position: relative; }
    .section-container.auto section,
    .section-container.auto .section {
      padding-top: 0;
      border: 0;
      position: static; }
      .section-container.auto section .title,
      .section-container.auto .section .title {
        width: auto;
        border: 1px solid #cccccc;
        border-right: 0;
        border-bottom: 0;
        position: absolute;
        z-index: 1; }
        .section-container.auto section .title a,
        .section-container.auto .section .title a {
          width: 100%; }
      .section-container.auto section:last-child .title,
      .section-container.auto .section:last-child .title {
        border-right: 1px solid #cccccc; }
      .section-container.auto section .content,
      .section-container.auto .section .content {
        border: 1px solid #cccccc;
        position: absolute;
        z-index: 10;
        top: -1px; }
      .section-container.auto section.active .title,
      .section-container.auto .section.active .title {
        background-color: white;
        z-index: 11;
        border-bottom: 0; }
      .section-container.auto section.active .content,
      .section-container.auto .section.active .content {
        position: relative; }

  .section-container.accordion .section {
    padding-top: 0 !important; }

  .section-container.vertical-nav {
    border: 1px solid #cccccc;
    border-top: none; }
    .section-container.vertical-nav section,
    .section-container.vertical-nav .section {
      padding-top: 0 !important; }
      .section-container.vertical-nav section .title a,
      .section-container.vertical-nav .section .title a {
        display: block;
        width: 100%; }
      .section-container.vertical-nav section .content,
      .section-container.vertical-nav .section .content {
        display: none; }
      .section-container.vertical-nav section.active .content,
      .section-container.vertical-nav .section.active .content {
        display: block;
        position: absolute;
        left: 100%;
        top: -1px;
        z-index: 999;
        min-width: 12.5em;
        border: 1px solid #cccccc; }

  .section-container.horizontal-nav {
    position: relative;
    background: #efefef;
    border: 1px solid #cccccc; }
    .section-container.horizontal-nav section,
    .section-container.horizontal-nav .section {
      padding-top: 0;
      border: 0;
      position: static; }
      .section-container.horizontal-nav section .title,
      .section-container.horizontal-nav .section .title {
        width: auto;
        border: 1px solid #cccccc;
        border-left: 0;
        top: -1px;
        position: absolute;
        z-index: 1; }
        .section-container.horizontal-nav section .title a,
        .section-container.horizontal-nav .section .title a {
          width: 100%; }
      .section-container.horizontal-nav section .content,
      .section-container.horizontal-nav .section .content {
        display: none; }
      .section-container.horizontal-nav section.active .content,
      .section-container.horizontal-nav .section.active .content {
        display: block;
        position: absolute;
        z-index: 999;
        left: 0;
        top: -2px;
        min-width: 12.5em;
        border: 1px solid #cccccc; } }
/* Wrapped around .top-bar to contain to grid width */
.contain-to-grid {
  width: 100%;
  background: #111111; }

.fixed {
  width: 100%;
  left: 0;
  position: fixed;
  top: 0;
  z-index: 99; }

.top-bar {
  overflow: hidden;
  height: 45px;
  line-height: 45px;
  position: relative;
  background: #111111;
  margin-bottom: 1.875em; }
  .top-bar ul {
    margin-bottom: 0;
    list-style: none; }
  .top-bar .row {
    max-width: none; }
  .top-bar form,
  .top-bar input {
    margin-bottom: 0; }
  .top-bar input {
    height: 2.45em; }
  .top-bar .button {
    padding-top: .5em;
    padding-bottom: .5em;
    margin-bottom: 0; }
  .top-bar .title-area {
    position: relative; }
  .top-bar .name {
    height: 45px;
    margin: 0;
    font-size: 16px; }
    .top-bar .name h1 {
      line-height: 45px;
      font-size: 1.0625em;
      margin: 0; }
      .top-bar .name h1 a {
        font-weight: bold;
        color: white;
        width: 50%;
        display: block;
        padding: 0 15px; }
  .top-bar .toggle-topbar {
    position: absolute;
    right: 0;
    top: 0; }
    .top-bar .toggle-topbar a {
      color: white;
      text-transform: uppercase;
      font-size: 0.8125em;
      font-weight: bold;
      position: relative;
      display: block;
      padding: 0 15px;
      height: 45px;
      line-height: 45px; }
    .top-bar .toggle-topbar.menu-icon {
      right: 15px;
      top: 50%;
      margin-top: -16px;
      padding-left: 40px; }
      .top-bar .toggle-topbar.menu-icon a {
        text-indent: -48px;
        width: 34px;
        height: 34px;
        line-height: 33px;
        padding: 0;
        color: white; }
        .top-bar .toggle-topbar.menu-icon a span {
          position: absolute;
          right: 0;
          display: block;
          width: 16px;
          height: 0;
          -webkit-box-shadow: 0 10px 0 1px white, 0 16px 0 1px white, 0 22px 0 1px white;
          box-shadow: 0 10px 0 1px white, 0 16px 0 1px white, 0 22px 0 1px white; }
  .top-bar.expanded {
    height: auto;
    background: transparent; }
    .top-bar.expanded .title-area {
      background: #111111; }
    .top-bar.expanded .toggle-topbar a {
      color: #888888; }
      .top-bar.expanded .toggle-topbar a span {
        -webkit-box-shadow: 0 10px 0 1px #888888, 0 16px 0 1px #888888, 0 22px 0 1px #888888;
        box-shadow: 0 10px 0 1px #888888, 0 16px 0 1px #888888, 0 22px 0 1px #888888; }

.top-bar-section {
  left: 0;
  position: relative;
  width: auto;
  -webkit-transition: left 300ms ease-out;
  -moz-transition: left 300ms ease-out;
  transition: left 300ms ease-out; }
  .top-bar-section ul {
    width: 100%;
    height: auto;
    display: block;
    background: #333333;
    font-size: 16px;
    margin: 0; }
  .top-bar-section .divider {
    border-bottom: solid 1px #4d4d4d;
    border-top: solid 1px #1a1a1a;
    clear: both;
    height: 1px;
    width: 100%; }
  .top-bar-section ul li > a {
    display: block;
    width: 100%;
    color: white;
    padding: 12px 0 12px 0;
    padding-left: 15px;
    font-size: 0.8125em;
    font-weight: bold;
    background: #333333;
    height: 45px; }
    .top-bar-section ul li > a:hover {
      background: #2b2b2b; }
    .top-bar-section ul li > a.button {
      background: #2ba6cb;
      font-size: 0.8125em; }
      .top-bar-section ul li > a.button:hover {
        background: #2284a1; }
    .top-bar-section ul li > a.button.secondary {
      background: #e9e9e9; }
      .top-bar-section ul li > a.button.secondary:hover {
        background: #d0d0d0; }
    .top-bar-section ul li > a.button.success {
      background: #5da423; }
      .top-bar-section ul li > a.button.success:hover {
        background: #457a1a; }
    .top-bar-section ul li > a.button.alert {
      background: #c60f13; }
      .top-bar-section ul li > a.button.alert:hover {
        background: #970b0e; }
  .top-bar-section ul li.active a {
    background: #2b2b2b; }
  .top-bar-section .has-form {
    padding: 15px; }
  .top-bar-section .has-dropdown {
    position: relative; }
    .top-bar-section .has-dropdown > a:after {
      content: "";
      display: block;
      width: 0;
      height: 0;
      border: solid 5px;
      border-color: transparent transparent transparent rgba(255, 255, 255, 0.5);
      margin-right: 15px;
      margin-top: -4.5px;
      position: absolute;
      top: 22px;
      right: 0; }
    .top-bar-section .has-dropdown.moved {
      position: static; }
      .top-bar-section .has-dropdown.moved > .dropdown {
        visibility: visible; }
  .top-bar-section .dropdown {
    position: absolute;
    left: 100%;
    top: 0;
    visibility: hidden;
    z-index: 99; }
    .top-bar-section .dropdown li {
      width: 100%; }
      .top-bar-section .dropdown li a {
        font-weight: normal;
        padding: 8px 15px; }
      .top-bar-section .dropdown li.title h5 {
        margin-bottom: 0; }
        .top-bar-section .dropdown li.title h5 a {
          color: white;
          line-height: 22.5px;
          display: block; }
    .top-bar-section .dropdown label {
      padding: 8px 15px 2px;
      margin-bottom: 0;
      text-transform: uppercase;
      color: #555555;
      font-weight: bold;
      font-size: 0.625em; }

.top-bar-js-breakpoint {
  width: 58.75em !important;
  visibility: hidden; }

.js-generated {
  display: block; }

@media only screen and (min-width: 58.75em) {
  .top-bar {
    background: #111111;
    *zoom: 1;
    overflow: visible; }
    .top-bar:before, .top-bar:after {
      content: " ";
      display: table; }
    .top-bar:after {
      clear: both; }
    .top-bar .toggle-topbar {
      display: none; }
    .top-bar .title-area {
      float: left; }
    .top-bar .name h1 a {
      width: auto; }
    .top-bar input,
    .top-bar .button {
      line-height: 2em;
      font-size: 0.875em;
      height: 2em;
      padding: 0 10px;
      position: relative;
      top: 8px; }
    .top-bar.expanded {
      background: #111111; }

  .contain-to-grid .top-bar {
    max-width: 62.5em;
    margin: 0 auto; }

  .top-bar-section {
    -webkit-transition: none 0 0;
    -moz-transition: none 0 0;
    transition: none 0 0;
    left: 0 !important; }
    .top-bar-section ul {
      width: auto;
      height: auto !important;
      display: inline; }
      .top-bar-section ul li {
        float: left; }
        .top-bar-section ul li .js-generated {
          display: none; }
    .top-bar-section li a:not(.button) {
      padding: 0 15px;
      line-height: 45px;
      background: #111111; }
      .top-bar-section li a:not(.button):hover {
        background: black; }
    .top-bar-section .has-dropdown > a {
      padding-right: 35px !important; }
      .top-bar-section .has-dropdown > a:after {
        content: "";
        display: block;
        width: 0;
        height: 0;
        border: solid 5px;
        border-color: rgba(255, 255, 255, 0.5) transparent transparent transparent;
        margin-top: -2.5px; }
    .top-bar-section .has-dropdown.moved {
      position: relative; }
      .top-bar-section .has-dropdown.moved > .dropdown {
        visibility: hidden; }
    .top-bar-section .has-dropdown:hover > .dropdown, .top-bar-section .has-dropdown:active > .dropdown {
      visibility: visible; }
    .top-bar-section .has-dropdown .dropdown li.has-dropdown > a:after {
      border: none;
      content: "\00bb";
      margin-top: -7px;
      right: 5px; }
    .top-bar-section .dropdown {
      left: 0;
      top: auto;
      background: transparent;
      min-width: 100%; }
      .top-bar-section .dropdown li a {
        color: white;
        line-height: 1;
        white-space: nowrap;
        padding: 7px 15px;
        background: #1e1e1e; }
      .top-bar-section .dropdown li label {
        white-space: nowrap;
        background: #1e1e1e; }
      .top-bar-section .dropdown li .dropdown {
        left: 100%;
        top: 0; }
    .top-bar-section > ul > .divider {
      border-bottom: none;
      border-top: none;
      border-right: solid 1px #2b2b2b;
      border-left: solid 1px black;
      clear: none;
      height: 45px;
      width: 0px; }
    .top-bar-section .has-form {
      background: #111111;
      padding: 0 15px;
      height: 45px; }
    .top-bar-section ul.right li .dropdown {
      left: auto;
      right: 0; }
      .top-bar-section ul.right li .dropdown li .dropdown {
        right: 100%; } }
.orbit-container {
  overflow: hidden;
  width: 100%;
  position: relative;
  background: whitesmoke; }
  .orbit-container .orbit-slides-container {
    list-style: none;
    margin: 0;
    padding: 0;
    position: relative; }
    .orbit-container .orbit-slides-container img {
      display: block; }
    .orbit-container .orbit-slides-container > * {
      position: relative;
      float: left;
      height: 100%; }
      .orbit-container .orbit-slides-container > * .orbit-caption {
        position: absolute;
        bottom: 0;
        background-color: black;
        background-color: rgba(0, 0, 0, 0.6);
        color: #fff;
        width: 100%;
        padding: 10px 14px;
        font-size: 0.875em; }
        .orbit-container .orbit-slides-container > * .orbit-caption * {
          color: white; }
  .orbit-container .orbit-slide-number {
    position: absolute;
    top: 10px;
    left: 10px;
    font-size: 12px; }
    .orbit-container .orbit-slide-number span {
      font-weight: 700; }
  .orbit-container .orbit-timer {
    position: absolute;
    top: 10px;
    right: 10px;
    height: 6px;
    width: 100px; }
    .orbit-container .orbit-timer .orbit-progress {
      height: 100%;
      background-color: black;
      background-color: rgba(0, 0, 0, 0.6);
      display: block;
      width: 0%; }
    .orbit-container .orbit-timer > span {
      display: none;
      position: absolute;
      top: 10px;
      right: 0px;
      width: 11px;
      height: 14px;
      border: solid 4px black;
      border-top: none;
      border-bottom: none; }
    .orbit-container .orbit-timer.paused > span {
      right: -6px;
      top: 9px;
      width: 11px;
      height: 14px;
      border: solid 8px;
      border-color: transparent transparent transparent black; }
  .orbit-container:hover .orbit-timer > span {
    display: block; }
  .orbit-container .orbit-prev,
  .orbit-container .orbit-next {
    position: absolute;
    top: 50%;
    margin-top: -25px;
    background-color: black;
    background-color: rgba(0, 0, 0, 0.6);
    width: 50px;
    height: 60px;
    line-height: 50px;
    color: white;
    text-indent: -9999px !important; }
    .orbit-container .orbit-prev > span,
    .orbit-container .orbit-next > span {
      position: absolute;
      top: 50%;
      margin-top: -16px;
      display: block;
      width: 0;
      height: 0;
      border: solid 16px; }
  .orbit-container .orbit-prev {
    left: 0; }
    .orbit-container .orbit-prev > span {
      border-color: transparent;
      border-right-color: #fff; }
    .orbit-container .orbit-prev:hover > span {
      border-right-color: #ccc; }
  .orbit-container .orbit-next {
    right: 0; }
    .orbit-container .orbit-next > span {
      border-color: transparent;
      border-left-color: #fff;
      left: 50%;
      margin-left: -8px; }
    .orbit-container .orbit-next:hover > span {
      border-left-color: #ccc; }

.orbit-bullets {
  margin: 0 auto 30px auto;
  overflow: hidden;
  position: relative;
  top: 10px; }
  .orbit-bullets li {
    display: block;
    width: 18px;
    height: 18px;
    background: #fff;
    float: left;
    margin-right: 6px;
    border: solid 2px black;
    -webkit-border-radius: 1000px;
    border-radius: 1000px; }
    .orbit-bullets li.active {
      background: #000; }
    .orbit-bullets li:last-child {
      margin-right: 0; }

.touch .orbit-container .orbit-prev,
.touch .orbit-container .orbit-next {
  display: none; }
.touch .orbit-bullets {
  display: none; }

@media only screen and (min-width: 48em) {
  .touch .orbit-container .orbit-prev,
  .touch .orbit-container .orbit-next {
    display: inherit; }
  .touch .orbit-bullets {
    display: block; } }
.reveal-modal-bg {
  position: fixed;
  height: 100%;
  width: 100%;
  background: black;
  background: rgba(0, 0, 0, 0.45);
  z-index: 98;
  display: none;
  top: 0;
  left: 0; }

.reveal-modal {
  visibility: hidden;
  display: none;
  position: absolute;
  left: 50%;
  z-index: 99;
  height: auto;
  background-color: #fff;
  margin-left: -40%;
  width: 80%;
  background-color: white;
  padding: 1.25em;
  border: solid 1px #666666;
  -webkit-box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
  top: 50px; }
  .reveal-modal .column,
  .reveal-modal .columns {
    min-width: 0; }
  .reveal-modal > :first-child {
    margin-top: 0; }
  .reveal-modal > :last-child {
    margin-bottom: 0; }
  .reveal-modal .close-reveal-modal {
    font-size: 1.375em;
    line-height: 1;
    position: absolute;
    top: 0.5em;
    right: 0.6875em;
    color: #aaaaaa;
    font-weight: bold;
    cursor: pointer; }

@media only screen and (min-width: 48em) {
  .reveal-modal {
    padding: 1.875em;
    top: 6.25em; }
    .reveal-modal.small {
      margin-left: -15%;
      width: 30%; }
    .reveal-modal.medium {
      margin-left: -20%;
      width: 40%; }
    .reveal-modal.large {
      margin-left: -30%;
      width: 60%; }
    .reveal-modal.xlarge {
      margin-left: -35%;
      width: 70%; }
    .reveal-modal.expand {
      margin-left: -47.5%;
      width: 95%; } }
@media print {
  div:not(.reveal-modal) {
    display: none; } }
/* Foundation Joyride */
.joyride-list {
  display: none; }

/* Default styles for the container */
.joyride-tip-guide {
  display: none;
  position: absolute;
  background: black;
  color: white;
  z-index: 101;
  top: 0;
  left: 2.5%;
  font-family: inherit;
  font-weight: normal;
  width: 95%; }

.lt-ie9 .joyride-tip-guide {
  max-width: 800px;
  left: 50%;
  margin-left: -400px; }

.joyride-content-wrapper {
  width: 100%;
  padding: 1.125em 1.25em 1.5em; }
  .joyride-content-wrapper .button {
    margin-bottom: 0 !important; }

/* Add a little css triangle pip, older browser just miss out on the fanciness of it */
.joyride-tip-guide .joyride-nub {
  display: block;
  position: absolute;
  left: 22px;
  width: 0;
  height: 0;
  border: solid 14px; }
  .joyride-tip-guide .joyride-nub.top {
    border-color: black;
    border-top-color: transparent !important;
    border-left-color: transparent !important;
    border-right-color: transparent !important;
    top: -28px;
    bottom: none; }
  .joyride-tip-guide .joyride-nub.bottom {
    border-color: black !important;
    border-bottom-color: transparent !important;
    border-left-color: transparent !important;
    border-right-color: transparent !important;
    bottom: -28px;
    bottom: none; }
  .joyride-tip-guide .joyride-nub.right {
    right: -28px; }
  .joyride-tip-guide .joyride-nub.left {
    left: -28px; }

/* Typography */
.joyride-tip-guide h1,
.joyride-tip-guide h2,
.joyride-tip-guide h3,
.joyride-tip-guide h4,
.joyride-tip-guide h5,
.joyride-tip-guide h6 {
  line-height: 1.25;
  margin: 0;
  font-weight: bold;
  color: white; }

.joyride-tip-guide p {
  margin: 0 0 1.125em 0;
  font-size: 0.875em;
  line-height: 1.3; }

.joyride-timer-indicator-wrap {
  width: 50px;
  height: 3px;
  border: solid 1px #555555;
  position: absolute;
  right: 1.0625em;
  bottom: 1em; }

.joyride-timer-indicator {
  display: block;
  width: 0;
  height: inherit;
  background: #666666; }

.joyride-close-tip {
  position: absolute;
  right: 12px;
  top: 10px;
  color: #777777 !important;
  text-decoration: none;
  font-size: 30px;
  font-weight: normal;
  line-height: 0.5 !important; }
  .joyride-close-tip:hover, .joyride-close-tip:focus {
    color: #eeeeee !important; }

.joyride-modal-bg {
  position: fixed;
  height: 100%;
  width: 100%;
  background: transparent;
  background: rgba(0, 0, 0, 0.5);
  z-index: 100;
  display: none;
  top: 0;
  left: 0;
  cursor: pointer; }

/* Styles for screens that are atleast 768px; */
@media only screen and (min-width: 48em) {
  .joyride-tip-guide {
    width: 300px;
    left: inherit; }
    .joyride-tip-guide .joyride-nub.bottom {
      border-color: black !important;
      border-bottom-color: transparent !important;
      border-left-color: transparent !important;
      border-right-color: transparent !important;
      bottom: -28px;
      bottom: none; }
    .joyride-tip-guide .joyride-nub.right {
      border-color: black !important;
      border-top-color: transparent !important;
      border-right-color: transparent !important;
      border-bottom-color: transparent !important;
      top: 22px;
      bottom: none;
      left: auto;
      right: -28px; }
    .joyride-tip-guide .joyride-nub.left {
      border-color: black !important;
      border-top-color: transparent !important;
      border-left-color: transparent !important;
      border-bottom-color: transparent !important;
      top: 22px;
      left: -28px;
      right: auto;
      bottom: none; } }
/* Clearing Styles */
[data-clearing] {
  *zoom: 1;
  margin-bottom: 0; }
  [data-clearing]:before, [data-clearing]:after {
    content: " ";
    display: table; }
  [data-clearing]:after {
    clear: both; }

.clearing-blackout {
  background: #111111;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 998; }
  .clearing-blackout .clearing-close {
    display: block; }

.clearing-container {
  position: relative;
  z-index: 998;
  height: 100%;
  overflow: hidden;
  margin: 0; }

.visible-img {
  height: 95%;
  position: relative; }
  .visible-img img {
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -50%;
    max-height: 100%;
    max-width: 100%; }

.clearing-caption {
  color: white;
  line-height: 1.3;
  margin-bottom: 0;
  text-align: center;
  bottom: 0;
  background: #111111;
  width: 100%;
  padding: 10px 30px;
  position: absolute;
  left: 0; }

.clearing-close {
  z-index: 999;
  padding-left: 20px;
  padding-top: 10px;
  font-size: 40px;
  line-height: 1;
  color: white;
  display: none; }
  .clearing-close:hover, .clearing-close:focus {
    color: #ccc; }

.clearing-assembled .clearing-container {
  height: 100%; }
  .clearing-assembled .clearing-container .carousel > ul {
    display: none; }

@media only screen and (min-width: 48em) {
  .clearing-main-prev,
  .clearing-main-next {
    position: absolute;
    height: 100%;
    width: 40px;
    top: 0; }
    .clearing-main-prev > span,
    .clearing-main-next > span {
      position: absolute;
      top: 50%;
      display: block;
      width: 0;
      height: 0;
      border: solid 16px; }

  .clearing-main-prev {
    left: 0; }
    .clearing-main-prev > span {
      left: 5px;
      border-color: transparent;
      border-right-color: white; }

  .clearing-main-next {
    right: 0; }
    .clearing-main-next > span {
      border-color: transparent;
      border-left-color: white; }

  .clearing-main-prev.disabled,
  .clearing-main-next.disabled {
    opacity: 0.5; }

  .clearing-feature ~ li {
    display: none; }

  .clearing-assembled .clearing-container .carousel {
    background: #111111;
    height: 150px;
    margin-top: 5px; }
    .clearing-assembled .clearing-container .carousel > ul {
      display: block;
      z-index: 999;
      width: 200%;
      height: 100%;
      margin-left: 0;
      position: relative;
      left: 0; }
      .clearing-assembled .clearing-container .carousel > ul li {
        display: block;
        width: 175px;
        height: inherit;
        padding: 0;
        float: left;
        overflow: hidden;
        margin-right: 1px;
        position: relative;
        cursor: pointer;
        opacity: 0.4; }
        .clearing-assembled .clearing-container .carousel > ul li.fix-height img {
          min-height: 100%;
          height: 100%;
          max-width: none; }
        .clearing-assembled .clearing-container .carousel > ul li a.th {
          border: none;
          -webkit-box-shadow: none;
          box-shadow: none;
          display: block; }
        .clearing-assembled .clearing-container .carousel > ul li img {
          cursor: pointer !important;
          min-width: 100% !important; }
        .clearing-assembled .clearing-container .carousel > ul li.visible {
          opacity: 1; }
  .clearing-assembled .clearing-container .visible-img {
    background: #111111;
    overflow: hidden;
    height: 75%; }

  .clearing-close {
    position: absolute;
    top: 10px;
    right: 20px;
    padding-left: 0;
    padding-top: 0; } }
/* Foundation Alerts */
.alert-box {
  border-style: solid;
  border-width: 1px;
  display: block;
  font-weight: bold;
  margin-bottom: 1.25em;
  position: relative;
  padding: 0.6875em 1.3125em 0.75em 0.6875em;
  font-size: 0.875em;
  background-color: #2ba6cb;
  border-color: #2284a1;
  color: white; }
  .alert-box .close {
    font-size: 1.375em;
    padding: 5px 4px 4px;
    line-height: 0;
    position: absolute;
    top: 0.4375em;
    right: 0.3125em;
    color: #333333;
    opacity: 0.3; }
    .alert-box .close:hover, .alert-box .close:focus {
      opacity: 0.5; }
  .alert-box.radius {
    -webkit-border-radius: 3px;
    border-radius: 3px; }
  .alert-box.round {
    -webkit-border-radius: 1000px;
    border-radius: 1000px; }
  .alert-box.success {
    background-color: #5da423;
    border-color: #457a1a;
    color: white; }
  .alert-box.alert {
    background-color: #c60f13;
    border-color: #970b0e;
    color: white; }
  .alert-box.secondary {
    background-color: #e9e9e9;
    border-color: #d0d0d0;
    color: #505050; }

/* Breadcrumbs */
.breadcrumbs {
  display: block;
  padding: 0.375em 0.875em 0.5625em;
  overflow: hidden;
  margin-left: 0;
  list-style: none;
  border-style: solid;
  border-width: 1px;
  background-color: #f6f6f6;
  border-color: gainsboro;
  -webkit-border-radius: 3px;
  border-radius: 3px; }
  .breadcrumbs li {
    margin: 0;
    padding: 0 0.75em 0 0;
    float: left; }
    .breadcrumbs li:hover a, .breadcrumbs li:focus a {
      text-decoration: underline; }
    .breadcrumbs li a,
    .breadcrumbs li span {
      font-size: 0.6875em;
      padding-left: 0.75em;
      text-transform: uppercase;
      color: #2ba6cb; }
    .breadcrumbs li.current a {
      cursor: default;
      color: #333333; }
    .breadcrumbs li.current:hover a, .breadcrumbs li.current:focus a {
      text-decoration: none; }
    .breadcrumbs li.unavailable a {
      color: #999999; }
    .breadcrumbs li.unavailable:hover a,
    .breadcrumbs li.unavailable a:focus {
      text-decoration: none;
      color: #999999;
      cursor: default; }
    .breadcrumbs li:before {
      content: "/";
      color: #aaaaaa;
      position: relative;
      top: 1px; }
    .breadcrumbs li:first-child a, .breadcrumbs li:first-child span {
      padding-left: 0; }
    .breadcrumbs li:first-child:before {
      content: " "; }

/* Keystroke Characters */
.keystroke,
kbd {
  background-color: #ededed;
  border-color: #dbdbdb;
  color: #222222;
  border-style: solid;
  border-width: 1px;
  margin: 0;
  font-family: "Consolas", "Menlo", "Courier", monospace;
  font-size: 0.9375em;
  padding: 0.125em 0.25em 0em;
  -webkit-border-radius: 3px;
  border-radius: 3px; }

/* Labels */
.label {
  font-weight: 500;
  text-align: center;
  text-decoration: none;
  line-height: 1;
  white-space: nowrap;
  display: inline-block;
  position: relative;
  padding: 0.1875em 0.625em 0.25em;
  font-size: 0.875em;
  background-color: #2ba6cb;
  color: #fff; }
  .label.radius {
    -webkit-border-radius: 3px;
    border-radius: 3px; }
  .label.round {
    -webkit-border-radius: 1000px;
    border-radius: 1000px; }
  .label.alert {
    background-color: #c60f13;
    color: #fff; }
  .label.success {
    background-color: #5da423;
    color: #fff; }
  .label.secondary {
    background-color: #e9e9e9;
    color: #333; }

/* Inline Lists */
.inline-list {
  margin: 0 auto 1.0625em auto;
  margin-left: -1.375em;
  margin-right: 0;
  padding: 0;
  list-style: none;
  overflow: hidden; }
  .inline-list > li {
    list-style: none;
    float: left;
    margin-left: 1.375em;
    display: block; }
    .inline-list > li > * {
      display: block; }

/* Pagination */
.pagination {
  display: block;
  height: 1.5em;
  margin-left: -0.3125em; }
  .pagination li {
    display: block;
    float: left;
    height: 1.5em;
    color: #222222;
    font-size: 0.875em;
    margin-left: 0.3125em; }
    .pagination li a {
      display: block;
      padding: 0.0625em 0.4375em 0.0625em;
      color: #999999; }
    .pagination li:hover a,
    .pagination li a:focus {
      background: #e6e6e6; }
    .pagination li.unavailable a {
      cursor: default;
      color: #999999; }
    .pagination li.unavailable:hover a, .pagination li.unavailable a:focus {
      background: transparent; }
    .pagination li.current a {
      background: #2ba6cb;
      color: white;
      font-weight: bold;
      cursor: default; }
      .pagination li.current a:hover, .pagination li.current a:focus {
        background: #2ba6cb; }

.pagination-centered {
  text-align: center; }
  .pagination-centered ul > li {
    float: none;
    display: inline-block; }

/* Panels */
.panel {
  border-style: solid;
  border-width: 1px;
  border-color: #d9d9d9;
  margin-bottom: 1.25em;
  padding: 1.25em;
  background: #f2f2f2; }
  .panel h1, .panel h2, .panel h3, .panel h4, .panel h5, .panel h6, .panel p {
    color: #333333; }
  .panel > :first-child {
    margin-top: 0; }
  .panel > :last-child {
    margin-bottom: 0; }
  .panel h1, .panel h2, .panel h3, .panel h4, .panel h5, .panel h6 {
    line-height: 1;
    margin-bottom: 0.625em; }
    .panel h1.subheader, .panel h2.subheader, .panel h3.subheader, .panel h4.subheader, .panel h5.subheader, .panel h6.subheader {
      line-height: 1.4; }
  .panel.callout {
    border-style: solid;
    border-width: 1px;
    border-color: #2284a1;
    margin-bottom: 1.25em;
    padding: 1.25em;
    background: #2ba6cb;
    -webkit-box-shadow: 0 1px 0 rgba(255, 255, 255, 0.5) inset;
    box-shadow: 0 1px 0 rgba(255, 255, 255, 0.5) inset; }
    .panel.callout h1, .panel.callout h2, .panel.callout h3, .panel.callout h4, .panel.callout h5, .panel.callout h6, .panel.callout p {
      color: white; }
    .panel.callout > :first-child {
      margin-top: 0; }
    .panel.callout > :last-child {
      margin-bottom: 0; }
    .panel.callout h1, .panel.callout h2, .panel.callout h3, .panel.callout h4, .panel.callout h5, .panel.callout h6 {
      line-height: 1;
      margin-bottom: 0.625em; }
      .panel.callout h1.subheader, .panel.callout h2.subheader, .panel.callout h3.subheader, .panel.callout h4.subheader, .panel.callout h5.subheader, .panel.callout h6.subheader {
        line-height: 1.4; }
  .panel.radius {
    -webkit-border-radius: 3px;
    border-radius: 3px; }

/* Pricing Tables */
.pricing-table {
  border: solid 1px #dddddd;
  margin-left: 0;
  margin-bottom: 1.25em; }
  .pricing-table * {
    list-style: none;
    line-height: 1; }
  .pricing-table .title {
    background-color: #dddddd;
    padding: 0.9375em 1.25em;
    text-align: center;
    color: #333333;
    font-weight: bold;
    font-size: 1em; }
  .pricing-table .price {
    background-color: #eeeeee;
    padding: 0.9375em 1.25em;
    text-align: center;
    color: #333333;
    font-weight: normal;
    font-size: 1.25em; }
  .pricing-table .description {
    background-color: white;
    padding: 0.9375em;
    text-align: center;
    color: #777777;
    font-size: 0.75em;
    font-weight: normal;
    line-height: 1.4;
    border-bottom: dotted 1px #dddddd; }
  .pricing-table .bullet-item {
    background-color: white;
    padding: 0.9375em;
    text-align: center;
    color: #333333;
    font-size: 0.875em;
    font-weight: normal;
    border-bottom: dotted 1px #dddddd; }
  .pricing-table .cta-button {
    background-color: whitesmoke;
    text-align: center;
    padding: 1.25em 1.25em 0; }

/* Progress Bar */
.progress {
  background-color: transparent;
  height: 1.5625em;
  border: 1px solid #cccccc;
  padding: 0.125em;
  margin-bottom: 0.625em; }
  .progress .meter {
    background: #2ba6cb;
    height: 100%;
    display: block; }
  .progress.secondary .meter {
    background: #e9e9e9;
    height: 100%;
    display: block; }
  .progress.success .meter {
    background: #5da423;
    height: 100%;
    display: block; }
  .progress.alert .meter {
    background: #c60f13;
    height: 100%;
    display: block; }
  .progress.radius {
    -webkit-border-radius: 3px;
    border-radius: 3px; }
    .progress.radius .meter {
      -webkit-border-radius: 2px;
      border-radius: 2px; }
  .progress.round {
    -webkit-border-radius: 1000px;
    border-radius: 1000px; }
    .progress.round .meter {
      -webkit-border-radius: 999px;
      border-radius: 999px; }

/* Side Nav */
.side-nav {
  display: block;
  margin: 0;
  padding: 0.875em 0;
  list-style-type: none;
  list-style-position: inside; }
  .side-nav li {
    margin: 0 0 0.4375em 0;
    font-size: 0.875em; }
    .side-nav li a {
      display: block;
      color: #2ba6cb; }
    .side-nav li.active a {
      color: #4d4d4d;
      font-weight: bold; }
    .side-nav li.divider {
      border-top: 1px solid;
      height: 0;
      padding: 0;
      list-style: none;
      border-top-color: #e6e6e6; }

/* Side Nav */
.sub-nav {
  display: block;
  width: auto;
  overflow: hidden;
  margin: -0.25em 0 1.125em;
  padding-top: 0.25em;
  margin-right: 0;
  margin-left: -0.5625em; }
  .sub-nav dt,
  .sub-nav dd {
    float: left;
    display: inline;
    margin-left: 0.5625em;
    margin-bottom: 0.625em;
    font-weight: normal;
    font-size: 0.875em; }
    .sub-nav dt a,
    .sub-nav dd a {
      color: #999999;
      text-decoration: none; }
    .sub-nav dt.active a,
    .sub-nav dd.active a {
      -webkit-border-radius: 1000px;
      border-radius: 1000px;
      font-weight: bold;
      background: #2ba6cb;
      padding: 0.1875em 0.5625em;
      cursor: default;
      color: white; }

/* Foundation Switches */
@media only screen {
  div.switch {
    position: relative;
    width: 100%;
    padding: 0;
    display: block;
    overflow: hidden;
    border-style: solid;
    border-width: 1px;
    margin-bottom: 1.25em;
    -webkit-animation: webkitSiblingBugfix infinite 1s;
    height: 36px;
    background: white;
    border-color: #cccccc; }
    div.switch label {
      position: relative;
      left: 0;
      z-index: 2;
      float: left;
      width: 50%;
      height: 100%;
      margin: 0;
      font-weight: bold;
      text-align: left;
      -webkit-transition: all 0.1s ease-out;
      -moz-transition: all 0.1s ease-out;
      transition: all 0.1s ease-out; }
    div.switch input {
      position: absolute;
      z-index: 3;
      opacity: 0;
      width: 100%;
      height: 100%; }
      div.switch input:hover, div.switch input:focus {
        cursor: pointer; }
    div.switch > span {
      position: absolute;
      top: -1px;
      left: -1px;
      z-index: 1;
      display: block;
      padding: 0;
      border-width: 1px;
      border-style: solid;
      -webkit-transition: all 0.1s ease-out;
      -moz-transition: all 0.1s ease-out;
      transition: all 0.1s ease-out; }
    div.switch input:not(:checked) + label {
      opacity: 0; }
    div.switch input:checked {
      display: none !important; }
    div.switch input {
      left: 0;
      display: block !important; }
    div.switch input:first-of-type + label,
    div.switch input:first-of-type + span + label {
      left: -50%; }
    div.switch input:first-of-type:checked + label,
    div.switch input:first-of-type:checked + span + label {
      left: 0%; }
    div.switch input:last-of-type + label,
    div.switch input:last-of-type + span + label {
      right: -50%;
      left: auto;
      text-align: right; }
    div.switch input:last-of-type:checked + label,
    div.switch input:last-of-type:checked + span + label {
      right: 0%;
      left: auto; }
    div.switch span.custom {
      display: none !important; }
    div.switch label {
      padding: 0 0.375em;
      line-height: 2.3em;
      font-size: 0.875em; }
    div.switch input:first-of-type:checked ~ span {
      left: 100%;
      margin-left: -2.1875em; }
    div.switch > span {
      width: 2.25em;
      height: 2.25em; }
    div.switch > span {
      border-color: #b3b3b3;
      background: white;
      background: -moz-linear-gradient(top, white 0%, #f2f2f2 100%);
      background: -webkit-linear-gradient(top, white 0%, #f2f2f2 100%);
      background: linear-gradient(to bottom, white 0%, #f2f2f2 100%);
      -webkit-box-shadow: 2px 0 10px 0 rgba(0, 0, 0, 0.07), 1000px 0 0 1000px #e1f5d1, -2px 0 10px 0 rgba(0, 0, 0, 0.07), -1000px 0 0 1000px whitesmoke;
      box-shadow: 2px 0 10px 0 rgba(0, 0, 0, 0.07), 1000px 0 0 980px #e1f5d1, -2px 0 10px 0 rgba(0, 0, 0, 0.07), -1000px 0 0 1000px whitesmoke; }
    div.switch:hover > span, div.switch:focus > span {
      background: white;
      background: -moz-linear-gradient(top, white 0%, #e6e6e6 100%);
      background: -webkit-linear-gradient(top, white 0%, #e6e6e6 100%);
      background: linear-gradient(to bottom, white 0%, #e6e6e6 100%); }
    div.switch:active {
      background: transparent; }
    div.switch.large {
      height: 44px; }
      div.switch.large label {
        padding: 0 0.375em;
        line-height: 2.3em;
        font-size: 1.0625em; }
      div.switch.large input:first-of-type:checked ~ span {
        left: 100%;
        margin-left: -2.6875em; }
      div.switch.large > span {
        width: 2.75em;
        height: 2.75em; }
    div.switch.small {
      height: 28px; }
      div.switch.small label {
        padding: 0 0.375em;
        line-height: 2.1em;
        font-size: 0.75em; }
      div.switch.small input:first-of-type:checked ~ span {
        left: 100%;
        margin-left: -1.6875em; }
      div.switch.small > span {
        width: 1.75em;
        height: 1.75em; }
    div.switch.tiny {
      height: 22px; }
      div.switch.tiny label {
        padding: 0 0.375em;
        line-height: 1.9em;
        font-size: 0.6875em; }
      div.switch.tiny input:first-of-type:checked ~ span {
        left: 100%;
        margin-left: -1.3125em; }
      div.switch.tiny > span {
        width: 1.375em;
        height: 1.375em; }
    div.switch.radius {
      -webkit-border-radius: 4px;
      border-radius: 4px; }
      div.switch.radius > span {
        -webkit-border-radius: 3px;
        border-radius: 3px; }
    div.switch.round {
      -webkit-border-radius: 1000px;
      border-radius: 1000px; }
      div.switch.round > span {
        -webkit-border-radius: 999px;
        border-radius: 999px; }
      div.switch.round label {
        padding: 0 0.5625em; }

  @-webkit-keyframes webkitSiblingBugfix {
    from {
      position: relative; }

    to {
      position: relative; } } }
[data-magellan-expedition] {
  background: white;
  z-index: 50;
  min-width: 100%;
  padding: 10px; }
  [data-magellan-expedition] .sub-nav {
    margin-bottom: 0; }
    [data-magellan-expedition] .sub-nav dd {
      margin-bottom: 0; }

/* Tables */
table {
  background: white;
  margin-bottom: 1.25em;
  border: solid 1px #dddddd; }
  table thead,
  table tfoot {
    background: whitesmoke;
    font-weight: bold; }
    table thead tr th,
    table thead tr td,
    table tfoot tr th,
    table tfoot tr td {
      padding: 0.5em 0.625em 0.625em;
      font-size: 0.875em;
      color: #222222;
      text-align: left; }
  table tr th,
  table tr td {
    padding: 0.5625em 0.625em;
    font-size: 0.875em;
    color: #222222; }
  table tr.even, table tr.alt, table tr:nth-of-type(even) {
    background: #f9f9f9; }
  table thead tr th,
  table tfoot tr th,
  table tbody tr td,
  table tr td,
  table tfoot tr td {
    display: table-cell;
    line-height: 1.125em; }

/* Image Thumbnails */
.th {
  display: inline-block;
  border: solid 4px white;
  -webkit-box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.2);
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.2);
  -webkit-transition: all 200ms ease-out;
  -moz-transition: all 200ms ease-out;
  transition: all 200ms ease-out; }
  .th:hover, .th:focus {
    -webkit-box-shadow: 0 0 6px 1px rgba(43, 166, 203, 0.5);
    box-shadow: 0 0 6px 1px rgba(43, 166, 203, 0.5); }
  .th.radius {
    -webkit-border-radius: 3px;
    border-radius: 3px; }

/* Tooltips */
.has-tip {
  border-bottom: dotted 1px #cccccc;
  cursor: help;
  font-weight: bold;
  color: #333333; }
  .has-tip:hover, .has-tip:focus {
    border-bottom: dotted 1px #196177;
    color: #2ba6cb; }
  .has-tip.tip-left, .has-tip.tip-right {
    float: none !important; }

.tooltip {
  display: none;
  position: absolute;
  z-index: 999;
  font-weight: bold;
  font-size: 0.9375em;
  line-height: 1.3;
  padding: 0.5em;
  max-width: 85%;
  left: 50%;
  width: 100%;
  color: white;
  background: black;
  -webkit-border-radius: 3px;
  border-radius: 3px; }
  .tooltip > .nub {
    display: block;
    left: 5px;
    position: absolute;
    width: 0;
    height: 0;
    border: solid 5px;
    border-color: transparent transparent black transparent;
    top: -10px; }
  .tooltip.opened {
    color: #2ba6cb !important;
    border-bottom: dotted 1px #196177 !important; }

.tap-to-close {
  display: block;
  font-size: 0.625em;
  color: #888888;
  font-weight: normal; }

@media only screen and (min-width: 48em) {
  .tooltip > .nub {
    border-color: transparent transparent black transparent;
    top: -10px; }
  .tooltip.tip-top > .nub {
    border-color: black transparent transparent transparent;
    top: auto;
    bottom: -10px; }
  .tooltip.tip-left, .tooltip.tip-right {
    float: none !important; }
  .tooltip.tip-left > .nub {
    border-color: transparent transparent transparent black;
    right: -10px;
    left: auto;
    top: 50%;
    margin-top: -5px; }
  .tooltip.tip-right > .nub {
    border-color: transparent black transparent transparent;
    right: auto;
    left: -10px;
    top: 50%;
    margin-top: -5px; } }
@media only screen and (max-width: 767px) {
  .f-dropdown {
    max-width: 100%;
    left: 0; } }
/* Foundation Dropdowns */
.f-dropdown {
  position: absolute;
  top: -9999px;
  list-style: none;
  padding: 1.25em;
  width: 100%;
  height: auto;
  max-height: none;
  background: white;
  border: solid 1px #cccccc;
  font-size: 16px;
  z-index: 99;
  margin-top: 2px;
  max-width: 200px; }
  .f-dropdown *:first-child {
    margin-top: 0; }
  .f-dropdown *:last-child {
    margin-bottom: 0; }
  .f-dropdown:before {
    content: "";
    display: block;
    width: 0;
    height: 0;
    border: solid 6px;
    border-color: transparent transparent white transparent;
    position: absolute;
    top: -12px;
    left: 10px;
    z-index: 99; }
  .f-dropdown:after {
    content: "";
    display: block;
    width: 0;
    height: 0;
    border: solid 7px;
    border-color: transparent transparent #cccccc transparent;
    position: absolute;
    top: -14px;
    left: 9px;
    z-index: 98; }
  .f-dropdown.right:before {
    left: auto;
    right: 10px; }
  .f-dropdown.right:after {
    left: auto;
    right: 9px; }
  .f-dropdown li {
    font-size: 0.875em;
    cursor: pointer;
    padding: 0.3125em 0.625em;
    line-height: 1.125em;
    margin: 0; }
    .f-dropdown li:hover, .f-dropdown li:focus {
      background: #eeeeee; }
    .f-dropdown li a {
      color: #555555; }
  .f-dropdown.content {
    position: absolute;
    top: -9999px;
    list-style: none;
    padding: 1.25em;
    width: 100%;
    height: auto;
    max-height: none;
    background: white;
    border: solid 1px #cccccc;
    font-size: 16px;
    z-index: 99;
    max-width: 200px; }
    .f-dropdown.content *:first-child {
      margin-top: 0; }
    .f-dropdown.content *:last-child {
      margin-bottom: 0; }
  .f-dropdown.tiny {
    max-width: 200px; }
  .f-dropdown.small {
    max-width: 300px; }
  .f-dropdown.medium {
    max-width: 500px; }
  .f-dropdown.large {
    max-width: 800px; }`
	NORMALIZE_CSS = `/*! normalize.css v2.1.0 | MIT License | git.io/normalize */

/* ==========================================================================
   HTML5 display definitions
   ========================================================================== */

/**
 * Correct block display not defined in IE 8/9.
 */

article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
main,
nav,
section,
summary {
    display: block;
}

/**
 * Correct inline-block display not defined in IE 8/9.
 */

audio,
canvas,
video {
    display: inline-block;
}

/**
 * Prevent modern browsers from displaying audio without controls.
 * Remove excess height in iOS 5 devices.
 */

audio:not([controls]) {
    display: none;
    height: 0;
}

/**
 * Address styling not present in IE 8/9.
 */

[hidden] {
    display: none;
}

/* ==========================================================================
   Base
   ========================================================================== */

/**
 * 1. Set default font family to sans-serif.
 * 2. Prevent iOS text size adjust after orientation change, without disabling
 *    user zoom.
 */

html {
    font-family: sans-serif; /* 1 */
    -webkit-text-size-adjust: 100%; /* 2 */
    -ms-text-size-adjust: 100%; /* 2 */
}

/**
 * Remove default margin.
 */

body {
    margin: 0;
}

/* ==========================================================================
   Links
   ========================================================================== */

/**
 * Address outline inconsistency between Chrome and other browsers.
 */

a:focus {
    outline: thin dotted;
}

/**
 * Improve readability when focused and also mouse hovered in all browsers.
 */

a:active,
a:hover {
    outline: 0;
}

/* ==========================================================================
   Typography
   ========================================================================== */

/**
 * Address variable h1 font-size and margin within section and article
 * contexts in Firefox 4+, Safari 5, and Chrome.
 */

h1 {
    font-size: 2em;
    margin: 0.67em 0;
}

/**
 * Address styling not present in IE 8/9, Safari 5, and Chrome.
 */

abbr[title] {
    border-bottom: 1px dotted;
}

/**
 * Address style set to bolder in Firefox 4+, Safari 5, and Chrome.
 */

b,
strong {
    font-weight: bold;
}

/**
 * Address styling not present in Safari 5 and Chrome.
 */

dfn {
    font-style: italic;
}

/**
 * Address differences between Firefox and other browsers.
 */

hr {
    -moz-box-sizing: content-box;
    box-sizing: content-box;
    height: 0;
}

/**
 * Address styling not present in IE 8/9.
 */

mark {
    background: #ff0;
    color: #000;
}

/**
 * Correct font family set oddly in Safari 5 and Chrome.
 */

code,
kbd,
pre,
samp {
    font-family: monospace, serif;
    font-size: 1em;
}

/**
 * Improve readability of pre-formatted text in all browsers.
 */

pre {
    white-space: pre-wrap;
}

/**
 * Set consistent quote types.
 */

q {
    quotes: "\201C" "\201D" "\2018" "\2019";
}

/**
 * Address inconsistent and variable font size in all browsers.
 */

small {
    font-size: 80%;
}

/**
 * Prevent sub and sup affecting line-height in all browsers.
 */

sub,
sup {
    font-size: 75%;
    line-height: 0;
    position: relative;
    vertical-align: baseline;
}

sup {
    top: -0.5em;
}

sub {
    bottom: -0.25em;
}

/* ==========================================================================
   Embedded content
   ========================================================================== */

/**
 * Remove border when inside a element in IE 8/9.
 */

img {
    border: 0;
}

/**
 * Correct overflow displayed oddly in IE 9.
 */

svg:not(:root) {
    overflow: hidden;
}

/* ==========================================================================
   Figures
   ========================================================================== */

/**
 * Address margin not present in IE 8/9 and Safari 5.
 */

figure {
    margin: 0;
}

/* ==========================================================================
   Forms
   ========================================================================== */

/**
 * Define consistent border, margin, and padding.
 */

fieldset {
    border: 1px solid #c0c0c0;
    margin: 0 2px;
    padding: 0.35em 0.625em 0.75em;
}

/**
 * 1. Correct color not being inherited in IE 8/9.
 * 2. Remove padding so people aren't caught out if they zero out fieldsets.
 */

legend {
    border: 0; /* 1 */
    padding: 0; /* 2 */
}

/**
 * 1. Correct font family not being inherited in all browsers.
 * 2. Correct font size not being inherited in all browsers.
 * 3. Address margins set differently in Firefox 4+, Safari 5, and Chrome.
 */

button,
input,
select,
textarea {
    font-family: inherit; /* 1 */
    font-size: 100%; /* 2 */
    margin: 0; /* 3 */
}

/**
 * Address Firefox 4+ setting line-height on input using !important in
 * the UA stylesheet.
 */

button,
input {
    line-height: normal;
}

/**
 * Address inconsistent text-transform inheritance for button and select.
 * All other form control elements do not inherit text-transform values.
 * Correct button style inheritance in Chrome, Safari 5+, and IE 8+.
 * Correct select style inheritance in Firefox 4+ and Opera.
 */

button,
select {
    text-transform: none;
}

/**
 * 1. Avoid the WebKit bug in Android 4.0.* where (2) destroys native audio
 *    and video controls.
 * 2. Correct inability to style clickable input types in iOS.
 * 3. Improve usability and consistency of cursor style between image-type
 *    input and others.
 */

button,
html input[type="button"], /* 1 */
input[type="reset"],
input[type="submit"] {
    -webkit-appearance: button; /* 2 */
    cursor: pointer; /* 3 */
}

/**
 * Re-set default cursor for disabled elements.
 */

button[disabled],
html input[disabled] {
    cursor: default;
}

/**
 * 1. Address box sizing set to content-box in IE 8/9.
 * 2. Remove excess padding in IE 8/9.
 */

input[type="checkbox"],
input[type="radio"] {
    box-sizing: border-box; /* 1 */
    padding: 0; /* 2 */
}

/**
 * 1. Address appearance set to searchfield in Safari 5 and Chrome.
 * 2. Address box-sizing set to border-box in Safari 5 and Chrome
 *    (include -moz to future-proof).
 */

input[type="search"] {
    -webkit-appearance: textfield; /* 1 */
    -moz-box-sizing: content-box;
    -webkit-box-sizing: content-box; /* 2 */
    box-sizing: content-box;
}

/**
 * Remove inner padding and search cancel button in Safari 5 and Chrome
 * on OS X.
 */

input[type="search"]::-webkit-search-cancel-button,
input[type="search"]::-webkit-search-decoration {
    -webkit-appearance: none;
}

/**
 * Remove inner padding and border in Firefox 4+.
 */

button::-moz-focus-inner,
input::-moz-focus-inner {
    border: 0;
    padding: 0;
}

/**
 * 1. Remove default vertical scrollbar in IE 8/9.
 * 2. Improve readability and alignment in all browsers.
 */

textarea {
    overflow: auto; /* 1 */
    vertical-align: top; /* 2 */
}

/* ==========================================================================
   Tables
   ========================================================================== */

/**
 * Remove most spacing between table cells.
 */

table {
    border-collapse: collapse;
    border-spacing: 0;
}`
)
