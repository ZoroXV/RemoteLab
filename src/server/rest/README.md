# REST API description

### `/uploadfile`

#### Request

Use `multipart-form/data` content type.
Fields:
- `name`: name to give to the file on the server
- `file`: content of the file

#### Response

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

### `/command/upload`

#### Request

Type: POST

Use `application/json` content type.
Format:
```json
{
    "serial_number": "<serial_number>",
    "start_address": "<start_address>",
    "port": "<port>",
    "fqbn": "<fqbn>",
    "filename": "<filename>"
}
``` 

Fields:
- `filename`: The name of the binary file to upload on the card (the file should be uploaded on the server before)
- If STM32 controller:
    - `serial_number`: Serial number of the controller (As given by `/command/list_controllers`)
    - `start_address`: The address of the flash, where to store the program in the controller
- If Arduino controller:
    - `port`: The port on which the card is linked (ex: "/dev/ttyUSB0")
    - `fqbn`: The type of the card, following the names of the `arduino-cli` (ex: "arduino:avr:uno")

#### Responses

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

### `/command/list_controllers`

#### Request

Type: GET

#### Response

## Responses

Use `application/json` content type.
Format:
```json
{
    "status": "<status>",
    "message": "<message>",
    "data": [
        {
            "vendor_name": "<vendor_name>",
            "product_name": "<product_name>",
            "serial_number": "<serial_number>",
            "port": "<port>",
            "fqbn": [
                "<fqbn>",
                ...
            ]
        },
        ...
    ]
}

```

Fields:
- `status`: The state of the command (OK or ERROR)
- `message`: A message describing the result of the command (could be empty)
- `data`: List of all microcontrollers connected to the device
- `vendor_name`: Vendor name of the controller (ex: Arduino SA)
- `product_name`: Product name of the controller (ex: Mega 2560 R3 (CDC ACM))
- `port`: The port on which the card is linked (ex: "/dev/ttyUSB0")
- If STM32 controller:
    - `serial_number`: Serial number of the controller
- If Arduino controller:
    - `fqbn`: A list containing the possible types of the card (only one choice if it use a custom vendor/product id, several choices if it use a generic vendor/product id), following the names of the `arduino-cli` (ex: "arduino:avr:uno"). If the device is not an Arduino, do not care of the value

### `/download/remotelab.py`

#### Request

Type: GET

#### Response

## Responses

The CLI tool