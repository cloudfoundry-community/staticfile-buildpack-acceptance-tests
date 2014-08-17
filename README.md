Acceptance Tests for Staticfiles Buildpack
==========================================

This test suite exercises various example applications against a pre-installed https://github.com/cloudfoundry-community/staticfile-buildpack.

This project is based on the CATS test project. A Cloud Foundry is targeted the same way, and the tests are run the same way.

Running the tests
-----------------

See [CATS readme](https://github.com/cloudfoundry/cf-acceptance-tests#running-the-tests) for how to run tests.

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
