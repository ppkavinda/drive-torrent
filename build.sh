rm -rf releases
mkdir releases
npm run dev
cp -R static releases/static

GOOS=darwin GOARCH=amd64 packr build
mv ./drive-torrent ./releases/darwin-drive-torrent
GOOS=linux GOARCH=amd64 packr build
mv ./drive-torrent ./releases/linux-drive-torrent
GOOS=windows GOARCH=386 packr build
mv ./drive-torrent.exe ./releases/drive-torrent.exe
packr clean