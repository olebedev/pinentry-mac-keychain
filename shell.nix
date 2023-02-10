let lock = builtins.fromJSON (builtins.readFile ./flake.lock); in

{ pkgs ? import (builtins.fetchTarball "https://github.com/NixOS/nixpkgs/archive/${lock.nodes.nixpkgs.locked.rev}.tar.gz") { }
, pinentry-mac-keychain ? pkgs.callPackage ./default.nix { }
}:

pkgs.mkShell {
  buildInputs = [
    pinentry-mac-keychain
  ];
}
