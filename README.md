# daytracker-cli

**Version 1.0**

`daytracker-cli` is a command-line application that allows you to track and save answers to configurable questions. You can define questions, provide multiple-choice answers or take free-form string answers, and specify default values in the configuration. The data is saved in a CSV file.

In the future, you will be able to upload the CSV file to [my Website](https://atiixx.github.io/tools-website) to view statistics and diagrams about your tracked days.

## Features

- Ask configurable questions from a JSON config file
- Provide multiple-choice answers or free-form string input
- Define default values for answers
- Save answers in a CSV file

## Planned Features (Future Releases)

- **Edit a specific date**: Modify or remove entries for a specific day.
- **CLI Configuration**: Change questions, default values, CSV file path, and name directly through different modes accessible via CLI flags.
- **Web Upload and Visualization**: Upload your CSV data to [my Website](https://atiixx.github.io/tools-website) and view detailed statistics and charts based on your tracked entries.

## Installation

To install and run the `daytracker-cli` app, follow these steps:

1. Ensure go is installed

2. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/daytracker-cli.git
   ```

3. Navigate inside the project folder and build the application:

   ```bash
   make build
   ```

   It will create a bin/ folder with the binary.

4. Navigate to the folder with the binary and run it:
   ```bash
   ./daytracker-cli
   ```

The app will generate a default config.json file when run for the first time. You can edit it with your own questions after that.

Configuration
The config.json file should look something like this:

```json
{
  "csv_filepath": "./",
  "csv_filename": "daytracking.csv",
  "questions": [
    {
      "title": "Name",
      "answers": [],
      "default_value": "John Smith",
      "csv": "name"
    },
    {
      "title": "💕 How do you feel?",
      "answers": ["Very good", "Good", "Okay", "Not good", "Bad", "Very bad"],
      "default_value": "2",
      "csv": "feel"
    }
  ]
}
```

- csv_filepath: The path where the CSV file will be saved.
- csv_filename: The name of the CSV file where the data will be stored.
- questions: An array of question objects that you want to answer in the app.
  - title: The question to ask the user.
  - answers: A list of possible answers (leave empty for free-form text input).
  - default_value: The default answer (if the user presses enter without input). Needs to be the number of the answer **(starting from 1)**.
  - csv: The column name for the CSV.

## Features to be Added

- Edit a Specific Date: Modify or remove an entry for a specific day.
- CLI Flags for Config Management: Add flags to change questions, default values, file path, and file name directly in the CLI app.
- Upload to Website for Statistics: Upload your CSV file to a website and view statistical data and diagrams.

## Contributing

Feel free to fork the repository and create a pull request if you'd like to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
