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
* **Post**: Command to post type requests
* **Handle**: Command to handle routes in general
* **HandleFunc**: Command to handle routes from a function
* **NewServeMux**: Command to create create a route handler
* **FileServer**: Command to create directory of files
* **NewRequest**: Commando to create a request customizing
* **Do**: Command to execute a modified request
* **Error**: Command to show server-side errors

### Package io:
* **ReadAll**: Command for general reading
* **CopyBuffer**: Command to copy a buffer to an output

### Package context:
* **Backgroud**: Command to create a new context in background
* **WithTimout**: Command to configure time an existing context
* **cancel**: Command to cancel current context
* **Done**: Returns if the context has ended
* **WithValue**: Command to create a new context with value
* **Value**: Command to retrieve a value passed in context creation 

## JSON

### Package encoding/json:
* **Marshal**: Command to convert a strutc to json
* **NewEncoder**: Command to convert a strutc to json, showing the result and or giving it to someone
* **Unmarshal**: Command to convert json to struct

## TEMPLATE

### Package text/templete:
* **New**: Command to create a new Template
* **Parse**: Command to parse templates
* **ParseFiles**: Command to parse multiple templates
* **Execute**: Command to process the template from a data source
* **Must**: Command to create and parse templates in a unified way

## Database

### Package database/sql:
* **Open**: Connects to the database, note: the corresponding database connection driver must be imported
* **Prepare**: Prepares the slq instruction to be executed next, command used to prevent sql inject
* **Close**: Close the connection
* **Stmt.Exec**: Execute prepared slq statement
* **Stmt.Query**: Execute prepared slq statement, with return of a list of data
* **Stmt.QueryRow**: Execute prepared slq statement, with return one line of data
