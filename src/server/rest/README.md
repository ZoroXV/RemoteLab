# REST API description

## Requests

### `/uploadfile`

Use `multipart-form/data` content type.
Fields:
- `name`: name to give to the file on the server
- `file`: content of the file

### `/command/upload`

Use `application/json` content type.
Format:
```json
{
    "port": "<port>",
    "fqbn": "<fqbn>",
    "filename": "<filename>"
}
``` 

Fields:
- `port`: The port on which the card is linked (ex: "/dev/ttyUSB0")
- `fqbn`: The type of the card, following the names of the `arduino-cli` (ex: "arduino:avr:uno")
- `filename`: The name of the binary file to upload on the card (the file should be uploaded on the server before)

#### Commands
* "UPLOAD"

## Responses

Use `application/json` content type.
Format:
```json
{
    "status": "<status>",
    "message": "<message>"
}
```

Fields:
- `status`: The state of the command (OK or ERROR)
- `message`: A message describing the result of the command (could be empty)
