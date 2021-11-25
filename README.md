# Webflow

## Webflow 

This is a tool for performing a series of web tasks, based on the chromedp(https://github.com/chromedp/chromedp) golang version.

## Usage

You can run the following command to view full help:

```sh
webflow v0.5.1

Usage:

  click <selector>    Click element by <selector>, like document.querySelector
  debug <true|false>    Show chromedp debug verbose log
  eval [$variable] <{Javascript Code}>  Evaluate Javascript code and save result to $variable
  flag <string>  <string|bool>  Custom browser flags.
  focus <selector>    Focus element by <selector>, like document.querySelector
  getcookie $variable [domain] [name]   Get all cookies, by [domain] or by [name], and save to $variable
  headless <true|false>   Set headless flag to chromedp, default is true means don't show browser window
  import <script|javascript_file>   Import <script|javascript_file> to evaluate on new document created
  keys <selector> <keys|$variable>  Type <keys> or $variable to element <selector>, use chromedp.SendKeys method
  loadcookie <@file> [domain]   Load cookies saved in <@file> with <domain>
  log <true|false>  Set enable/disable flow verbose log.
  open <url>  Open <url> in browser
  poll <expressions>  [milliseconds]  Poll result of <expressions>, timeout with <milliseconds> is optional
  print <string|$variable>  Print out <string> or $variable value
  printf <format> <field1> <$variable> ...  Print format string, bridge to fmt.Printf function
  save <@file> <string|$variable>   Save strings or $variable value to <@file>
  screen <width>x<height>   Setup chrome window size, eg: screen 1080x720
  screenshot <selector> <path>  [interval]  Take screenshot of <selector> and save to <path>, [interval] is optional
  setcookie <name> <value> <domain>   Set cookie key/value to browser context with <domain>
  setupload <selector> <file1> <file2> ...  Set files of element <selector> for upload
  setvalue <selector> <string|$variable>  Set the value of element <selector> with <string> or $variable
  shell <script|$variable>  Run custom shell scripts, support $variable
  switch <URL>  Switch context which URL contains <URL>.
  timeout <milliseconds>  Setup timeout duration, by milliseconds
  useragent <string>  Setup browser UserAgent string
  userdata <string>   Setup browser userdata storage path.
  var $variable <string>  Define $variable with value <string>
  wait <selector|milliseconds>  Wait for target with <selector> OR timeout with <milliseconds>

  [..]  Parameter is optional;
  <..>  Parameter is required;
  ... Follow with one or more parameters;

Options:

  -f, --file      Specify Flowfile path.
  -c, --flow      Using raw flow content string.
  -i, --insert    Insert new flow before the flow content.
  -a, --append    Append new flow to the end of flow content.
  -v, --version   Print webflow version.
  -d, --verbose   Verbose detail mode.
  -h, --help      Show help.
```

### Simple Examples

1. Open www.baidu.com and print 'It works':

```sh
webflow -c 'headless false; open "https://www.baidu.com"; print "It works!";'
```

2. Open www.baidu.com and search 'hello':

```sh
webflow -c 'headless false; open "https://www.baidu.com/"; wait "#kw"; keys "#kw" "hello"; click "#su"; wait 10000;'
```

3. Open www.baidu.com and print web content:

```sh
webflow -c 'headless false; open "https://www.baidu.com/"; wait "#kw"; eval $output "document.body.textContent"; print $output;'
```

## License

MIT

