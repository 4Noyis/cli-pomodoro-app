# CLI Pomodoro App

## Overview

The CLI Pomodoro App is a command-line tool designed to help you boost productivity and manage your time effectively. The app follows the Pomodoro Technique, a time management method that breaks work into intervals, traditionally 25 minutes in length, separated by short breaks. This tool is fully customizable, allowing you to set your own work and break times, as well as the number of cycles.

## Features

- Customize work duration, break duration, and the number of cycles.
- Real-time progress tracking during work and break intervals.
- A visually appealing countdown timer.
- Tracks total work and break times for a session.

## Installation

1. Clone the repository:
    ```
    git clone https://github.com/4Noyis/cli-pomodoro-app.git
    ```

1. Navigate to the project directory:
    ```
    cd cli-pomodoro-app
    ```

3. Build the application:
    ```
    go build -o pomodoro
    ```

## Usage

Run the application with the following flags:

### Example:

```
./pomodoro -w 30 -b 10 -l 5
```


### Flags:

- `-w, --worktime` (int): Set the duration of the work period (in minutes).

- `-b, --breaktime` (int): Set the duration of the break period (in minutes).

- `-l, --loopcount` (int): Specify the number of Pomodoro cycles.


### Help Command:

```
./pomodoro -h
```

Displays detailed usage instructions.

## Purpose

The purpose of this project is to provide a lightweight, efficient tool to implement the Pomodoro Technique, enabling users to:

- Stay focused on tasks.

- Prevent burnout by taking regular breaks.

- Track productivity through session summaries.


This app is ideal for developers, students, and anyone who wants to improve their time management skills.

## Contributing

Contributions are welcome! If you have ideas or improvements, feel free to open an issue or submit a pull request.


## Author

Developed by [4Noyis](https://github.com/4Noyis).