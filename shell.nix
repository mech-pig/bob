let
  nixpkgs = builtins.fetchGit {
    name = "nixos-unstable-2023-05-30";
    url = "https://github.com/NixOS/nixpkgs/";
    ref = "refs/heads/nixos-unstable";
    rev = "5e871d8aa6f57cc8e0dc087d1c5013f6e212b4ce";
  };

  pkgs = import nixpkgs {};
in
  with pkgs;
    mkShell {
      buildInputs = [
        go
        gotools
        go-tools
        gopls
        go-outline
        gocode
        gopkgs
        gocode-gomod
        godef
        golint
        delve
      ];
    }