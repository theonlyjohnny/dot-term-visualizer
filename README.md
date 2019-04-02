
# Dot Term Visualizer

## Description
Dot Term Visualizer is a Golang CLI tool that renders [Graphviz](http://graphviz.gitlab.io/) dotfiles directly in the terminal. Useful for debugging [Python memory leaks](https://mg.pov.lt/objgraph/), debugging [graphics pipelines](https://gstreamer.freedesktop.org/documentation/tutorials/basic/debugging-tools.html#getting-pipeline-graphs) and any other Graphviz usecase.
  
  Unlike other Graphviz visualizers, this tool renders direct to stdout. I do much of my work via remote SSH connection, and I got sick of generating dotfiles on a remote host, converting them to images, then SCPing them to my local system to view them.

### TODO (PRs welcome!):
 - [ ] Rename this project
 - [ ] Properly render multiple dotfiles
 - [ ] Fix scaling/render math in rect.go
 - [ ] Redirect log output to file so the tscreen doesn't take it over
 - [ ] Scrolling
 - [ ] [More shapes!](https://graphviz.gitlab.io/_pages/doc/info/attrs.html)
