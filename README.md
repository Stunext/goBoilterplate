<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>

[![Sourcegraph](https://sourcegraph.com/github.com/labstack/echo/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/labstack/echo?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/labstack/echo)
[![Go Report Card](https://goreportcard.com/badge/github.com/labstack/echo?style=flat-square)](https://goreportcard.com/report/github.com/labstack/echo)
[![Join the chat at https://gitter.im/labstack/echo](https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square)](https://gitter.im/labstack/echo)
[![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://forum.labstack.com)
[![Twitter](https://img.shields.io/badge/twitter-@labstack-55acee.svg?style=flat-square)](https://twitter.com/labstack)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)

## Feature Overview

-   Optimized HTTP router which smartly prioritize routes
-   Build robust and scalable RESTful APIs
-   Group APIs
-   Extensible middleware framework
-   Define middleware at root, group or route level
-   Data binding for JSON, XML and form payload
-   Handy functions to send variety of HTTP responses
-   Centralized HTTP error handling
-   Template rendering with any template engine
-   Define your format for the logger
-   Highly customizable
-   Automatic TLS via Let’s Encrypt
-   HTTP/2 support

## How to run

```sh
$ go run main.go
```

## About GORM

The fantastic ORM library for Golang, aims to be developer friendly.

[![go report card](https://goreportcard.com/badge/github.com/jinzhu/gorm "go report card")](https://goreportcard.com/report/github.com/jinzhu/gorm)
[![wercker status](https://app.wercker.com/status/8596cace912c9947dd9c8542ecc8cb8b/s/master "wercker status")](https://app.wercker.com/project/byKey/8596cace912c9947dd9c8542ecc8cb8b)
[![Join the chat at https://gitter.im/jinzhu/gorm](https://img.shields.io/gitter/room/jinzhu/gorm.svg)](https://gitter.im/jinzhu/gorm?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/jinzhu/gorm?status.svg)](https://godoc.org/github.com/jinzhu/gorm)

## Overview

-   Full-Featured ORM (almost)
-   Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
-   Hooks (Before/After Create/Save/Update/Delete/Find)
-   Preloading (eager loading)
-   Transactions
-   Composite Primary Key
-   SQL Builder
-   Auto Migrations
-   Logger
-   Extendable, write Plugins based on GORM callbacks
-   Every feature comes with tests
-   Developer Friendly

## Other dependencies (Not Tracket)

-   [Testify/Assert](https://github.com/stretchr/testify/assert).

Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend.

Install:

```sh
$ go get -u github.com/stretchr/testify/assert
```

Use:

```sh
$ go test ./tests
```

-   [Swagger](https://github.com/swaggo/swag).

Swag converts Go annotations to Swagger Documentation 2.0. And provides a variety of builtin web framework lib. Let you can quickly integrated in existing golang project(using Swagger UI).

Install:

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```

Make sure you have a GOBIN system enviroument variable.

Use:

```sh
$ swag init
```

Especification in: https://swaggo.github.io/swaggo.io/
