# Load Go Rules and Gazelle
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.3/rules_go-0.15.3.tar.gz"],
    sha256 = "97cf62bdef33519412167fd1e4b0810a318a7c234f5f8dc4f53e2da86241c492",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# And our actual deps

go_repository(
    name = "com_github_deckarep_golang_set",
    commit = "1d4478f51bed434f1dadf96dcd9b43aabac66795",
    importpath = "github.com/deckarep/golang-set",
)

go_repository(
    name = "com_github_emirpasic_gods",
    commit = "f6c17b524822278a87e3b3bd809fec33b51f5b46",
    importpath = "github.com/emirpasic/gods",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    commit = "22d885f9ecc78bf4ee5d72b937e4bbcdc58e8cae",
    importpath = "github.com/gin-contrib/sse",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    commit = "b869fe1415e4b9eb52f247441830d502aece2d4d",
    importpath = "github.com/gin-gonic/gin",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "aa810b61a9c79d51363740d207bb46cf8e620ed5",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_jbenet_go_context",
    commit = "d14ea06fba99483203c19d92cfcd13ebe73135f4",
    importpath = "github.com/jbenet/go-context",
)

go_repository(
    name = "com_github_json_iterator_go",
    commit = "1624edc4454b8682399def8740d46db5e4362ba4",
    importpath = "github.com/json-iterator/go",
)

go_repository(
    name = "com_github_kevinburke_ssh_config",
    commit = "81db2a75821ed34e682567d48be488a1c3121088",
    importpath = "github.com/kevinburke/ssh_config",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    commit = "ae18d6b8b3205b561c79e8e5f69bff09736185f4",
    importpath = "github.com/mitchellh/go-homedir",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    commit = "bacd9c7ef1dd9b15be4a9909b8ac7a4e313eec94",
    importpath = "github.com/modern-go/concurrent",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    commit = "4b7aa43c6742a2c18fdef89dd197aaae7dac7ccd",
    importpath = "github.com/modern-go/reflect2",
)

go_repository(
    name = "com_github_pelletier_go_buffruneio",
    commit = "c37440a7cf42ac63b919c752ca73a85067e05992",
    importpath = "github.com/pelletier/go-buffruneio",
)

go_repository(
    name = "com_github_sergi_go_diff",
    commit = "1744e2970ca51c86172c8190fadad617561ed6e7",
    importpath = "github.com/sergi/go-diff",
)

go_repository(
    name = "com_github_src_d_gcfg",
    commit = "f187355171c936ac84a82793659ebb4936bc1c23",
    importpath = "github.com/src-d/gcfg",
)

go_repository(
    name = "com_github_ugorji_go",
    commit = "b4c50a2b199d93b13dc15e78929cfb23bfdf21ab",
    importpath = "github.com/ugorji/go",
)

go_repository(
    name = "com_github_xanzy_ssh_agent",
    commit = "640f0ab560aeb89d523bb6ac322b1244d5c3796c",
    importpath = "github.com/xanzy/ssh-agent",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v8",
    commit = "5f1438d3fca68893a817e4a66806cea46a9e4ebf",
    importpath = "gopkg.in/go-playground/validator.v8",
)

go_repository(
    name = "in_gopkg_src_d_go_billy_v4",
    commit = "59952543636f55de3f860b477b615093d5c2c3e4",
    importpath = "gopkg.in/src-d/go-billy.v4",
)

go_repository(
    name = "in_gopkg_src_d_go_git_v4",
    commit = "d3cec13ac0b195bfb897ed038a08b5130ab9969e",
    importpath = "gopkg.in/src-d/go-git.v4",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    commit = "ec4a0fea49c7b46c2aeb0b51aac55779c607e52b",
    importpath = "gopkg.in/warnings.v0",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "0709b304e793a5edb4a2c0145f281ecdc20838a4",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "161cd47e91fd58ac17490ef4d742dc98bb4cf60e",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "917fdcba135dcbaccd57425db91723541b4427c8",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)
