{ pkgs }:
pkgs.mkShell {
  name = "pinentry-touchid";

  # Add build dependencies
  packages = with pkgs; [
    go
    gopls
    gotools
  ];

  # Add environment variables
  env = { };

  # Load custom bash code
  shellHook = ''
    unset GOPATH GOROOT
    export NIX_LDFLAGS="-framework CoreFoundation -framework LocalAuthentication -framework Foundation $NIX_LDFLAGS";
  '';
}
