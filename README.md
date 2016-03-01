# Rabbitbeat

Welcome to Rabbitbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/cero-t`

## Getting Started with Rabbitbeat

### Init Project
To get running with Rabbitbeat, run the following commands:

```
glide update --no-recursive
make update
```


To push Rabbitbeat in the git repository, run the following commands:

```
git init
git add .
git commit
git remote set-url origin https://github.com/cero-t/rabbitbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Rabbitbeat run the command below. This will generate a binary
in the same directory with the name rabbitbeat.

```
make
```


### Run

To run Rabbitbeat with debugging output enabled, run:

```
./rabbitbeat -c rabbitbeat.yml -e -d "*"
```


### Test

To test Rabbitbeat, run the following commands:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/rabbitbeat.template.json and etc/rabbitbeat.asciidoc

```
make update
```


### Cleanup

To clean  Rabbitbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Rabbitbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/cero-t
cd ${GOPATH}/github.com/cero-t
git clone https://github.com/cero-t/rabbitbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
