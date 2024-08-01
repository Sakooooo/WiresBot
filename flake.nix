{
  description = "wiresbot go thing";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default-linux";
  };

  outputs = {
    self,
    nixpkgs,
    systems,
  }: let
    inherit (nixpkgs) lib;
    eachSystem = lib.genAttrs (import systems);
    pkgsFor = eachSystem (system: import nixpkgs {localSystem = system;});

    inherit (builtins) concatStringsSep match;

    date = concatStringsSep "-" (match "(.{4})(.{2})(.{2}).*" self.lastModifiedDate);
  in {
    packages = eachSystem (system: {
      default = self.packages.${system}.wiresbot-go;
      wiresbot-go = pkgsFor.${system}.callPackage ./nix/package.nix {inherit date;};
    });

    devShells = eachSystem (system: {
      default = pkgsFor.${system}.mkShell {
        name = "wiresbot-shell";
        packages = with pkgsFor.${system}; [go gopls gotools go-tools imagemagick.dev];
      };
    });

    #
  };
}
