{ pkgs ? import ../../../nix { } }:
let furyd = (pkgs.callPackage ../../../. { });
in
furyd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-furyd.patch
  ];
})
