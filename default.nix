{ lib
, buildGoApplication
, buildPackages
, rev ? "dirty"
}:
let
  version = "latest";
  pname = "furyd";
  tags = [ "netgo" ];
  ldflags = lib.concatStringsSep "\n" ([
    "-X github.com/cosmos/cosmos-sdk/version.Name=fury"
    "-X github.com/cosmos/cosmos-sdk/version.AppName=${pname}"
    "-X github.com/cosmos/cosmos-sdk/version.Version=${version}"
    "-X github.com/cosmos/cosmos-sdk/version.BuildTags=${lib.concatStringsSep "," tags}"
    "-X github.com/cosmos/cosmos-sdk/version.Commit=${rev}"
  ]);
in
buildGoApplication rec {
  inherit pname version tags ldflags;
  go = buildPackages.go_1_20;
  src = ./.;
  modules = ./gomod2nix.toml;
  doCheck = false;
  pwd = src; # needed to support replace
  subPackages = [ "cmd/furyd" ];
  CGO_ENABLED = "1";

  meta = with lib; {
    description = "Fury is a scalable and interoperable blockchain, built on Proof-of-Stake with fast-finality using the Cosmos SDK which runs on top of CometBFT Core consensus engine.";
    homepage = "https://github.com/exfury/fury";
    license = licenses.asl20;
    mainProgram = "furyd";
  };
}
