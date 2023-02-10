let lock = builtins.fromJSON (builtins.readFile ./flake.lock); in

{ pkgs ? import (builtins.fetchTarball "https://github.com/NixOS/nixpkgs/archive/${lock.nodes.nixpkgs.locked.rev}.tar.gz") { } }:

with pkgs;

buildGoModule {
  name = "pinentry-mac-keychain";
  src = ./.;
  patches = [
    (pkgs.substituteAll {
      src = ./pinentry-mac-location.patch;
      pinentryMac = "${pinentry_mac}/Applications/pinentry-mac.app/Contents/MacOS/pinentry-mac";
    })
  ];
  vendorSha256 = "sha256-7UzdYGhi9asQRVb9EbaW0ijXf+lDnkM01Pv6yGsAghM=";

  postInstall = ''
    echo "Add the following line to your ~/.gnupg/gpg-agent.conf file:"
    echo "  pinentry-program $out/bin/pinentry-mac-keychain"
  '';

  meta = with lib; {
    description = "Pinentry with macOs keychain support";
    longDescription = builtins.readFile ./README.md;
    license = licenses.mit;
    platforms = platforms.darwin;
    maintainers = with maintainers; [ olebedev ];
  };
}
