{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  outputs = { self, nixpkgs, flake-utils }: flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = nixpkgs.legacyPackages.${system};
      pinentry-mac-keychain = pkgs.callPackage ./default.nix { };
    in
    rec {
      packages.default = pinentry-mac-keychain;
      packages.pinentry-mac-keychain = pinentry-mac-keychain;
      devShells.default = pkgs.callPackage ./shell.nix {
        inherit pinentry-mac-keychain;
      };
    }
  );
}
