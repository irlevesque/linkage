# Linkage

![Linkage Logo](assets/linkage.png)

## Project

A self-hosted search engine that redirects to specific sites (similar to go/links).

## Features:

- Redirects to specific sites if the "search term" is available
- Redirects searches to a fallback search engine if no links are found

## Planned Features

- Appends any additionally-supplied parameters to the redirected URL
- Token-based authenticated REST API for link management and fallback search engine behavior

## Example Usage

- `gh`: redirects to https://github.com
- `gh/irlevesque`: redirects to https://github.com/irlevesque

## Setup

On Chrome (and deriviative browsers), go to `settings` > `search engine` and add a new "Site search" engine with the name "Linkage" and the URL "http://localhost:8080/?s=%s"
