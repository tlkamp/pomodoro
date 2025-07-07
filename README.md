# Pomodoro Timer

A simple command-line Pomodoro timer written in Go.

## What is the Pomodoro method?

1. Decide on the task to be done.
2. Set the Pomodoro timer (default: 25 minutes).
3. Work on the task.
4. When the timer rings, take a short break (default: 5 minutes).
5. Repeat steps 2–4 for four cycles.
6. After four Pomodoros, take a long break (default: 20 minutes).

Repeat as needed.

## Features

- Configurable Pomodoro, short break, and long break durations
- Automatic cycle management (work/break/long break)
- Interruptible sessions with context cancellation
- Easy to extend and customize

## Usage

Build and run the app:

```console
make build
./pomodoro-cli "Your task"
```

By default, the timer runs with:

* Pomodoro: 25 minutes
* Short break: 5 minutes
* Long break: 20 minutes
* 4 intervals per cycle

It can be customized using flags:

```
./pomodoro-cli
A simple command-line Pomodoro timer

Usage:
  pomodoro TASK [flags]

Flags:
  -h, --help                   help for pomodoro
      --intervals int          Number of pomodoro intervals before a long break (default 4)
      --long-break duration    Duration of the long break timer (default 15m0s)
      --pomodoro duration      Duration of the pomodoro timer (default 25m0s)
      --short-break duration   Duration of the short break timer (default 5m0s)
```


## Project Layout

```
.
├── main.go
├── internal/
│   └── session/
├── pomodoro/
├── go.mod
└── README.md
```

## Sample Code
You can adjust durations and intervals in main.go using the session options:

```go
s := session.New(
    session.WithPomodoro(pomodoro.New(25 * time.Minute)),
    session.WithShortBreak(5 * time.Minute),
    session.WithLongBreak(20 * time.Minute),
    session.WithIntervals(4),
)
```