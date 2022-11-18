# Go-Packages
Repository with important go packages

## File Handling

### Package os:
* **Create**: Command to create files
* **Write**: Command to write data to file, for bytes[]
* **WriteString**: Command to write data to file, for string
* **ReadFile**: Command to read data to file
* **Open**: Command to open files
* **Close**: Command to close files
* **Remove**: Command to remove created files

### Package bufio:
* **NewReader**: Command to to create a ReaderBuffer (to read the file in parts it is necessary to create a buffer: make([]byte, x))
* **Read**: Command to read the file by parts. (reader.Read(buffer))

## HTTP Request

### Package net/http:
* **Get**: Command to get type requests

### Package io:
* **ReadAll**: Command for general reading
