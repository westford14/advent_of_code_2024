package go_verify

import (
	"flag"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type xlang struct{}

func NewLanguage() language.Language {
	return &xlang{}
}

func (x *xlang) Name() string {
	return "go_fmt_verify"
}

func (x *xlang) Kinds() map[string]rule.KindInfo {
	return map[string]rule.KindInfo{
		"sh_test": {
			MatchAttrs: []string{"name"},
		},
	}
}

func (x *xlang) Loads() []rule.LoadInfo {
	return []rule.LoadInfo{}
}

func (x *xlang) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
}

func (x *xlang) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

func (x *xlang) KnownDirectives() []string {
	return nil
}

func (x *xlang) Configure(c *config.Config, rel string, f *rule.File) {
}

func (x *xlang) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	rules := make([]*rule.Rule, 0)
	imports := make([]interface{}, 0)
	goFiles := []string{}
	for _, f := range args.RegularFiles {
		if strings.HasSuffix(f, ".go") {
			goFiles = append(goFiles, f)
		}
	}
	goFiles = append(goFiles, "//:go.mod", "//:go.sum", "@go_sdk//:tools", "@go_sdk//:files")

	for _, r := range args.OtherGen {
		if r.Kind() == "go_library" {

			verify := rule.NewRule("sh_test", "go_fmt_verify")

			verify.SetAttr("srcs", []string{"//bzl_tools/go_verify:gofmt.sh"})
			verify.SetAttr("data", goFiles)
			rules = append(rules, verify)
			imports = append(imports, nil)
		}
	}

	return language.GenerateResult{
		Gen:     rules,
		Imports: imports,
	}
}

func (x *xlang) Fix(c *config.Config, f *rule.File) {
}

func (x *xlang) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	return nil
}

func (x *xlang) Embeds(r *rule.Rule, from label.Label) []label.Label {
	return nil
}

func (x *xlang) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, imports interface{}, from label.Label) {
}
