{ pkgs ? import (builtins.fetchGit {
        # Descriptive name to make the store path easier to identify                
        name = "dev-go";                        
        url = "https://github.com/NixOS/nixpkgs";                       
         ref = "refs/heads/master";                     
         rev = "b31e87640a5553fbd972e5061d02b067412317a0"; 
}) {} }:

with pkgs;

mkShell {
  buildInputs = [
    gitlint
    go_1_22
    go-tools
    golangci-lint
    goreleaser
    gosec
    gotools
    gofumpt
    golint
    pre-commit
    awscli2
    act
  ];

  shellHook =
    ''
    # Setup the binaries installed via `go install` to be accessible globally.
      export PATH="$(go env GOPATH)/bin:$PATH"
      export GOPROXY="https://proxy.golang.org,direct"
      export GOSUMDB="sum.golang.org"
    # Install pre-commit hooks.
      pre-commit install

    # Install Go binaries.
      which protoc-gen-grpc-gateway || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
      which protoc-gen-openapiv2 || go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
      which gocritic || go install github.com/go-critic/go-critic/cmd/gocritic@latest
      which goreturns || go install github.com/sqs/goreturns@latest
      which swag || go get -u github.com/swaggo/swag
      which mockgen || go install github.com/golang/mock/mockgen@v1.6.0
      
    # Clear the terminal screen.
      clear
    '';
}
