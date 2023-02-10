# pinentry-mac-keychain

Custom GPG pinentry program for macOS that fetching the password from the macOS keychain. _WARNING: Use it on your own risk._                               

### How does it work

The `pinentry-mac-keychain` acts like a proxy between the `gpg-agent` and the original `pinentry-mac` 
executable in order to hijack requests for PIN and pull the PIN from the users keychain instead of
poping up the `pinentry-mac` UI each time. This is useful if you can be sure no one will attempt
to use your macOs machine to perform malicuous actions while you away but your Smart Card, like Yubikey,
is plugged in. Otherwise it will be a significant security risk.

The program proxies only the requests that can be recongnised as the requirests for the PIN, coming from the GnuPG toolset (see the RegExp that detects that in [`./utils.go#L13`](./utils.go#L13)) and doesn't proxies any other requirests, like `MESSAGE` and `CONFIRM`.

### Install

Installation is being performed via Nix package manager. So you'd need to have it installed first.
Once Nix is installed with Flakes support, you can install `pinentry-mac-keychain` via:

```
$ nix profile install github:olebedev/pinentry-mac-keychain
```

Otherwise, if your Nix installation doesn't support Flakes you can install it via:
```
$ nix-env -i -f https://github.com/olebedev/pinentry-mac-keychain/archive/main.tar.gz
```

After that, make sure you added the following line 
```
pinentry-program /Users/<your user name>/.nix-profile/bin/pinentry-mac-keychain
```
to your `~/.gnupg/gpg-agent.conf` file and restarted GPG agent via:
```
gpgconf --kill all
```

### Build from source

```
go install .
```

### Debug

It can be built so it writes down logs into a file:
```
go install -ldflags "-X main.logfile=$HOME/pinentry.log"
```

### TODO:
  - [ ] setup github actions

