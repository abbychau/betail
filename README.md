# BeTail

This is a Go program that mimics the behavior of `tail -f` by continuously displaying the appended content of a file.

## Usage

1. Make sure you have Go installed on your system.
2. Clone this repository or download the `tail.go` file.
3. Open a terminal and navigate to the directory where `tail.go` is saved.
4. Execute the following command:

   ```shell
   go run tail.go <filename>
   ```

Replace <filename> with the path to the file you want to tail.

1. The program will start displaying the appended content of the file.
2. If the file is truncated (its size becomes smaller than the current offset position), the program will start reading from the beginning of the file.

# Customization

You can adjust the sleep duration in the main loop of the `main.go` file to control the frequency at which the program checks for new content in the file. Modify the line:

    ```go
    time.Sleep(1 * time.Second) // Adjust the sleep duration as per your needs
    ```

to your desired duration.

# License

This project is licensed under the MIT License.

Feel free to customize the content and formatting according to your specific needs.

# Why make this?

Because I thought that `tail -f` cannot keep reading truncated files, but I was wrong, lol.