# README

## About

This is an searching app which can search and return the occurence of input word.

## Development

- GUI:
Wails framework enables React frontend integration.

- Main Logic:
Implented by Golang. Using gse package for terms indexing instead of marking every words.

## Building

To build a redistributable, production mode package, use `wails build`, or if webkitgtk isn't v4.0, run `wails build -tags webkit2_41` instead.
