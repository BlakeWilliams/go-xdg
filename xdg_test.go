package xdg

import (
	"os"
	"testing"
)

func TestConfigHome(t *testing.T) {
	testCases := []struct {
		desc    string
		env     map[string]string
		expects string
		error   string
	}{
		{
			desc: "empty env fails",
			env: map[string]string{
				"HOME":            "",
				"XDG_CONFIG_HOME": "",
			},
			error: "could not get user's home directory: $HOME is not defined",
		},
		{
			desc: "returns $XDG_CONFIG_HOME if set",
			env: map[string]string{
				"HOME":            "",
				"XDG_CONFIG_HOME": "/wow",
			},
			expects: "/wow",
		},
		{
			desc: "returns $HOME/.config if $XDG_CONFIG_HOME is not set",
			env: map[string]string{
				"HOME":            "/Users/Mulder",
				"XDG_CONFIG_HOME": "",
			},
			expects: "/Users/Mulder/.config",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			clearEnv := stubEnv(tC.env)
			defer clearEnv()

			val, err := ConfigHome()
			if err != nil {
				if err.Error() != tC.error {
					t.Errorf("Expected error %v, got %v", tC.error, err)
				}

				return
			}

			if val != tC.expects {
				t.Errorf("Expected %v, got %v", tC.expects, val)
			}
		})
	}
}

func TestDataHome(t *testing.T) {
	testCases := []struct {
		desc    string
		env     map[string]string
		expects string
		error   string
	}{
		{
			desc: "empty env fails",
			env: map[string]string{
				"HOME":          "",
				"XDG_DATA_HOME": "",
			},
			error: "could not get user's home directory: $HOME is not defined",
		},
		{
			desc: "returns $XDG_CONFIG_HOME if set",
			env: map[string]string{
				"HOME":          "",
				"XDG_DATA_HOME": "/wow",
			},
			expects: "/wow",
		},
		{
			desc: "returns $HOME/.config if $XDG_DATA_HOME is not set",
			env: map[string]string{
				"HOME":          "/Users/Mulder",
				"XDG_DATA_HOME": "",
			},
			expects: "/Users/Mulder/.local/share",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			clearEnv := stubEnv(tC.env)
			defer clearEnv()

			val, err := DataHome()
			if err != nil {
				if err.Error() != tC.error {
					t.Errorf("Expected error %v, got %v", tC.error, err)
				}

				return
			}

			if val != tC.expects {
				t.Errorf("Expected %v, got %v", tC.expects, val)
			}
		})
	}
}

func TestCacheHome(t *testing.T) {
	testCases := []struct {
		desc    string
		env     map[string]string
		expects string
		error   string
	}{
		{
			desc: "empty env fails",
			env: map[string]string{
				"HOME":          "",
				"XDG_DATA_HOME": "",
			},
			error: "could not get user's home directory: $HOME is not defined",
		},
		{
			desc: "returns $XDG_CACHE_HOME if set",
			env: map[string]string{
				"HOME":           "",
				"XDG_CACHE_HOME": "/wow",
			},
			expects: "/wow",
		},
		{
			desc: "returns $HOME/.cache if $XDG_CACHE_HOME is not set",
			env: map[string]string{
				"HOME":          "/Users/Mulder",
				"XDG_DATA_HOME": "",
			},
			expects: "/Users/Mulder/.cache",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			clearEnv := stubEnv(tC.env)
			defer clearEnv()

			val, err := CacheHome()
			if err != nil {
				if err.Error() != tC.error {
					t.Errorf("Expected error %v, got %v", tC.error, err)
				}

				return
			}

			if val != tC.expects {
				t.Errorf("Expected %v, got %v", tC.expects, val)
			}
		})
	}
}

func TestStateHome(t *testing.T) {
	testCases := []struct {
		desc    string
		env     map[string]string
		expects string
		error   string
	}{
		{
			desc: "empty env fails",
			env: map[string]string{
				"HOME":           "",
				"XDG_STATE_HOME": "",
			},
			error: "could not get user's home directory: $HOME is not defined",
		},
		{
			desc: "returns $XDG_STATE_HOME if set",
			env: map[string]string{
				"HOME":           "",
				"XDG_STATE_HOME": "/wow",
			},
			expects: "/wow",
		},
		{
			desc: "returns $HOME/.cache if $XDG_STATE_HOME is not set",
			env: map[string]string{
				"HOME":           "/Users/Mulder",
				"XDG_STATE_HOME": "",
			},
			expects: "/Users/Mulder/.local/state",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			clearEnv := stubEnv(tC.env)
			defer clearEnv()

			val, err := StateHome()
			if err != nil {
				if err.Error() != tC.error {
					t.Errorf("Expected error %v, got %v", tC.error, err)
				}

				return
			}

			if val != tC.expects {
				t.Errorf("Expected %v, got %v", tC.expects, val)
			}
		})
	}
}

func stubEnv(env map[string]string) func() {
	orig := make(map[string]string, len(env))

	for key, val := range env {
		orig[key] = os.Getenv(key)
		os.Setenv(key, val)
	}

	return func() {
		for key, val := range orig {
			if val == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, val)
			}
		}
	}
}
