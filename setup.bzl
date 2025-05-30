load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def github_archive(name, repo_name, commit, kind = "zip", strip_prefix = "", **kwargs):
    """Defines a GitHub commit-based repository rule."""
    project = repo_name[repo_name.index("/"):]
    if "sha256" in kwargs:
        print("Ignoring unstable GitHub sha256 argument in", name)
        kwargs.pop("sha256")
    http_archive(
        name = name,
        strip_prefix = "{project}-{commit}/{prefix}".format(
            project = project,
            commit = commit,
            prefix = strip_prefix,
        ),
        urls = [u.format(commit = commit, repo_name = repo_name, kind = kind) for u in [
            "https://mirror.bazel.build/github.com/{repo_name}/archive/{commit}.{kind}",
            "https://github.com/{repo_name}/archive/{commit}.{kind}",
        ]],
        canonical_id = commit,
        **kwargs
    )

def kythe_rule_repositories():
    """Defines external repositories for Kythe Bazel rules.

    These repositories must be loaded before calling external.bzl%kythe_dependencies.
    """
    maybe(
        http_archive,
        name = "platforms",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/platforms/releases/download/0.0.8/platforms-0.0.8.tar.gz",
            "https://github.com/bazelbuild/platforms/releases/download/0.0.8/platforms-0.0.8.tar.gz",
        ],
        sha256 = "8150406605389ececb6da07cbcb509d5637a3ab9a24bc69b1101531367d89d74",
    )

    maybe(
        http_archive,
        name = "bazel_skylib",
        urls = [
            "https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
        ],
        sha256 = "66ffd9315665bfaafc96b52278f57c7e2dd09f5ede279ea6d39b2be471e7e3aa",
    )

    maybe(
        http_archive,
        name = "io_bazel_rules_go",
        integrity = "sha256-yruF04LPIjMYaJjb6XivkvXxyZQv/R0Y6mob/2jyNCo=",
        strip_prefix = "rules_go-077f15fe11b9da6aa0e3271db1260929f04fef87",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/archive/077f15fe11b9da6aa0e3271db1260929f04fef87.zip",
            "https://github.com/bazelbuild/rules_go/archive/077f15fe11b9da6aa0e3271db1260929f04fef87.zip",
        ],
    )

    maybe(
        http_archive,
        name = "rules_cc",
        urls = ["https://github.com/bazelbuild/rules_cc/releases/download/0.0.9/rules_cc-0.0.9.tar.gz"],
        sha256 = "2037875b9a4456dce4a79d112a8ae885bbc4aad968e6587dca6e64f3a0900cdf",
        strip_prefix = "rules_cc-0.0.9",
    )

    maybe(
        http_archive,
        name = "rules_java",
        urls = [
            # Note: when updating rules_java, please check if the ModuleName check in tools/javacopts.bazelrc can be re-enabled.
            "https://mirror.bazel.build/github.com/bazelbuild/rules_java/releases/download/6.5.2/rules_java-6.5.2.tar.gz",
            "https://github.com/bazelbuild/rules_java/releases/download/6.5.2/rules_java-6.5.2.tar.gz",
        ],
        sha256 = "16bc94b1a3c64f2c36ceecddc9e09a643e80937076b97e934b96a8f715ed1eaa",
    )

    maybe(
        http_archive,
        name = "rules_proto",
        sha256 = "71fdbed00a0709521ad212058c60d13997b922a5d01dbfd997f0d57d689e7b67",
        strip_prefix = "rules_proto-6.0.0-rc2",
        url = "https://github.com/bazelbuild/rules_proto/releases/download/6.0.0-rc2/rules_proto-6.0.0-rc2.tar.gz",
    )

    maybe(
        http_archive,
        name = "bazel_gazelle",
        sha256 = "5d80e62a70314f39cc764c1c3eaa800c5936c9f1ea91625006227ce4d20cd086",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.42.0/bazel-gazelle-v0.42.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.42.0/bazel-gazelle-v0.42.0.tar.gz",
        ],
    )

    maybe(
        http_archive,
        name = "aspect_rules_js",
        sha256 = "d66f8abf914a0454a69181b7b17acaae56d7b0e2784cb26b40cb3273c4d836d1",
        strip_prefix = "rules_js-2.2.0",
        url = "https://github.com/aspect-build/rules_js/releases/download/v2.2.0/rules_js-v2.2.0.tar.gz",
    )

    maybe(
        http_archive,
        name = "rules_jvm_external",
        sha256 = "f86fd42a809e1871ca0aabe89db0d440451219c3ce46c58da240c7dcdc00125f",
        strip_prefix = "rules_jvm_external-5.2",
        urls = ["https://github.com/bazelbuild/rules_jvm_external/releases/download/5.2/rules_jvm_external-5.2.tar.gz"],
    )

    maybe(
        http_archive,
        name = "aspect_rules_ts",
        sha256 = "4263532b2fb4d16f309d80e3597191a1cb2fb69c19e95d91711bd6b97874705e",
        strip_prefix = "rules_ts-3.5.0",
        url = "https://github.com/aspect-build/rules_ts/releases/download/v3.5.0/rules_ts-v3.5.0.tar.gz",
    )

    maybe(
        http_archive,
        name = "aspect_rules_jasmine",
        sha256 = "0d2f9c977842685895020cac721d8cc4f1b37aae15af46128cf619741dc61529",
        strip_prefix = "rules_jasmine-2.0.0",
        url = "https://github.com/aspect-build/rules_jasmine/releases/download/v2.0.0/rules_jasmine-v2.0.0.tar.gz",
    )

    maybe(
        http_archive,
        name = "rules_python",
        sha256 = "e5470e92a18aa51830db99a4d9c492cc613761d5bdb7131c04bd92b9834380f6",
        strip_prefix = "rules_python-4b84ad270387a7c439ebdccfd530e2339601ef27",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
            "https://github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
        ],
    )

    maybe(
        http_archive,
        name = "rules_rust",
        sha256 = "db89135f4d1eaa047b9f5518ba4037284b43fc87386d08c1d1fe91708e3730ae",
        urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.27.0/rules_rust-v0.27.0.tar.gz"],
    )

    maybe(
        http_archive,
        name = "bazelruby_rules_ruby",
        strip_prefix = "rules_ruby-0.4.1",
        sha256 = "abfc2758cc379e0ff0eb9824e3b507c1633d4c8f99f24735aef63c7180be50f0",
        urls = [
            "https://github.com/bazelruby/rules_ruby/archive/v0.4.1.zip",
        ],
        patches = [
            "@io_kythe//third_party:rules_ruby_allow_empty.patch",
        ],
        patch_args = ["-p1"],
    )

    maybe(
        http_archive,
        name = "rules_foreign_cc",
        sha256 = "e60cfd0a8426fa4f5fd2156e768493ca62b87d125cb35e94c44e79a3f0d8635f",
        strip_prefix = "rules_foreign_cc-0.2.0",
        url = "https://github.com/bazelbuild/rules_foreign_cc/archive/0.2.0.zip",
    )

    maybe(
        github_archive,
        repo_name = "llvm/llvm-project",
        commit = "cc46d00a86f89b57008ca878e89538d724b7df90",
        name = "llvm-raw",
        build_file_content = "#empty",
        patch_args = ["-p1"],
        patches = ["@io_kythe//third_party:llvm-bazel-glob.patch"],
        integrity = "sha256-eDavYz6rzf56ip8jl5Ag6c62oz0+eQh1Ac6Hb8wy2sw=",
    )

    maybe(
        github_archive,
        repo_name = "hedronvision/bazel-compile-commands-extractor",
        integrity = "sha256-Df55O1d5hVz3Oz7p9DDgAiX1HzjHBVWTbU3W8bPGXmY=",
        name = "hedron_compile_commands",
        commit = "d6734f1d7848800edc92de48fb9d9b82f2677958",
    )

    # proto_library, cc_proto_library, and java_proto_library rules implicitly
    # depend on @com_google_protobuf for protoc and proto runtimes.
    # Note that if you update protobuf, you must also update some generated
    # proto files:
    #   bazel run //kythe/proto:update
    maybe(
        http_archive,
        name = "com_google_protobuf",
        # sha256 = "8ff511a64fc46ee792d3fe49a5a1bcad6f7dc50dfbba5a28b0e5b979c17f9871",
        strip_prefix = "protobuf-30.0",
        urls = [
            "https://github.com/protocolbuffers/protobuf/releases/download/v30.0/protobuf-30.0.tar.gz",
        ],
        repo_mapping = {"@zlib": "@net_zlib"},
    )
