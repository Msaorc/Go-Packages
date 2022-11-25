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
* **Handle**: Command to handle routes in general
* **HandleFunc**: Command to handle routes from a function
* **NewServeMux**: Command to create create a route handler
* **FileServer**: Command to create directory of files

### Package io:
* **ReadAll**: Command for general reading

## JSON

### Package encoding/json:
* **Marshal**: Command to convert a strutc to json
* **NewEncoder**: Command to convert a strutc to json, showing the result and or giving it to someone.
* **Unmarshal**: Command to convert json to struct.
