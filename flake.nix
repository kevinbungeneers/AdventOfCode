{
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils }:
        flake-utils.lib.eachDefaultSystem (system:
            let
                pkgs = nixpkgs.legacyPackages.${system};

                go_setup = ''
                    export GOPATH=$PWD/.nix/go
                    mkdir -p $GOPATH
                    export PATH=$GOPATH/bin:$PATH
                '';
            in {
                formatter = pkgs.nixpkgs-fmt;

                devShells.default = pkgs.mkShell {
                    buildInputs = with pkgs; [
                        go_1_19
                    ];

                    shellHook = ''
                        ${go_setup}
                    '';
                };
            }
        );
}
