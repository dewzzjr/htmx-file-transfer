# New App Idea

KirimAja

## Initial Ideas

Application to transfer file from one device to another device in the same local network

- **Explorer**: read file from provided path. Connected receiver can explore the provided path from the sender.
- **P2P connection**: using ftp or tcp, connection initiated by receiver to open port and sender to scan network
- **File transfer**: file transfer from app should be streamed internally and will be downloaded when transfer complete

## Nice to have

- **Create Session**: device from the same network can connect directly by joining the session. Session can be protected by password but can be search by custom name.
- **Transfer file online**: _Sender&rarr;Server&rarr;Receiver_ easier to implement but will need _HUGE_ amount of bandwith from server.
- **Transfer file online P2P version**: harder to implement but server will act only as middleman.

## Technical requirement

- ~~Embed golang binary to Web Application.~~ It is impossible so we need pivot to desktop application. Let's try [wails.io](https://wails.io/docs/gettingstarted/installation). It's also has interesting [wails-htmx-tailwind](https://github.com/PylotLight/wails-htmx-templ-template)
- Scan local network. [github.com/stefanwichmann/lanscan](https://pkg.go.dev/github.com/stefanwichmann/lanscan)
- FTP Server. [github.com/jlaffaye/ftp](https://pkg.go.dev/github.com/jlaffaye/ftp)
- P2P Network by Internet. [github.com/libp2p/go-libp2p](https://github.com/libp2p/go-libp2p)

## wails.io

### About

This is the official Wails Vanilla template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: [https://wails.io/docs/reference/project-config](https://wails.io/docs/reference/project-config)

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on [port 34155](http://localhost:34115). Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.
