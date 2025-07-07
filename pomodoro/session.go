package pomodoro

import (
	"context"
	"fmt"
	"time"
)

type Session struct {
	task       string
	pomodoro   time.Duration
	shortBreak time.Duration
	longBreak  time.Duration
	intervals  int
	topic      chan string
}

// NewSession creates a new pomodoro session with the specified task name and options.
func NewSession(taskName string, opts ...Option) *Session {
	session := &Session{
		task:       taskName,
		pomodoro:   25 * time.Minute,
		shortBreak: 5 * time.Minute,
		longBreak:  20 * time.Minute,
		intervals:  4,
	}

	for _, opt := range opts {
		opt(session)
	}

	return session
}

// Start begins the pomodoro session for the specified task. It will run until the context is cancelled.
func (s *Session) Start(ctx context.Context) error {
	fmt.Println("Starting pomodoro session for task:", s.task)

	for {
		for i := range s.intervals {
			if err := runTimer(ctx, s.pomodoro, "pomodoro"); err != nil {
				return err
			}

			breakDuration := s.shortBreak

			if i == 3 {
				fmt.Println("Scheduling long break")
				breakDuration = s.longBreak
			}

			if err := runTimer(ctx, breakDuration, "break"); err != nil {
				return err
			}
		}
	}
}

// Option is a functional option for configuring a Session.
type Option func(s *Session)

// WithPomorodo sets the duration of a pomodoro session.
func WithPomodoro(d time.Duration) Option {
	return func(s *Session) {
		s.pomodoro = d
	}
}

// WithShortBreak sets the duration of a short break.
func WithShortBreak(d time.Duration) Option {
	return func(s *Session) {
		s.shortBreak = d
	}
}

// WithLongBreak sets the duration of a long break.
func WithLongBreak(d time.Duration) Option {
	return func(s *Session) {
		s.longBreak = d
	}
}

// WithIntervals sets the number of pomodoro intervals before a long break.
func WithIntervals(i int) Option {
	return func(s *Session) {
		s.intervals = i
	}
}

func WithTopic(topic chan string) Option {
	return func(s *Session) {
		s.topic = topic
	}
}

func runTimer(ctx context.Context, duration time.Duration, purpose string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(duration):
		fmt.Printf("%s timer completed (%s)\n", purpose, duration)
		return nil
	}
}
