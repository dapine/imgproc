# imgproc

imgproc is an image processing server on top of RabbitMQ and libvips.

## How it works
Each image processing operation (e.g.: resize, rotate) has an exclusive queue. The server expects messages with an image binary and a set of headers, as the processing parameters.
Then, the server responds to client with the processed image or a error message in case of failure.

## Supported operations
- Rotation
- Resizing
- Converting
- Cropping
- Enlarging
- Extracting
- Flipping

For a complete description of what each operation does, follow the [bimg docs](https://pkg.go.dev/github.com/h2non/bimg).

## Sending messages
To send a message, you,ll need to pass the exchange name, the queue key, the image bytes as message body and processing parameters as message headers.
Each operation (queue key name) has a set of key-value parameters (headers).
e.g.:
```elixir
exchange = "image_processing"
key      = "rotate"
bytes    = File.read("image_path")
headers  = [{:angle, 90}]

# Sends a message asking the server to rotate the image by 90 degrees
AmqpClient.send(exchange, key, bytes, headers)
```

The possible operation-headers are:
| operation | headers |
|:----------|:--------|
| `resize`    | `[{"width": <integer>}, {"height": <integer>}]` |
| `rotate`    | `[{"angle": <integer>}]` |
| `convert`    | `[{"target_image_type": "jpeg" \| "jpg" \| "webp" \| "png" \| "tiff" \| "gif" \| "pdf" \| "svg" \| "magick" \| "miff" \| "heif" \| "avif"}]` |
| `crop`    | `[{"width": <integer>}, {"height": <integer>}, {"gravity": "centre" \| "north" \| "east" \| "south" \| "west" \| "smart"}]` |
| `enlarge`    | `[{"width": <integer>}, {"height": <integer>}]` |
| `extract`    | `[{"width": <integer>}, {"height": <integer>}, {"x": <integer>}, {"y": <integer>}]` |
| `flip`    | `[{"axis": "y" \| "vertical" \| "x" \| "horizontal"}]` |


## Client implementation
I have built a client for testing purposes. It's an HTTP API written in Elixir Phoenix supporting all operations implemented by the server:
<https://github.com/dapine/imgproc_api>

## Dependencies
- libvips >= 8.12.0
