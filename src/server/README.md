# Configuration file of the server

The default name of the configuration file is `config.json` and should be place on the same directory as the server binary. It has the JSON format.

Format:
```json
{
    "vhosts":
    [
        {
            "protocol": "<protocol>",
            "port": "<port>"
        }
    ]
}
```

Fields:
- `protocol`: The protocol to use for this vhost.
- `port`: The port on which this vhost will be listening (Some ports need to run the server as super user).

## Protocols
- REST
