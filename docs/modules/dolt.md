# Dolt

Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

## Introduction

The Testcontainers module for Dolt.

## Adding this module to your project dependencies

Please run the following command to add the Dolt module to your Go dependencies:

```
go get github.com/testcontainers/testcontainers-go/modules/dolt
```

## Usage example

<!--codeinclude-->
[Creating a Dolt container](../../modules/dolt/examples_test.go) inside_block:runDoltContainer
<!--/codeinclude-->

## Module Reference

### Run function

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.32.0"><span class="tc-version">:material-tag: v0.32.0</span></a>

!!!info
    The `RunContainer(ctx, opts...)` function is deprecated and will be removed in the next major release of _Testcontainers for Go_.

The Dolt module exposes one entrypoint function to create the Dolt container, and this function receives three parameters:

```golang
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*DoltContainer, error)
```

- `context.Context`, the Go context.
- `string`, the Docker image to use.
- `testcontainers.ContainerCustomizer`, a variadic argument for passing options.

#### Image

Use the second argument in the `Run` function to set a valid Docker image.
In example: `Run(context.Background(), "dolthub/dolt-sql-server:1.32.4")`.

### Container Options

When starting the Dolt container, you can pass options in a variadic way to configure it.

#### Set username, password and database name

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

If you need to set a different database, and its credentials, you can use `WithUsername`, `WithPassword`, `WithDatabase`
options.

!!!info
The default values for the username is `root`, for password is `test` and for the default database name is `test`.

#### WithScripts

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

If you would like to perform DDL or DML operations in the Dolt container, add one or more `*.sql`, `*.sql.gz`, or `*.sh`
scripts to the container request, using the `WithScripts(scriptPaths ...string)`. Those files will be copied under `/docker-entrypoint-initdb.d`.

#### WithDoltCloneRemoteUrl

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

If you would like to clone data from a remote into the Dolt container, add an `*.sh`
scripts to the container request, using the `WithScripts(scriptPaths ...string)`. Additionally, use `WithDoltCloneRemoteUrl(url string)` to specify
the remote to clone, and use `WithDoltCredsPublicKey(key string)` along with `WithCredsFile(credsFile string)` to authorize the Dolt container to clone from the remote.

<!--codeinclude-->
[Example of Clone script](../../modules/dolt/testdata/clone-db.sh)
<!--/codeinclude-->

#### WithConfigFile

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

If you need to set a custom configuration, you can use `WithConfigFile` option to pass the path to a custom configuration file.

{% include "../features/common_functional_options_list.md" %}

### Container Methods

#### ConnectionString

- Since <a href="https://github.com/testcontainers/testcontainers-go/releases/tag/v0.31.0"><span class="tc-version">:material-tag: v0.31.0</span></a>

This method returns the connection string to connect to the Dolt container, using the default `3306` port.
It's possible to pass extra parameters to the connection string, e.g. `tls=skip-verify` or `application_name=myapp`, in a variadic way.

<!--codeinclude-->
[Get connection string](../../modules/dolt/dolt_test.go) inside_block:connectionString
<!--/codeinclude-->
