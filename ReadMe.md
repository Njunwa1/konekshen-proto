# gRPC Project Setup and Deployment Guide

## Protobuf Code Generation

### Generate Go Code with protoc

To generate Go code from your `.proto` files, use the following protoc command:

```bash
protoc --go_out=./golang --go_opt=paths=source_relative \
       --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=./golang --grpc-gateway_opt=paths=source_relative \
       --grpc-gateway_opt=generate_unbound_methods=true \
       ./**/*.proto
```

#### Command Breakdown:
- `--go_out=./golang`: Generate Go structs
- `--go_opt=paths=source_relative`: Use relative import paths
- `--go-grpc_out=./golang`: Generate gRPC service code
- `--grpc-gateway_out=./golang`: Generate REST gateway code
- `./**/*.proto`: Recursively find all .proto files

## Git Workflow

### Committing Changes

```bash
# Stage all changes
git add .

# Commit with a descriptive message
git commit -m "feat(users): add user service implementation"

# Push to the current branch
git push origin <branch-name>
```

### Creating and Pushing a Release Tag

```bash
# Create a new tag for a specific module
git tag golang/users/v1.1.16

# Push the tag to remote repository
git push origin golang/users/v1.1.16

# To push all tags
git push --tags
```

### Creating a GitHub Release

1. Go to your GitHub repository
2. Click on "Releases"
3. Click "Draft a new release"
4. Choose the tag you just created
5. Add release notes
6. Publish the release

## Best Practices

- Use semantic versioning (MAJOR.MINOR.PATCH)
- Include a changelog with each release
- Ensure all tests pass before creating a tag
- Use descriptive commit messages

## Prerequisites

- protoc (Protobuf Compiler)
- Go
- gRPC tools
- grpc-gateway

## Troubleshooting

- Ensure all dependencies are installed
- Check protoc version compatibility
- Verify Go module configuration
```