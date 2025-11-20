# Linkage

A self-hosted replacement search engine that redirects Linkage Terms to specific addresses (inspired by [golinks](https://meta.wikimedia.org/wiki/Go_links)). Unmatched terms are passed through to your preferred search engine.

## Features:

- Redirects you to specific sites if the "search term" is a Linkage term.
- Appends any additionally-supplied parameters to the redirected linkage URL.
- Redirects searches to a user-defined fallback search engine if no links are found.
- No analytics or tracking. Everything is in your control.

### Planned Features

- Web-based interface for link management and fallback search engine behavior
- Preserve URL query parameters

### Currently Unplanned Features

- Support for HTTPS

## Example Usage

If you configure the Linkage search engine to be the default search engine, simply type a Linkage term or search term into the address bar.
Otherwise, you can type `go` followed by a space and the linkage / search term.

The default configuration includes the `gh` linkage which redirects to github. Therefore, you can test the "Search engine" configuration via:
- _(go)_ `gh`: redirects to https://github.com
- _(go)_ `gh/irlevesque`: redirects to https://github.com/irlevesque

## Setup

### Chrome (and deriviative browsers):

Go to `Settings` > `Search Engine` and add a new `Site search` engine:
- Name: "Linkage"
- Shortcut: "go"
- URL: "http://localhost:8080/?s=%s" (assuming you're running locally on the default port).
- Optionally, after saving the search engine, you can make it default using the triple-dot dropdown menu.
