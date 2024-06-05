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

- ~~Embed golang binary to Web Application.~~ It is impossible so we need pivot to desktop application. Let's try [wails.io](https://wails.io/docs/gettingstarted/installation)
- Scan local network. [github.com/stefanwichmann/lanscan](https://pkg.go.dev/github.com/stefanwichmann/lanscan)
- FTP Server. [github.com/jlaffaye/ftp](https://pkg.go.dev/github.com/jlaffaye/ftp)
- P2P Network by Internet. [github.com/libp2p/go-libp2p](https://github.com/libp2p/go-libp2p)
