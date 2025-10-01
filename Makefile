build: clean
	mkdir -p ./bin/
	go build -o ./bin/qube .

test:
	./bin/qube https://archlinux.org/releng/releases/2025.09.01/torrent/

clean:
	rm -rf ./bin/
	rm /var/folders/cx/p0t_w01s06l4ytpfgxrj8jz80000gn/T/somefilenametest & echo ""

