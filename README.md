# Gocat

Gocat is a Go implementation of the Unix command 'cat', which displays the content of files. Similar to the original 'cat' command, Gocat provides functionality to read one file and output their content to standard output. It's still a WIP project tho, so there's probably some bugs and missing functionalities, like concatenating, creating files, and appending. 


## Features

- Display content of a single file.
- Error handling for file opening and reading.
    


## Installation

Installing with ```go install```

```
  go install github.com/mathesukkj/gocat@latest
```
    
## Usage

```
gocat [FILE]
```

**Examples:**

Display content of a single file:
```
gocat file.txt
```

Display content of multiple files:
```
gocat file.txt file2.txt file3.txt
```

Concatenate two text files and write them to a new file:
```
gocat file.txt file2.txt > newfile.txt
```

Copy the contents of file1 into file2:
```
gocat file.txt > file2.txt
```

Append the contents of file1 to file2:
```
gocat file.txt >> file2.txt
```

Concatenate the files, sort the set of files, and write the output to a new file:
```
gocat file1.txt file2.txt file3.txt | sort > test4
```
