package env

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/env"
	"github.com/rsteube/carapace/pkg/style"
)

func init() {
	knownVariables["carapace"] = func() variables {
		_bool := carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)
		return variables{
			Variables: map[string]string{
				// carapace
				"CARAPACE_COVERDIR":      "coverage directory for sandbox tests",
				"CARAPACE_ENV":           "register get-env, set-env and unset-env",
				"CARAPACE_HIDDEN":        "show hidden commands/flags",
				"CARAPACE_LENIENT":       "allow unknown flags",
				"CARAPACE_LOG":           "enable logging",
				"CARAPACE_MATCH":         "match case insensitive",
				"CARAPACE_SANDBOX":       "mock context for sandbox tests",
				"CARAPACE_ZSH_HASH_DIRS": "zsh hash directories",
				// carapace-bin
				"CARAPACE_EXCLUDES": "internal completers to exclude",
				"CARAPACE_BRIDGES":  "implicit bridges",
			},
			VariableCompletion: map[string]carapace.Action{
				// carapace
				"CARAPACE_COVERDIR": carapace.ActionDirectories(),
				"CARAPACE_ENV":      _bool,
				"CARAPACE_HIDDEN":   _bool,
				"CARAPACE_LENIENT":  _bool,
				"CARAPACE_LOG":      _bool,
				"CARAPACE_MATCH": carapace.ActionValuesDescribed(
					"0", "CASE_SENSITIVE",
					"1", "CASE_INSENSITIVE",
				).StyleF(style.ForKeyword),
				"CARAPACE_SANDBOX": carapace.ActionValues(),
				// carapace-bin
				"CARAPACE_EXCLUDES": carapace.Batch(
					carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						c.Setenv(env.CARAPACE_EXCLUDES, "")
						return action.ActionCompleters(action.CompleterOpts{Internal: true}).Invoke(c).ToA()
					}),
					carapace.ActionValuesDescribed("*", "exclude all"),
				).ToA().UniqueList(","),
				"CARAPACE_BRIDGES": carapace.ActionStyledValues(
					"bash", "#d35673",
					"fish", "#7ea8fc",
					"inshellisense", style.Default,
					"zsh", "#efda53",
				).UniqueList(","),
			},
		}
	}

}
