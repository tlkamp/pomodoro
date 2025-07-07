package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tlkamp/pomodoro/pomodoro"
)

const (
	pomodoroStr = "pomodoro"
	shortBreak  = "short-break"
	longBreak   = "long-break"
	intervals   = "intervals"
)

func NewRootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:          "pomodoro TASK",
		Short:        "A simple command-line Pomodoro timer",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go func() {
				sig := <-sigChan
				cancel()
				cmd.Printf("\nReceived signal: %s. Exiting...\n", sig)
			}()

			// Normally I would handle the errors but Cobra seems to validate them pretty well
			pt, _ := time.ParseDuration(cmd.Flag(pomodoroStr).Value.String())
			sbt, _ := time.ParseDuration(cmd.Flag(shortBreak).Value.String())
			lbt, _ := time.ParseDuration(cmd.Flag(longBreak).Value.String())
			itvls, _ := cmd.Flags().GetInt(intervals)

			session := pomodoro.NewSession(
				args[0],
				pomodoro.WithPomodoro(pt),
				pomodoro.WithShortBreak(sbt),
				pomodoro.WithLongBreak(lbt),
				pomodoro.WithIntervals(itvls),
			)

			if err := session.Start(ctx); err != nil {
				// If the user cancels the session, exit successfully
				if errors.Is(err, context.Canceled) {
					return nil
				}

				return errors.Wrap(err, "error in pomodoro session")
			}

			return nil
		},
	}

	rootCmd.PersistentFlags().Duration(pomodoroStr, 25*time.Minute, "Duration of the pomodoro timer")
	rootCmd.PersistentFlags().Duration(shortBreak, 5*time.Minute, "Duration of the short break timer")
	rootCmd.PersistentFlags().Duration(longBreak, 15*time.Minute, "Duration of the long break timer")
	rootCmd.PersistentFlags().Int(intervals, 4, "Number of pomodoro intervals before a long break")

	return rootCmd
}
