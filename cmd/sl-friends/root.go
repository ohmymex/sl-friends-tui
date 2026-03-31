package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ohmymex/sl-friends-tui/internal/config"
	"github.com/ohmymex/sl-friends-tui/internal/notify"
	"github.com/ohmymex/sl-friends-tui/internal/tui"
	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:           "sl-friends",
	Short:         "Second Life friends monitor with TUI dashboard",
	Long:          "A terminal UI application to monitor your Second Life friends, groups, and Linden Dollar balance.",
	Version:       version,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          run,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./config.yaml)")
	rootCmd.Flags().String("token", "", "Second Life session token")
	rootCmd.Flags().Bool("token-encoded", false, "token is base64 encoded")
	rootCmd.Flags().String("filter", "", "status filter: online, offline, all")
	rootCmd.Flags().Duration("refresh", 0, "refresh interval (e.g. 5s, 1m)")
	rootCmd.Flags().Bool("show-internal-names", false, "show internal SL names")
	rootCmd.Flags().Bool("show-lindens", false, "show L$ balance")
	rootCmd.Flags().Bool("show-groups", false, "show groups")
	rootCmd.Flags().String("layout", "", "TUI layout: dashboard")
	rootCmd.Flags().StringSlice("notify", nil, "friends to watch (comma-separated)")

	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	viper.BindPFlag("token_encoded", rootCmd.Flags().Lookup("token-encoded"))
	viper.BindPFlag("filter", rootCmd.Flags().Lookup("filter"))
	viper.BindPFlag("refresh", rootCmd.Flags().Lookup("refresh"))
	viper.BindPFlag("show_internal_names", rootCmd.Flags().Lookup("show-internal-names"))
	viper.BindPFlag("show_lindens", rootCmd.Flags().Lookup("show-lindens"))
	viper.BindPFlag("show_groups", rootCmd.Flags().Lookup("show-groups"))
	viper.BindPFlag("layout", rootCmd.Flags().Lookup("layout"))
	viper.BindPFlag("notify.users", rootCmd.Flags().Lookup("notify"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/sl-friends")
	}

	viper.SetEnvPrefix("SLF")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadWithViper(viper.GetViper())
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	token, err := cfg.ResolveToken()
	if err != nil {
		return fmt.Errorf("token required: set via --token flag, SLF_TOKEN env var, or config.yaml\n  %w", err)
	}

	if notifyUsers := viper.GetStringSlice("notify.users"); len(notifyUsers) > 0 {
		cfg.Notify.Enabled = true
		cfg.Notify.Users = notifyUsers
	}

	client := sl.NewClient(token)

	var notifier notify.Notifier
	if cfg.Notify.Enabled {
		notifier = notify.NewDesktopNotifier()
	}

	app := tui.New(client, cfg, notifier)
	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}

	return nil
}
