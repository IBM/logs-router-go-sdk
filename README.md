[![Build Status](https://travis-ci.com/IBM/logs-router-go-sdk.svg?token=eW5FVD71iyte6tTby8gr&branch=main)](https://travis-ci.com/IBM/logs-router-go-sdk.svg?token=eW5FVD71iyte6tTby8gr&branch=main)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Logs Routing Go SDK 1.1.0
Go client library to interact with the various [IBM Cloud Logs Router APIs](https://cloud.ibm.com/apidocs/logs-router-service-api/logs-router-v1).

> [!WARNING]
> As of 28 March 2024 the logdna service is deprecated and will no longer be supported as of 30 March 2025.
> The IBM Cloud Logs Routing Go SDK will stop supporting `logdna` targets at the same time and no logs will be routed to these type of targets after that date.
> You should make sure that you have configured your tenant to direct your logs to another destination before 30 March 2025.
> Any `logdna` targets still configured after 30 April 2025 will be removed automatically from your IBM Cloud Logs Routing configuration.


## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Go modules](#go-modules)
    - [`go get` command](#go-get-command)
  - [Using the SDK](#using-the-sdk)
  - [Questions](#questions)
  - [Issues](#issues)
  - [Open source @ IBM](#open-source--ibm)
  - [Contributing](#contributing)
  - [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Logs Router Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Logs Routing](https://cloud.ibm.com/apidocs/logs-router-service-api/logs-router-v1) | ibmcloudlogsroutingv0

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.19 or above.

## Installation
The current version of this SDK: 1.0.5

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `ibmcloudlogsroutingv0` part of the import path is the package name
associated with IBM Cloud Logs Routing.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](github.com/IBM/logs-router-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
