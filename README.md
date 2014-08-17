Acceptance Tests for Staticfile Buildpack
=========================================

This test suite exercises various example applications against a pre-installed https://github.com/cloudfoundry-community/staticfile-buildpack.

This project is based on the CATS test project. A Cloud Foundry is targeted the same way, and the tests are run the same way.

Prerequisites
-------------

Prior to running this test suite, the latest [staticfile-buildpack](https://github.com/cloudfoundry-community/staticfile-buildpack) release must be imported into your Cloud Foundry by an admin user.

See [Administrator Upload](https://github.com/cloudfoundry-community/staticfile-buildpack#administrator-upload) section of Buildpack readme.

```bash
$ cf buildpacks
Getting buildpacks...

buildpack               position   enabled   locked   filename
staticfiles_buildpack   1          true      false    staticfile-buildpack-v0.4.1.zip
java_buildpack          2          true      false    java-buildpack-v2.4.zip
...
```

Running the tests
-----------------

See [CATS readme](https://github.com/cloudfoundry/cf-acceptance-tests#running-the-tests) for how to run tests.

The following script will configure these prerequisites for a [bosh-lite](https://github.com/cloudfoundry/bosh-lite) installation. Replace credentials and URLs as appropriate for your environment.

```bash
#! /bin/bash

cat > integration_config.json <<EOF
{
  "api": "api.10.244.0.34.xip.io",
  "admin_user": "admin",
  "admin_password": "admin",
  "apps_domain": "10.244.0.34.xip.io"
}
EOF
```

To run the tests:

```
$ export CONFIG=$PWD/integration_config.json
$ ./bin/test -v
...
Ran 5 of 5 Specs in 136.036 seconds
SUCCESS! -- 5 Passed | 0 Failed | 0 Pending | 0 Skipped PASS
```

Changing CATs
-------------

### Dependency Management

These Acceptance Tests use [godep](https://github.com/tools/godep) to manage `go` dependencies.

All `go` packages required to run CATs are vendored into the `cf-acceptance-tests/Godeps` directory.

When making changes to the test suite that bring in additional `go` packages, you should use the workflow described in the[Add or Update a Dependency](https://github.com/tools/godep#add-or-update-a-dependency) section of the godep documentation.

As the project has no `.go` files in the root folder, to save any changes to dependencies:

```bash
godep save ./...
```
