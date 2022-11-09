{ pkgs }: {
    deps = [
        pkgs.dirb
        pkgs.openssh_with_kerberos
        pkgs.go_1_17
        pkgs.gopls
    ];
}