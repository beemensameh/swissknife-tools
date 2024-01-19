# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  channel = "stable-23.11"; # "stable-23.11" or "unstable"
  # Use https://search.nixos.org/packages to  find packages
  packages = [
    pkgs.go
  ];
  # Sets environment variables in the workspace
  env = {};
  # search for the extension on https://open-vsx.org/ and use "publisher.id"
  idx.extensions = [
    "golang.go"
  ];
  idx.workspace = {
    # runs when a workspace is first created with this `dev.nix` file
    # to run something each time the environment is rebuilt, use the `onStart` hook
    onCreate = {
      install-go-tools = ''
        go run github.com/golang/vscode-go/tools/installtools@latest
      '';
      };
  };
  # preview configuration, identical to monospace.json
  # idx.previews = {
  #   enable = true;
  #   previews = [
  #     {
  #       command = ["go" "run" "server.go" "-addr" "localhost:$PORT"];
  #       manager = "web";
  #       id = "web";
  #     }
  #     {
  #       manager = "ios";
  #       id = "ios";
  #     }
  #   ];
  # };
}